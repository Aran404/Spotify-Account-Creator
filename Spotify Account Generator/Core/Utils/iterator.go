package utils

import (
	"sync"
)

type Iterator struct {
	Mutex    *sync.Mutex
	Index    int
	FileSize int
	Data     []string
}

func NewFromFile(path string) *Iterator {
	Lines := Readlines(path)

	return &Iterator{
		Data:     Lines,
		FileSize: len(Lines),
		Index:    0,
		Mutex:    &sync.Mutex{},
	}
}

func (in *Iterator) Next() string {
	in.Mutex.Lock()
	defer in.Mutex.Unlock()

	if in.Index >= in.FileSize {
		in.Index = 0
	}

	in.Index++
	return in.Data[in.Index-1]
}
