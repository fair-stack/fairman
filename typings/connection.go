// Package global
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-04 11:04
package typings

import "io"

type Connection interface {
	Open() error
	Close() error

	// ExportRaw write the db contents to filename as json (the schema needs defining)
	// takedbtakejsontake（take）
	ExportRaw(filename string) error

	//Rollback(force bool) error
	//MigrateData(migratorParams *database.MigratorParameters, force bool) error

	// TODO: this one is very database specific atm
	BackupTo(w io.Writer) error
	GetDatabaseFilename() string
	GetStorePath() string

	SetServiceName(bucketName string) error
	GetObject(bucketName string, key []byte, object interface{}) error
	UpdateObject(bucketName string, key []byte, object interface{}) error
	DeleteObject(bucketName string, key []byte) error
	DeleteAllObjects(bucketName string, matching func(o interface{}) (id string, ok bool)) error
	// GetNextIdentifier(bucketName string) int
	CreateObject(bucketName string, id string, fn func(id string) (string, interface{})) error
	CreateObjectWithId(bucketName string, id string, obj interface{}) error
	CreateObjectWithSetSequence(bucketName string, id string, obj interface{}) error
	GetAll(bucketName string, obj interface{}, append func(o interface{}) (interface{}, error)) error
	GetAllWithJsoniter(bucketName string, obj interface{}, append func(o interface{}) (interface{}, error)) error
	ConvertToKey(v string) []byte
}
