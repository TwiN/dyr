package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Note struct {
	Id   uint64
	Text string
	Tags []string
}

func (n Note) ToBytes() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(n)
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

func NoteFromBytes(b []byte) (*Note, error) {
	var buffer bytes.Buffer
	decoder := gob.NewDecoder(&buffer)
	buffer.Write(b)
	var note Note
	if err := decoder.Decode(&note); err != nil {
		return nil, fmt.Errorf("failed to decode note from gdstore: %s", err.Error())
	}
	return &note, nil
}
