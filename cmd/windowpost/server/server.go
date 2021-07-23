package main

import (
	"context"
	"errors"
	pb "go-snark/cmd/windowpost"

	"github.com/golang/glog"
)

type SnarkServer struct {
	pb.UnimplementedWindowGrpcServer
}

func NewSnarkServer() *SnarkServer {
	return &SnarkServer{}
}

// AllocateTask 接收计算任务
func (s *SnarkServer) AllocateTask(ctx context.Context, task *pb.TaskRequest) (*pb.TaskResponse, error) {
	glog.Info("receive task: ", task.MinerID)
	return nil, nil
}

// HeartBeat 接收心跳检测
func (s *SnarkServer) HeartBeat(ctx context.Context, req *pb.HeartBeatRequest) (*pb.HeartBeatResponse, error) {
	glog.Info("receive heart beat: ", req.SentTime)
	return nil, nil
}

// RegisterWorker snark server 端无需实现
func (s *SnarkServer) RegisterWorker(ctx context.Context, req *pb.WindowWorkerRequest) (*pb.WindowWorkerResponse, error) {
	return nil, errors.New("unimplements")
}
