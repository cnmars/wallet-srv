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
	flag.Lookup("log_dir").Value.Set("./logs/sign")
	flag.Lookup("v").Value.Set("0")

	gin.SetMode(gin.ReleaseMode)
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	defer glog.Flush()

	router := gin.Default()

	lib.AuthMachineID()

	//IP whitelist
	whiteList := lib.GetIPWhiteList(conf.ACCESS_SIGNAPI_WHITELIST)
	router.Use(lib.IPWhiteList(whiteList))

	//sign api v1
	v1 := router.Group("/v1")
	{
		v1.POST("/btc/sign", action.BtcSignAction)   // BTC
		v1.POST("/hd/sign", action.HdSignAction)     // BCH / DASH / DOGE / QTUM / LTC / BSV
		v1.POST("/eth/sign", action.EthSignAction)   // ETH / HT / BSC / MATIC / ETC
		v1.POST("/trx/sign", action.TrxSignAction)   // TRX
		v1.POST("/sol/sign", action.SolSignAction)   // SOL
		v1.POST("/fil/sign", action.FilSignAction)   // FIL
		v1.POST("/ada/sign", action.AdaSignAction)   // ADA
		v1.POST("/xrp/sign", action.XrpSignAction)   // XRP
		v1.POST("/luna/sign", action.LunaSignAction) // LUNA
		v1.POST("/near/sign", action.NearSignAction) // NEAR
		v1.POST("/dot/sign", action.DotSignAction)   // DOT
		v1.POST("/avax/sign", action.AvaxSignAction) // AVAX
	}

	v2 := router.Group("get")
	{
		// get coin address no used list
		v2.POST("/address", action.FindAddress)
	}

	glog.Infof("SignAPIServer start. Listen Address: %s\n", conf.SIGN_LISTEN)
	err := router.Run(conf.SIGN_LISTEN)

	if err != nil {
		glog.Errorf("SignAPIServer start failed. ", err)
		return
	}
}
