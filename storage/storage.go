package storage

import (
	"fmt"
	"github.com/TwinProduction/dyr/config"
	"github.com/TwinProduction/dyr/core"
	"github.com/TwinProduction/gdstore"
)

const (
	DatabaseFileName = "dyr.data"
)

func openStore() *gdstore.GDStore {
	return gdstore.New(fmt.Sprintf("%s/%s", config.GetConfigDir(), DatabaseFileName))
}

func SaveNote(note *core.Note) error {
	store := openStore()
	defer store.Close()
	return store.Put(fmt.Sprintf("%d", note.Id), note.ToBytes())
}

func GetNoteById(id uint64) (*core.Note, error) {
	store := openStore()
	defer store.Close()
	bytes, _ := store.Get(fmt.Sprintf("%d", id))
	return core.NoteFromBytes(bytes)
}
