package lib

import (
	"wallet-srv/conf"
	"wallet-srv/lib/encrypt"
)

//私钥 加密方法
func SignAESEncrypt(content string) (string, error) {
	return encrypt.AESEncrypt(content, []byte(conf.PRIV_AES_IK), []byte(conf.PRIV_AES_IV))
}

//私钥 解密方法
func SignAESDecrypt(content string) (string, error) {
	return encrypt.AESDecrypt(content, []byte(conf.PRIV_AES_IK), []byte(conf.PRIV_AES_IV))
}

//API 请求加密方法
func APIAESEncrypt(content string) (string, error) {
	return encrypt.AESEncrypt(content, []byte(conf.API_AES_IK), []byte(conf.API_AES_IV))
}

//API 请求解密方法
func APIAESDecrypt(content string) (string, error) {
	return encrypt.AESDecrypt(content, []byte(conf.API_AES_IK), []byte(conf.API_AES_IV))
}

//私钥 二次加密方法
func XXTEAEncrypt(content string) (string, error) {
	return encrypt.XXTEAEncrypt(content)
}

//私钥 二次解密方法
func XXTEADecrypt(content string) (string, error) {
	return encrypt.XXTEADecrypt(content)
}

//API POST XXTEA解密
// func XXTEAPOSTDecrypt(content string) (string, error) {

// }

// //API POST XXTEA加密
// func XXTEAPOSTEncrypt(content string) (string, error) {

// }
