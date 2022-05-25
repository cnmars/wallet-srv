package main

import (
	"flag"
	"runtime"
	"wallet-srv/action"
	"wallet-srv/conf"
	"wallet-srv/lib"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func init() {
	flag.Parse()
	flag.Lookup("alsologtostderr").Value.Set("true")
	flag.Lookup("log_dir").Value.Set("../logs/encrypt")
	flag.Lookup("v").Value.Set("3")

	gin.SetMode(gin.ReleaseMode)
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	defer glog.Flush()

	lib.AuthMachineID()

	router := gin.Default()

	//IP white list
	whiteList := lib.GetIPWhiteList(conf.ACCESS_ENCRYPT_WHITELIST)
	router.Use(lib.IPWhiteList(whiteList))

	v1 := router.Group("/v1")
	{
		v1.POST("/encrypt", action.EncryptAction)
		v1.POST("/decrypt", action.DecryptAction)
	}

	glog.Infof("EncryptServer start. Listen Address: %s\n", conf.ENCRYPT_LISTEN)
	err := router.Run(conf.ENCRYPT_LISTEN)

	if err != nil {
		glog.Info(err)
		return
	}

}
