package handlers

import (
	"encoding/hex"
	"go-snark/conf"
	"go-snark/dao"
	"go-snark/model"
	"net/http"

	ffi "github.com/filecoin-project/filecoin-ffi"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

// PingPong 调度机服务pingpong
func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, "PONG")
}

func SealCommitPhase2(c *gin.Context) {
	var data model.Commit2Out
	err := c.BindJSON(&data)
	if nil != err {
		glog.Infof("resolve param: %s", err.Error())
		return
	}

	maddr, err := address.NewFromString(data.Miner)
	if nil != err {
		glog.Infof("NewFromString %s: %s", data.Miner, err.Error())
		return
	}
	actorID, err := address.IDFromAddress(maddr)
	if nil != err {
		glog.Infof("IDFromAddress %s: %s", maddr.String(), err.Error())
		return
	}
	go func() {
		var (
			status      int
			finalResult string
		)

		result, err := ffi.SealCommitPhase2([]byte(data.Phase1Out), abi.SectorNumber(data.SectorID), abi.ActorID(actorID))
		if nil != err {
			glog.Infof("sector %d compute failed: %s", data.SectorID, err.Error())
			status = 3
		} else {
			status = 2
			finalResult = hex.EncodeToString(result)
		}
		glog.Infof("sector: %d, result: %s", data.SectorID, finalResult)
		// TODO, 先往redis里面放一份，避免连接mysql的外网出问题

		err = dao.ChangeTaskStatus(finalResult, data.Miner, conf.Conf.Server.IpAddr, status, data.SectorID)
		if nil != err {
			glog.Infof("update task status failed: %s", err.Error())
			return
		}

		glog.Infof("update %s %d task status success", data.Miner, data.SectorID)
		glog.Info("worker is free now !")
	}()

	c.JSON(http.StatusOK, "OK")
}
