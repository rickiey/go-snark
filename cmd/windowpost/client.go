package main

import (
	"context"
	pb "go-snark/cmd/windowpost/proto"
	"log"

	"google.golang.org/grpc"
)

type SnarkClient struct {
	RemoteAddr string
}

// NewSnarkClient ..
func NewSnarkClient(remoteAddr string) *SnarkClient {
	return &SnarkClient{
		RemoteAddr: remoteAddr,
	}
}

func (c *SnarkClient) connect(function string) (pb.WindowGrpcClient, *grpc.ClientConn, error) {
	log.Printf("window snark client(%s) connect to %s\n", function, c.RemoteAddr)
	conn, err := grpc.Dial(c.RemoteAddr, grpc.WithInsecure()) //连接gRPC服务器
	if err != nil {
		return nil, nil, err
	}
	client := pb.NewWindowGrpcClient(conn) //建立客户端
	return client, conn, nil
}

// RegisterWorker ..
func (c *SnarkClient) RegisterWorker(ctx context.Context, localAddr string) (int, error) {
	client, conn, err := c.connect("RegisterWorker")
	if err != nil {
		return -1, err
	}
	defer conn.Close()

	req := new(pb.WindowWorkerRequest)
	req.WorkerAddr = localAddr

	resp, err := client.RegisterWorker(ctx, req) //调用方法
	if err != nil {
		return -1, err
	}
	return int(resp.Status), nil
}
