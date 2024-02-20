package main

import (
	"context"
	"github.com/idomath/Blockchain/node"
	"github.com/idomath/Blockchain/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	makeNode(":3000", []string{})
	time.Sleep(time.Second)
	makeNode(":3001", []string{":3000"})
	time.Sleep(time.Second)
	makeNode(":3002", []string{":3001"})

	select {}
}

func makeNode(listenAddr string, bootstrapNodes []string) *node.Node {
	n := node.NewNode()
	go n.Start(listenAddr, bootstrapNodes)

	return n
}

func makeTransaction() {
	client, err := grpc.Dial(":3000",
		grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := proto.NewNodeClient(client)

	tx := &proto.Transaction{
		Version: 1,
	}

	_, err = c.HandleTransaction(context.TODO(), tx)
	if err != nil {
		log.Fatal(err)
	}
	version := &proto.Version{
		Version:    "0.1",
		Height:     1,
		ListenAddr: ":4000",
	}

	_, err = c.Handshake(context.TODO(), version)
	if err != nil {
		log.Fatal(err)
	}
}
