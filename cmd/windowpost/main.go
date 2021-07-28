package main

import (
	"context"
	"flag"
	pb "go-snark/cmd/windowpost/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var LocalAddr = flag.String("local_addr", "", "the local addr:port")
var RemoteAddr = flag.String("remote_addr", "", "remote server addr:port")

/**
snark （每台挂载存储，2/4张卡，对应2/4个进程，端口区分，各为一个worker）
	server: 1. 接收任务进行计算，并返回结果
	client: 1. 向window机注册snark server
window机
	server: 1. 管理client
			2. 对client做心跳，当前可用数不足，则电话告警
	client: 1. 调用snark server端分配任务，并接收计算结果， 1个partition，一个client
*/
func init() {
	flag.Parse()

	if *LocalAddr == "" || *RemoteAddr == "" {
		panic("local addr or remote add cannot empty")
	}
}

func main() {
	listener, err := net.Listen("tcp", *LocalAddr) // 监听本地端口
	if err != nil {
		panic(err)
	}
	log.Println("grpc server Listing on", *LocalAddr)

	grpcServer := grpc.NewServer() // 新建gRPC服务器实例

	client := NewSnarkClient(*RemoteAddr)
	status, err := client.RegisterWorker(context.Background(), *LocalAddr)

	if nil != err {
		panic(err)
	}

	if status != 1 {
		panic("register failed")
	}

	log.Println("register success")

	server := NewSnarkServer()

	pb.RegisterWindowGrpcServer(grpcServer, server)

	if err = grpcServer.Serve(listener); err != nil { //用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
		panic(err)
	}
}
