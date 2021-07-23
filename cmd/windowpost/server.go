package main

import (
	"context"
	"errors"
	pb "go-snark/cmd/windowpost/proto"
	"log"

	ffi "github.com/filecoin-project/filecoin-ffi"
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
	resp := new(pb.TaskResponse)
	log.Println("receive task: ", task.MinerID, " rand: ", task.Random)
	log.Println("privbyte: ", task.Privsectors)
	//ss := make([]ffi.PrivateSectorInfo, 0)
	//err := json.Unmarshal(task.Privsectors, &ss)
	ss := &ffi.SortedPrivateSectorInfo{}
	err := ss.UnmarshalJSON(task.Privsectors)
	if nil != err {
		log.Println(err)
		return resp, err
	}

	log.Println("priv: ", ss.Values())

	return resp, nil
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
