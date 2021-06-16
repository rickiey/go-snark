package services

import (
	"fmt"
	"go-snark/dao"

	"github.com/golang/glog"
)

func RegisterGpuWorker(gpuType, ipAddr string) error {
	methodPath := "services.RegisterGpuWorker"
	// 先查询，如果已经注册，则更改状态为重新上线
	exists, err := dao.QueryWorker(ipAddr)
	if nil != err {
		return fmt.Errorf("%s.QueryWorker: %s", methodPath, err.Error())
	}

	if exists {
		glog.Infof("change worker online: %s", ipAddr)
		err = dao.ChangeWorkerStatus(ipAddr, "is_online", 1)
		if nil != err {
			return fmt.Errorf("%s.ChangeWorkerStatus: %s", methodPath, err.Error())
		}
		glog.Info("change success!")
		return nil
	}

	// 注册
	glog.Infof("register worker: %s", ipAddr)
	err = dao.InsertWorker(gpuType, ipAddr)
	if nil != err {
		return fmt.Errorf("%s.InsertWorker: %s", methodPath, err.Error())
	}
	glog.Info("worker register success!")

	return nil
}
