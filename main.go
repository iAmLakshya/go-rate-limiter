package main

import (
	"errors"
	"fmt"
	"hash/fnv"
	"log"
	"net"
	"net/http"
)

type KVStore struct {
	size int
	data map[uint64]string // [key]value
}
func (m *KVStore) set(k string, v string) {
	h := hash(k)
	if _, ok := m.data[h]; !ok {
		m.size++
	}
	m.data[hash(k)] = v;
}
func (m *KVStore) get(k string) (string, error) {
	h := hash(k)
	val, ok := m.data[h]
	if ok {
		return val, nil
	}
	return val, errors.New("this key does not exist in the map") 
}

func hash(id string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(id))
	return h.Sum64()
}

func main() {
	m := KVStore{size: 0, data: make(map[uint64]string)}
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("welcome")
		fmt.Println(r.RemoteAddr)
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Fatal("Invalid IP address")
		}
		fmt.Println(ip)
		
		m.set(ip, "RAHH")
		val, _ := m.get(ip)
		fmt.Println(val)
	});

	http.ListenAndServe(":8080", nil)
}