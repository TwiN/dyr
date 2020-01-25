package core

import (
	"bytes"
	"encoding/gob"
)

type Note struct {
	Id   uint64
	Text string
	Tags []string
}

func (n Note) ToBytes() []byte {
	var buffer bytes.Buffer            // Stand-in for a buffer connection
	encoder := gob.NewEncoder(&buffer) // Will write to buffer.
	err := encoder.Encode(n)
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
}
