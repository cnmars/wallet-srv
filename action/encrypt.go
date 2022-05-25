package action

import (
	"wallet-srv/lib"
	"wallet-srv/lib/resp"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

type reqData struct {
	PrivKey string `json:"data"`
}

//二次加密私钥
func EncryptAction(c *gin.Context) {

	var red reqData
	err := lib.DecodeHTTPReqJSON(c, &red)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("EncryptAction: ", err)
		return
	}

	encryptPriv, err := lib.XXTEAEncrypt(red.PrivKey)
	if err != nil {
		resp.Error(c, resp.XXTEAM_ENCRYPT_ERROR)
		glog.Error("EncryptAction: ", err)
		return
	}

	decryptPriv, err := lib.XXTEADecrypt(encryptPriv)
	if err != nil {
		resp.Error(c, resp.XXTEAM_ENCRYPT_ERROR)
		glog.Error("EncryptAction: ", err)
		return
	}

	if decryptPriv != red.PrivKey {
		resp.Error(c, resp.XXTEAM_ENCRYPT_ERROR)
		glog.Error("EncryptAction: encrypt privKey not same to decrypt privkey")
		return
	}

	resp.Ok(c, encryptPriv)
}

//二次解密私钥
func DecryptAction(c *gin.Context) {

	var red reqData
	err := lib.DecodeHTTPReqJSON(c, &red)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("DecryptAction: ", err)
		return
	}

	decryptPriv, err := lib.XXTEADecrypt(red.PrivKey)
	if err != nil {
		resp.Error(c, resp.XXTEAM_ENCRYPT_ERROR)
		glog.Error("DecryptAction: ", err)
		return
	}

	encryptPriv, err := lib.XXTEAEncrypt(decryptPriv)
	if err != nil {
		resp.Error(c, resp.XXTEAM_ENCRYPT_ERROR)
		glog.Error("DecryptAction: ", err)
		return
	}

	if encryptPriv != red.PrivKey {
		resp.Error(c, resp.XXTEAM_ENCRYPT_ERROR)
		glog.Error("DecryptAction: encrypt privKey not same to decrypt privkey")
		return
	}

	resp.Ok(c, decryptPriv)
}
