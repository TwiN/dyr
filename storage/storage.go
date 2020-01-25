package storage

import (
	"fmt"
	"github.com/TwinProduction/dyr/config"
	"github.com/TwinProduction/dyr/core"
	"github.com/dgraph-io/badger"
)

const (
	DatabaseFileName = "badger.db"
)

func openDatabase() *badger.DB {
	db, err := badger.Open(badger.DefaultOptions(fmt.Sprintf("%s/%s", config.GetConfigDir(), DatabaseFileName)))
	if err != nil {
		panic(err)
	}
	return db
}

func SaveNote(note *core.Note) error {
	db := openDatabase()
	defer db.Close()
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("1"), note.ToBytes())
		return err
	})
	return err
}
