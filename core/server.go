// Package core
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-11 16:35
package core

import (
	"cnic/fairman-gin/global"
	"cnic/fairman-gin/initialize"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/cloudflare/tableflip"
	"go.uber.org/zap"
)

type server interface {
	Serve(l net.Listener) error
	Shutdown(ctx context.Context) error
}

func RunServer() {
	// judgepidjudge
	if global.FairConf.System.PIDFile == "" {
		log.Fatalln("pidFile not found!")
	}

	pidFile := global.FairConf.System.PIDFile
	upg, err := tableflip.New(tableflip.Options{
		PIDFile: pidFile,
	})
	if err != nil {
		panic(err)
	}
	defer upg.Stop()

	// Do an upgrade on SIGHUP
	var exit bool
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGHUP, syscall.SIGUSR2, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
		for s := range sig {
			switch s {
			case syscall.SIGHUP, syscall.SIGUSR2:
				log.Println("Upgrade start:", s)
				err := upg.Upgrade()
				if err != nil {
					log.Println("Upgrade failed:", err)
					continue
				}
				log.Println("Upgrade succeeded")
			case syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT:
				log.Println("Shutdown Server start:", s)
				exit = true
				upg.Stop()
				log.Println("Shutdown Server ...")
			}
		}
	}()

	address := fmt.Sprintf(":%s", global.FairConf.System.Port)
	ln, err := upg.Fds.Listen("tcp", address)
	if err != nil {
		log.Fatalln("Can't listen:", err)
	}

	Router := initialize.Routers()
	s := initServer(Router)
	go func() {
		if err := s.Serve(ln); err != nil {
			log.Fatalln("Server error:", err)
		}
	}()
	// Ensure text order output
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.FairLog.Info("server run success on ", zap.String("address: ", address))

	fmt.Printf(`
        ______      _      __  ___           _____                          
       / ____/___ _(_)____/  |/  /___ _____ / ___/___  ______   _____  _____
      / /_  / __ /  / ___/ /|_/ / __ / __  \\__ \/ _ \/ ___/ | / / _ \/ ___/
     / __/ / /_/ / / /  / /  / / /_/ / / / /__/ /  __/ /   | |/ /  __/ /    
    /_/    \__,_/_/_/  /_/  /_/\__,_/_/ /_/____/\___/_/    |___/\___/_/     

	Welcome to cnic/fairman-gin
	current version:V0.0.2 Beta
	Default backend interface running address:http://127.0.0.1%s
	`, address)

	err = ioutil.WriteFile(pidFile, []byte(strconv.Itoa(os.Getpid())), 0755)
	if err != nil {
		panic(err)
	}
	if err := upg.Ready(); err != nil {
		panic(err)
	}
	<-upg.Exit()

	// Make sure to set a deadline on exiting the process
	// after upg.Exit() is closed. No new upgrades can be
	// performed if the parent doesn't exit.
	time.AfterFunc(20*time.Second, func() {
		log.Println("Graceful shutdown timed out")
		os.Exit(1)
	})
	// Wait for connections to drain.
	err = s.Shutdown(context.Background())
	if err != nil {
		global.FairLog.Error(err.Error())
	}

	if exit {
		_ = os.Remove(pidFile)
	}
}
