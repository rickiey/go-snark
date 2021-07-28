package main

import (
	"context"
	"encoding/json"
	"errors"
	pb "go-snark/cmd/windowpost/proto"
	"io/ioutil"
	"log"

	ffi "github.com/filecoin-project/filecoin-ffi"
	"github.com/filecoin-project/go-state-types/abi"
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
	//log.Println("receive task: ", task.MinerID, " rand: ", task.Random)
	//log.Println("privbyte: ", task.Privsectors)

	// 将参数记录在文件里
	param, err := json.Marshal(task)
	if nil != err {
		log.Println("SnarkServer.Marshal: ", err.Error())
		return resp, err
	}

	err = ioutil.WriteFile("window_param.out", param, 0644)
	if nil != err {
		log.Println("SnarkServer.WriteFile: ", err.Error())
		return resp, err
	}

	ss := &ffi.SortedPrivateSectorInfo{}
	err = ss.UnmarshalJSON(task.Privsectors)
	if nil != err {
		log.Println("SnarkServer.AllocateTask: ", err.Error())
		return resp, err
	}

	//log.Println("priv: ", ss.Values())
	proof, faulty, err := ffi.GenerateWindowPoSt(abi.ActorID(task.MinerID), *ss, abi.PoStRandomness(task.Random))

	if nil != err {
		log.Println("SnarkServer.GenerateWindowPoSt: ", err.Error())
		return resp, err
	}

	for _, p := range proof {
		temp := &pb.PoStProof{
			PoStProof:  int64(p.PoStProof),
			ProofBytes: p.ProofBytes,
		}
		resp.Proofs = append(resp.Proofs, temp)
	}

	for _, s := range faulty {
		temp := uint64(s)
		resp.SectorNumber = append(resp.SectorNumber, temp)
	}

	return resp, nil
}

// HeartBeat 接收心跳检测
func (s *SnarkServer) HeartBeat(ctx context.Context, req *pb.HeartBeatRequest) (*pb.HeartBeatResponse, error) {
	log.Println("receive heart beat: ", req.SentTime)
	resp := &pb.HeartBeatResponse{
		Status: "OK",
	}

	return resp, nil
}

// RegisterWorker snark server 端无需实现
func (s *SnarkServer) RegisterWorker(ctx context.Context, req *pb.WindowWorkerRequest) (*pb.WindowWorkerResponse, error) {
	return nil, errors.New("unimplements")
}
