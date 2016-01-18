package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/rasmusrygaard/chord/node"
)

func main() {
	server := node.Server{
		Backend: &node.LinkedListNode{ID: 123, Next: nil, Prev: nil},
	}
	go server.Launch()

	time.Sleep(5 * time.Second)
	remote := node.Remote{Port: 8080, IP: net.IPv4(127, 0, 0, 1)}
	client, err := node.NewClient(remote)
	if err != nil {
		log.Fatalf("Fatal error with client: err=%q", err)
	}
	fmt.Println(client.Identifier())
}
