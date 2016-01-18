package node

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/rasmusrygaard/chord/chord"
)

type Remote struct {
	Port int
	IP   net.IP
}

type Client struct {
	client *rpc.Client
}

func NewClient(n Remote) (*Client, error) {
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("%s:%d", n.IP, n.Port))
	if err != nil {
		log.Fatalf("Failed to connect to remote node, err=%q", err)
		return nil, err
	}
	return &Client{client: client}, nil
}

func (c *Client) Identifier() chord.ID {
	var id chord.ID
	err := c.client.Call("Server.Identifier", new(int), &id)
	if err != nil {
		log.Fatalf("Error when calling Identifier(), err=%q", err)
	}
	return id
}
