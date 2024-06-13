package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)
func main() {
	fmt.Println("Hello, World!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("welcome")
		fmt.Println(r.RemoteAddr)
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Fatal("Invalid IP address")
		}
		fmt.Println(ip)
	});

	http.ListenAndServe(":8080", nil)
}