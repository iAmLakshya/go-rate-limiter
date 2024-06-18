package main

import (
	"errors"
	"hash/fnv"
	"log"
	"net"
	"net/http"
	"sync"
)

type KVStore struct {
	size int
	data map[uint64]int // [key]value
	mu sync.Mutex
}
func (m *KVStore) set(k string, v int) {
	h := hash(k)
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[h]; !ok {
		m.size++
	}
	m.data[hash(k)] = v;
}
func (m *KVStore) get(k string) (int, error) {
	h := hash(k)
	m.mu.Lock()
	defer m.mu.Unlock()
	val, ok := m.data[h]
	if ok {
		return val, nil
	}
	return val, errors.New("this key does not exist in the map") 
}

func (m *KVStore) inc(k string) {
	val, err := m.get(k)
	if err != nil {
		m.set(k, 1)
		return
	}
	m.set(k, val + 1)
}

func hash(id string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(id))
	return h.Sum64()
}

func main() {
	m := KVStore{size: 0, data: make(map[uint64]int)}
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Fatal("Invalid IP address")
		}
		
		m.inc(ip)
	});

	http.ListenAndServe(":8080", nil)
}