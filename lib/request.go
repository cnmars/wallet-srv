package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

type reqJson struct {
	Content string `json:"content"`
	Token   string `json:"token"`
}

func DecodeHTTPReqJSON(c *gin.Context, data interface{}) error {

	var r reqJson
	err := c.ShouldBind(&r)
	if err != nil {
		return err
	}

	//校验token是否有效
	sToken, err := APIAESDecrypt(r.Token)
	if err != nil {
		glog.Error(err)
		return errors.New("token decrypt fail")
	}
	sTime, _ := strconv.ParseInt(sToken, 10, 64)
	if time.Now().Unix()-sTime > 15 {
		return errors.New("token expired")
	}

	//解密数据
	decodeStr, err := APIAESDecrypt(r.Content)
	if err != nil {
		glog.Info("decode request content fail.")
		return err
	}

	err = json.Unmarshal([]byte(decodeStr), &data)
	if err != nil {
		glog.Info("decode content is not json data.")
		return err
	}

	return nil
}

func IPWhiteList(whiteList map[string]bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		if !whiteList[clientIP] {
			fmt.Println("IP:", clientIP, " request denied.")
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":    http.StatusForbidden,
				"message": "Permission denied",
				"data":    "",
			})
			return
		}
	}
}

func GetIPWhiteList(iplist string) map[string]bool {
	var ipList []string
	ipList = strings.Split(iplist, ",")
	if len(ipList) == 0 {
		ipList = append(ipList, "127.0.0.1")
	}

	whiteList := make(map[string]bool)
	for _, ip := range ipList {

		if ip == "" {
			continue
		}
		whiteList[ip] = true

	}

	return whiteList
}
