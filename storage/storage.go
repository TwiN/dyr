package storage

import (
	"fmt"
	"github.com/TwinProduction/dyr/config"
	"github.com/TwinProduction/dyr/core"
	"github.com/TwinProduction/gdstore"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

const (
	DatabaseFileName = "dyr.data"
	CurrentNoteIdKey = "current_note_id"
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
	lastId, ok, err := store.GetInt(CurrentNoteIdKey)
	if err != nil {
		return err
	}
	if !ok {
		note.Id = 0
		_ = store.Put(CurrentNoteIdKey, []byte("0"))
	} else {
		note.Id = uint64(lastId) + 1
		_ = store.Put(CurrentNoteIdKey, []byte(fmt.Sprintf("%d", note.Id)))
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
	sort.Slice(keys, func(i, j int) bool {
		first, err := strconv.ParseInt(keys[i], 10, 0)
		if err != nil {
			return false
		}
		second, err := strconv.ParseInt(keys[j], 10, 0)
		if err != nil {
			return false
		}
		return second > first
	})
	for _, key := range keys {
		if key == CurrentNoteIdKey {
			continue
		}
		bytes, _ := store.Get(key)
		note, err := core.NoteFromBytes(bytes)
		if err != nil {
			continue
		}
		notes = append(notes, note)
	}
	return notes, nil
}

func GetNotesByTag(tag string) ([]*core.Note, error) {
	var notes []*core.Note
	store := openStore()
	defer store.Close()
	values := store.Values()
	for _, value := range values {
		note, err := core.NoteFromBytes(value)
		if err != nil {
			continue
		}
		for _, t := range note.Tags {
			if tag == t {
				notes = append(notes, note)
				break
			}
		}
	}
	return notes, nil
}

func GetRandomNote() (*core.Note, error) {
	store := openStore()
	keys := store.Keys()
	store.Close()
	var validNoteIds []string
	for _, key := range keys {
		if key != CurrentNoteIdKey {
			validNoteIds = append(validNoteIds, key)
		}
	}
	if validNoteIds == nil || len(validNoteIds) == 0 {
		return nil, fmt.Errorf("no valid notes available")
	}
	var err error
	var id int
	if len(validNoteIds) == 1 {
		id, err = strconv.Atoi(validNoteIds[0])
	} else {
		rand.Seed(time.Now().UnixNano())
		randomKey := validNoteIds[rand.Intn(len(validNoteIds))]
		id, err = strconv.Atoi(randomKey)
	}
	if err != nil {
		return nil, err
	}
	return GetNoteById(uint64(id))
}

func DeleteNoteById(id uint64) error {
	store := openStore()
	defer store.Close()
	return store.Delete(fmt.Sprintf("%d", id))
}
