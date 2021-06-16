package main

import (
	"flag"
	"fmt"
	"go-snark/conf"
	"go-snark/dao"
	"go-snark/router"
	"go-snark/services"

	"os"
	"os/signal"
	"syscall"

	ffi "github.com/filecoin-project/filecoin-ffi"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func init() {
	flag.Set("stderrthreshold", "INFO")
	flag.Parse()

	glog.MaxSize = 1024 * 1024 * 50

	// 加载配置文件
	if err := conf.Init(); nil != err {
		panic(err)
	}
	gin.SetMode(conf.Conf.Server.Env)
	// 设置环境变量，限定该进程可见的GPU
	err := os.Setenv("CUDA_VISIBLE_DEVICES", conf.Conf.Server.GpuVisible)
	if nil != err {
		panic(err)
	}
	glog.Infof("CUDA_VISIBLE_DEVICES is %s", conf.Conf.Server.GpuVisible)
	glog.Infof("This worker's addr is {}", conf.Conf.Server.IpAddr)
	glog.Infof("GPU is %s", conf.Conf.Server.GpuType)
	err = CheckGpu()
	if nil != err {
		panic(err)
	}

	dao.InitDB()

	// 如果改gpu worker 还未注册，则进行注册
	err = services.RegisterGpuWorker(conf.Conf.Server.GpuType, conf.Conf.Server.IpAddr)
	if nil != err {
		panic(err)
	}
}

func CheckGpu() error {
	gpus, err := ffi.GetGPUDevices()
	if nil != err {
		return err
	}

	if len(gpus) < 2 {
		return fmt.Errorf("current gpu number is %d", len(gpus))
	}

	glog.Infof("This worker has %d GPU devices", len(gpus))

	return nil
}

func checkErr(err error, desc string) {
	if nil == err {
		glog.Infof("%s close success", desc)
	} else {
		glog.Infof("%s close error: %s", desc, err.Error())
	}
}

func main() {
	// 加载路由
	r := router.InitRouter()
	// 在协程中启动服务
	go func() {
		err := r.Run(conf.Conf.Server.Port)
		if nil != err {
			panic(err)
		}
	}()
	glog.Infof("server is running %s", conf.Conf.Server.Port)
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	s := <-exit
	glog.Infof("server get a signal %s", s.String())
	switch s {
	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
		// 服务停止时，关闭数据库连接
		err := dao.DB.Close()
		checkErr(err, "DB")
		glog.Infof("server exit")
		glog.Flush()
		return
	case syscall.SIGHUP:
	default:
		return
	}
}
