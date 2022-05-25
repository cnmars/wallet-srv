package conf

var (

	/*** API请求加密秘钥 ****/
	API_AES_IK = "imMeLx9MZM36ru52yoCkb60yXtjdHgWN"
	API_AES_IV = "Pijl45zRIOsAuBvJ"

	//一次加密秘钥，使用公私钥对称加密算法
	PRIV_AES_IK = "lVlv3irxebksvrcmuLntyklgddx1tpqt" //私钥一次加密IK 32位
	PRIV_AES_IV = "5phpriLlqwurPyd2"                 //私钥一次加密IV 16位

	//二次加密秘钥，采用自定义XXTEA加密算法
	PRIV_XXTEA_KEY = "paIBMluOaEjlkwFgm/Eh/r4X9OW9UUSA5DQY/jUd50g=" //私钥二次加密秘钥

	//钱包地址库链接方式
	MYSQL_DSN = "root:workspace@tcp(127.0.0.1:3306)/hannan?charset=utf8&parseTime=True&loc=Local"

	//验签访问白名单 - 只提供给需要验签的节点访问
	ACCESS_SIGNAPI_WHITELIST = "127.0.0.1,127.0.0.2"
	//加密服务访问白名单 - 只提供给验签服务访问
	ACCESS_ENCRYPT_WHITELIST = "127.0.0.1"

	//二次加解密网关监听地址
	ENCRYPT_LISTEN = "127.0.0.1:4030" //加密HTTP监听地址
	//验签API网关监听地址
	SIGN_LISTEN = "0.0.0.0:4040" //验签HTTP监听地址

	//二次加解密网关API
	ENCRYPT_HTTP_API = "http://127.0.0.1:4030/v1"

	//绑定部署机器的机器码 - 获取机器码的方式为 bin/wallet-tool
	MACHINE_UUID = "45B3CF5E2441CFC42B7A097A7F03F50F" //"59BA9AA693227F825A59C5AEE80001C6"

	DEFAULT_CONFIG = "../conf/config.json"
)
