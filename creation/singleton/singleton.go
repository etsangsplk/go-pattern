package singleton

import (
	"log"
	"sync"
)

type Singleton map[string]string

var (
	once     sync.Once
	instance Singleton
)

func Instance() Singleton {
	once.Do(func() {
		log.Println("Create an instance.")
		instance = make(Singleton)
	})
	return instance
}
