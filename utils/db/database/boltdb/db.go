// Package boltdb
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-04 10:54
package boltdb

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/pkg/errors"

	bolt "go.etcd.io/bbolt"
)

const (
	DatabaseFileName = "fairman.db"
)

var (
	ErrObjectNotFound = errors.New("Object not found inside the database")
)

type DbConnection struct {
	Path string

	*bolt.DB
}

func (connection *DbConnection) GetDatabaseFilename() string {
	return DatabaseFileName
}

func (connection *DbConnection) GetStorePath() string {
	return connection.Path
}

// Open opens and initializes the BoltDB database.
// Open Open and initializeBoltDBOpen and initialize。
func (connection *DbConnection) Open() error {

	// Disabled for now.  Can't use feature flags due to the way that works
	// databaseExportPath := path.Join(connection.Path, fmt.Sprintf("raw-%s-%d.json", DatabaseFileName, time.Now().Unix()))
	// if err := connection.ExportRaw(databaseExportPath); err != nil {
	// 	log.Printf("raw export to %s error: %s", databaseExportPath, err)
	// } else {
	// 	log.Printf("raw export to %s success", databaseExportPath)
	// }

	databasePath := path.Join(connection.Path, DatabaseFileName)
	fmt.Println()

	db, err := bolt.Open(databasePath, 0600, &bolt.Options{Timeout: 10 * time.Second})
	if err != nil {
		return err
	}
	connection.DB = db
	return nil
}

// Close closes the BoltDB database.
// Safe to being called multiple times.
// Close closeBoltDBclose。
// Can be called multiple times。
func (connection *DbConnection) Close() error {
	if connection.DB != nil {
		return connection.DB.Close()
	}
	return nil
}

// BackupTo backs up db to a provided writer.
// It does hot backup and doesn't block other database reads and writes
// Backing up the database to the provided writer。
// It performs hot backup，It performs hot backup
func (connection *DbConnection) BackupTo(w io.Writer) error {
	return connection.View(func(tx *bolt.Tx) error {
		_, err := tx.WriteTo(w)
		return err
	})
}

// ExportRaw exports the raw database to a JSON file.
// ExportRaw Export the original database toJSONExport the original database to。
func (connection *DbConnection) ExportRaw(filename string) error {
	databasePath := path.Join(connection.Path, DatabaseFileName)
	if _, err := os.Stat(databasePath); err != nil {
		return fmt.Errorf("stat on %s failed: %s", databasePath, err)
	}

	b, err := exportJson(databasePath)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, b, 0600)
}

// ConvertToKey returns an 8-byte big endian representation of v.
// This function is typically used for encoding integer IDs to byte slices
// so that they can be used as BoltDB keys.
// ConvertToKey takevtake8take。
// This function is typically used to encode integersIDThis function is typically used to encode integers，This function is typically used to encode integersBoltDBThis function is typically used to encode integers。
// Therefore, they can be used asBoltDBTherefore, they can be used as。
// 2022.08.04 10:54 takeinttakestring
func (connection *DbConnection) ConvertToKey(v string) []byte {
	return []byte(v)
}

// SetServiceName CreateBucket is a generic function used to create a bucket inside a database database.
// SetServiceName CreateBucketIt is a universal function，It is a universal functionbucket。
func (connection *DbConnection) SetServiceName(bucketName string) error {
	return connection.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return nil
	})
}

// GetObject is a generic function used to retrieve an unmarshalled object from a database database.
// GetObject It is a universal function，It is a universal function。
func (connection *DbConnection) GetObject(bucketName string, key []byte, object interface{}) error {
	var data []byte

	err := connection.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		value := bucket.Get(key)
		if value == nil {
			return ErrObjectNotFound
		}

		data = make([]byte, len(value))
		copy(data, value)

		return nil
	})
	if err != nil {
		return err
	}

	return UnmarshalObject(data, object)
}

