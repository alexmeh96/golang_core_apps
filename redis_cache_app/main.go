package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type Store struct {
	data  map[int]string
	cache Cacher
}

func NewStore(c Cacher) *Store {
	data := map[int]string{
		1: "Message1",
		2: "Message2",
		3: "Message3",
	}
	return &Store{
		data:  data,
		cache: c,
	}
}

func (s *Store) Get(key int) (string, error) {
	val, ok := s.cache.Get(key)
	if ok {
		if err := s.cache.Remove(key); err != nil {
			fmt.Println(err)
		}
		fmt.Println("return the value from the cache")
		return val, nil
	}

	val, ok = s.data[key]
	if !ok {
		return "", fmt.Errorf("key not found: %d", key)
	}

	if err := s.cache.Set(key, val); err != nil {
		return "", err
	}

	fmt.Println("returning key from internal storage")

	return val, nil
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// срок жизни данных в кэше
	ttl := time.Second * 3
	s := NewStore(NewRedisCache(client, ttl))

	for i := 0; i < 6; i++ {
		val, err := s.Get(3)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(val)

		if i == 2 {
			time.Sleep(4 * time.Second)
			fmt.Println("timeout ttl")
		}
	}
}
