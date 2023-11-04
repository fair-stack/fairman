// Package database
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-04 15:51
package database

import (
	"cnic/fairman-gin/typings"
	"cnic/fairman-gin/utils/db/database/boltdb"
	"fmt"
)

func NewDatabase(storeType, storePath string) (connection typings.Connection, err error) {
	switch storeType {
	case "boltdb":
		return &boltdb.DbConnection{Path: storePath}, nil
	}
	return nil, fmt.Errorf("Unknown storage database: %s", storeType)
}
