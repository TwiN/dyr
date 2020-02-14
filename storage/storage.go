package storage

import (
	"fmt"
	"github.com/TwinProduction/dyr/config"
	"github.com/TwinProduction/dyr/core"
	"github.com/TwinProduction/gdstore"
	"sort"
)

const (
	DatabaseFileName = "dyr.data"
)

func openStore() *gdstore.GDStore {
	return gdstore.New(fmt.Sprintf("%s/%s", config.GetConfigDir(), DatabaseFileName))
}

func SaveNote(text string, tags []string) error {
	store := openStore()
	defer store.Close()
	note := core.Note{
		Id:   0,
		Text: text,
		Tags: tags,
	}
	lastId, ok, err := store.GetInt("current_note_id")
	if err != nil {
		return err
	}
	if !ok {
		note.Id = 0
		_ = store.Put("current_note_id", []byte("0"))
	} else {
		note.Id = uint64(lastId) + 1
		_ = store.Put("current_note_id", []byte(fmt.Sprintf("%d", note.Id)))
	}
	return store.Put(fmt.Sprintf("%d", note.Id), note.ToBytes())
}

func GetNoteById(id uint64) (*core.Note, error) {
	store := openStore()
	defer store.Close()
	bytes, _ := store.Get(fmt.Sprintf("%d", id))
	return core.NoteFromBytes(bytes)
}

func GetAllNotes() ([]*core.Note, error) {
	var notes []*core.Note
	store := openStore()
	defer store.Close()
	keys := store.Keys()
	sort.Strings(keys)
	for _, key := range keys {
		bytes, _ := store.Get(key)
		note, err := core.NoteFromBytes(bytes)
		if err != nil {
			continue
		}
		notes = append(notes, note)
	}
	return notes, nil
}
