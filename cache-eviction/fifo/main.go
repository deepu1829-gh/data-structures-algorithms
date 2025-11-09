package main

import "fmt"

type FIFO struct {
	capacity int
	queue    []string
	store    map[string]string
}

func NewFifo(capacity int) *FIFO {
	return &FIFO{
		capacity: capacity,
		store:    make(map[string]string),
	}
}

func (f *FIFO) Set(key, value string) {
	if len(f.queue) >= f.capacity {
		evict := f.queue[0]
		f.queue = f.queue[1:]
		delete(f.store, evict)
	}

	f.queue = append(f.queue, key)
	f.store[key] = value
}

func (f *FIFO) Get(key string) (string, bool) {
	val, ok := f.store[key]

	return val, ok
}

func main() {
	cache := NewFifo(2)

	cache.Set("1", "hello 3")
	cache.Set("2", "hello 2")
	cache.Set("3", "hello 3")

	fmt.Println(cache.Get("1"))
	fmt.Println(cache.Get("3"))
}
