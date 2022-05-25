package action

import (
	"wallet-srv/lib"
	"wallet-srv/lib/resp"
	"wallet-srv/lib/wallet"
	"wallet-srv/model"

	ctypes "wallet-srv/lib/types"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func FindAddress(c *gin.Context) {
	var req ctypes.ReqAddress
	err := lib.DecodeHTTPReqJSON(c, &req)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("FindAddress: ", err)
		return
	}

	cModel, table := wallet.GetDBModel(req.Coin)

	var data *[]model.CoinAddress
	data = cModel.FindData(table, req.Limit)

	var list []string
	if len(*data) > 0 {
		for _, row := range *data {
			ok := cModel.UpdateData(table, row.ID)
			if ok {
				list = append(list, row.Address)
			}
		}
	}

	resp.Ok(c, list)
}
