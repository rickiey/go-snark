package handlers

import (
	"encoding/json"
	"fmt"
	"go-snark/model"
	"go-snark/resp"
	"net/http"

	ffi "github.com/filecoin-project/filecoin-ffi"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

// PingPong 调度机服务pingpong
func PingPong(c *gin.Context) {
	resp.NormalResult(c, "PONG")
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
			//finalResult string
			//status int
			res struct {
				Proof []byte
			}
		)
		result, err := ffi.SealCommitPhase2([]byte(data.Phase1Out), abi.SectorNumber(data.SectorID), abi.ActorID(actorID))
		if nil != err {
			glog.Infof("sector %d compute failed: %s", data.SectorID, err.Error())
			//status = 3
		} else {
			fmt.Println(string(result))
			err = json.Unmarshal(result, &res)
			if nil != err {
				glog.Infof("sector %d json unmarshal failed: %s", data.SectorID, err.Error())
				//status = 3
			} else {

			}
		}

	}()

	c.JSON(http.StatusOK, "OK")
}
