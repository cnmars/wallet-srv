package wallet

import (
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/golang/glog"
)

func CheckTrxAddress(addr string) bool {

	_, err := common.DecodeCheck(addr)
	if err != nil {
		glog.Error("CheckTrxAddress error: ", err)
		return false
	}

	if len(addr) != 34 {
		return false
	}

	prefix := addr[0:1]
	return prefix == "T"
}

func CheckSolAddress(addr string) bool {
	return len(addr) == 44
}