// UpdateObject is a generic function used to update an object inside a database database.
// UpdateObject It is a universal function，It is a universal function。
func (connection *DbConnection) UpdateObject(bucketName string, key []byte, object interface{}) error {
	return connection.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		data, err := MarshalObject(object)
		if err != nil {
			return err
		}

		err = bucket.Put(key, data)
		if err != nil {
			return err
		}

		return nil
	})
}

// DeleteObject is a generic function used to delete an object inside a database database.
// DeleteObject It is a universal function，It is a universal function。
func (connection *DbConnection) DeleteObject(bucketName string, key []byte) error {
	return connection.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		return bucket.Delete(key)
	})
}

// DeleteAllObjects delete all objects where matching() returns (id, ok).
// DeleteAllObjects Delete all matching objects（id，ok）。
// TODO: think about how to return the error inside (maybe change ok to type err, and use "notfound"?
// Consider how to return internal errors（Consider how to return internal errorsokConsider how to return internal errorserr，Consider how to return internal errors“notfound”）？
func (connection *DbConnection) DeleteAllObjects(bucketName string, matching func(o interface{}) (id string, ok bool)) error {
	return connection.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var obj interface{}
			err := UnmarshalObject(v, &obj)
			if err != nil {
				return err
			}

			if id, ok := matching(obj); ok {
				err := bucket.Delete(connection.ConvertToKey(id))
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}

// GetNextIdentifier is a generic function that returns the specified bucket identifier incremented by 1.
// GetNextIdentifier It is a universal function，It is a universal functionbucketIt is a universal function1。
// 2022.08.04 10:54 Abandonment，AbandonmentidAbandonment
func (connection *DbConnection) GetNextIdentifier(bucketName string) int {
	var identifier int

	connection.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		id, err := bucket.NextSequence()
		if err != nil {
			return err
		}
		identifier = int(id)
		return nil
	})

	return identifier
}

// CreateObject creates a new object in the bucket, using the next bucket sequence id
// CreateObject staybucketstay，staybucketstayid
func (connection *DbConnection) CreateObject(bucketName string, id string, fn func(id string) (string, interface{})) error {
	return connection.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		newId, obj := fn(id)

		data, err := MarshalObject(obj)
		if err != nil {
			return err
		}

		return bucket.Put(connection.ConvertToKey(newId), data)
	})
}

// CreateObjectWithId creates a new object in the bucket, using the specified id
// CreateObjectWithId staybucketstay，stayid
func (connection *DbConnection) CreateObjectWithId(bucketName string, id string, obj interface{}) error {
	return connection.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		data, err := MarshalObject(obj)
		if err != nil {
			return err
		}

		return bucket.Put(connection.ConvertToKey(id), data)
	})
}

// CreateObjectWithSetSequence creates a new object in the bucket, using the specified id, and sets the bucket sequence
// avoid this :)
// CreateObjectWithSetSequence staybucketstay，stayid，staybucketstay
// Avoid this situation：）
func (connection *DbConnection) CreateObjectWithSetSequence(bucketName string, id string, obj interface{}) error {
	return connection.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		data, err := MarshalObject(obj)
		if err != nil {
			return err
		}

		return bucket.Put(connection.ConvertToKey(id), data)
	})
}

// GetAll returns all objects in the bucket.
// GetAll returnbucketreturn。
func (connection *DbConnection) GetAll(bucketName string, obj interface{}, append func(o interface{}) (interface{}, error)) error {
	err := connection.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			err := UnmarshalObject(v, obj)
			if err != nil {
				return err
			}
			obj, err = append(obj)
			if err != nil {
				return err
			}
		}

		return nil
	})
	return err
}

// GetAllWithJsoniter
// TODO: decide which Unmarshal to use, and use one...
// Decide which unmarshaller to use，Decide which unmarshaller to use。。。
func (connection *DbConnection) GetAllWithJsoniter(bucketName string, obj interface{}, append func(o interface{}) (interface{}, error)) error {
	err := connection.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			err := UnmarshalObjectWithJsoniter(v, obj)
			if err != nil {
				return err
			}
			obj, err = append(obj)
			if err != nil {
				return err
			}
		}

		return nil
	})
	return err
}
