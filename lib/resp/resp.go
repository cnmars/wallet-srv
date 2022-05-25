package resp

import (
	"encoding/json"
	"net/http"
	"wallet-srv/lib"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, code int) {

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": ERRORCODE[code],
		"data":    "",
	})
}

func Ok(c *gin.Context, data interface{}) {

	d, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    SERVER_INNER_ERROR,
			"message": ERRORCODE[SERVER_INNER_ERROR],
			"data":    "",
		})
		return
	}

	encryptData, err := lib.APIAESEncrypt(string(d))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    API_ENCRYPT_ERROR,
			"message": ERRORCODE[API_ENCRYPT_ERROR],
			"data":    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    HTTP_OK,
		"message": ERRORCODE[HTTP_OK],
		"data":    encryptData,
	})
}
