// Package initialize
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-04 15:56
package initialize

//import (
//	"cnic/fairman-gin/global"
//	"cnic/fairman-gin/utils"
//	"cnic/fairman-gin/utils/db/database"
//	"cnic/fairman-gin/utils/db/datastore"
//	"fmt"
//	"os"
//	"path"
//	"time"
//)
//
//func WhileInitDataStore() error {
//	fmt.Println("while init data store")
//	//fmt.Println(global.FairStore)
//	if global.FairStore != nil {
//		return nil
//	}
//	var i = 1
//	for true {
//		fmt.Println(i)
//		i++
//		time.Sleep(time.Second * 2)
//		err := InitDataStore()
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//
//		if global.FairStore != nil && global.FairDB != nil {
//			endpoints, err := global.FairStore.Endpoint().Endpoints()
//			if err != nil {
//				continue
//			}
//			fmt.Println("endpoints length: ", len(endpoints))
//			break
//		}
//	}
//	return nil
//}
//
//func InitDataStore() error {
//	pathStr := path.Join(global.FairConf.System.Root, "db", "boltdb")
//
//	if ok, _ := utils.PathExists(pathStr); !ok {
//		err := os.MkdirAll(pathStr, 0755)
//		if err != nil {
//			return err
//		}
//	}
//
//	connection, err := database.NewDatabase("boltdb", pathStr)
//	if err != nil {
//		return err
//	}
//
//	store := datastore.NewStore(connection)
//
//	err = store.Open()
//	if err != nil {
//		return utils.Errorf("failed opening store: %v", err)
//	}
//
//	global.FairDB = connection
//	global.FairStore = store
//
//	return nil
//}
