package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Alfred-Jijo/todoApp-api/api"
	"github.com/Alfred-Jijo/todoApp-api/storage"
)

func main() {
	listenAddr := flag.String("address", ":8080", "localhost port")
	flag.Parse()

	server := api.NewServer(*listenAddr, storage.PageData{})
	fmt.Println("Server Listening on port:", *listenAddr)
	log.Fatal(server.Start())
}
