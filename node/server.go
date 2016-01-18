package node

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/rasmusrygaard/chord/chord"
)

type Server struct {
	Backend chord.Node
}

func (n *Server) Launch() {
	err := rpc.Register(n)
	if err != nil {
		log.Fatalf("Format of NetworkNode service isn't correct. %s", err)
	}

	rpc.HandleHTTP()

	//start listening for messages on port 1234
	l, e := net.Listen("tcp", fmt.Sprintf("127.0.0.1:8080"))
	if e != nil {
		log.Fatalf("Couldn't start listening on port 1234. Error %s", e)
	}
	log.Println("Serving RPC handler")
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatalf("Error serving: %s", err)
	}
}

type NoArgs int

func (s *Server) Identifier(_ *NoArgs, ret *chord.ID) error {
	id := s.Backend.Identifier()
	ret = &id
	return nil
}

func (s *Server) FindSuccessor(id *chord.ID, ret *chord.Node) error {
	succ := s.Backend.FindSuccessor(*id)
	ret = &succ
	return nil
}

func (s *Server) Successor(_ *NoArgs, ret *chord.Node) error {
	succ := s.Backend.Successor()
	ret = &succ
	return nil
}

func (s *Server) Predecessor(_ *NoArgs, ret *chord.Node) error {
	succ := s.Backend.Predecessor()
	ret = &succ
	return nil
}

func (s *Server) Join(n *chord.Node, ret *NoArgs) error {
	return s.Backend.Join(*n)
}
