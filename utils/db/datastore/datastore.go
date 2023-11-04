// Package datastore
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-08-04 15:46
package datastore

import (
	"cnic/fairman-gin/typings"
	"io"
)

// NewStore initializes a new Store and the associated services
func NewStore(connection typings.Connection) *Store {
	return &Store{
		connection: connection,
	}
}

func (store *Store) Open() (err error) {
	err = store.connection.Open()
	if err != nil {
		return err
	}

	err = store.initServices()
	if err != nil {
		return err
	}

	return nil
}

func (store *Store) Close() error {
	return store.connection.Close()
}

// BackupTo backs up db to a provided writer.
// It does hot backup and doesn't block other database reads and writes
func (store *Store) BackupTo(w io.Writer) error {
	return store.connection.BackupTo(w)
}
