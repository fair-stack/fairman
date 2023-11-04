// Package exec
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2023-08-23 01:05
package exec

import (
	"net"

	"golang.org/x/crypto/ssh"
)

func NewSSHDeployer(ip, user, pass string) (client *ssh.Client, err error) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	return ssh.Dial("tcp", ip+":22", config)
}
