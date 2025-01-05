package main

import (
	"log"
)

func main() {
	server := NewAPIServer(":3000")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
