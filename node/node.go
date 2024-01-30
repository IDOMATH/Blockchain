package node

import (
	"context"
	"fmt"
	"github.com/idomath/Blockchain/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"net"
)

type Node struct {
	version string
	peers   map[net.Addr]*grpc.ClientConn
	proto.UnimplementedNodeServer
}

func NewNode() *Node {
	return &Node{
		version: "0.1",
	}
}

func (n *Node) Handshake(ctx context.Context, v *proto.Version) (*proto.Version, error) {
	ourVersion := &proto.Version{
		Version: n.version,
		Height:  100,
	}

	p, _ := peer.FromContext(ctx)
	fmt.Printf("received version: %+v from %s", v, p.Addr)

	return ourVersion, nil
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Println("received tx from: ", peer)
	return &proto.Ack{}, nil
}
