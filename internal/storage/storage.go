package storage

import (
	"fmt"
	"sync"
)

type InMEM struct {
	mu   sync.Mutex
	data map[string]string
}

var i InMEM

func NewMem() *InMEM {
	var i = &InMEM{
		data: make(map[string]string),
	}
	return i
}

func (i *InMEM) Put(orig, short string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.data[short] = orig

	return nil
}

func (i *InMEM) Get(short string) (s string, e error) {
	i.mu.Lock()

	defer i.mu.Unlock()

	s, ok := i.data[short]
	if !ok {
		s = ""
		e = fmt.Errorf("not found")
	}
	return s, e

}
