package resp

const (
	HTTP_OK              = 0
	REQ_ERROR            = 10000
	PARAM_ERROR          = 10001
	XXTEAM_ENCRYPT_ERROR = 10002
	SERVER_INNER_ERROR   = 10003
	API_ENCRYPT_ERROR    = 10004
	TOKEN_EXPIRED        = 10005
	SIGN_ERROR           = 10006
	FROM_ADDRESS         = 10007
)

var ERRORCODE = map[int]string{
	0:     "ok",
	10000: "invalid request",
	10001: "param error",
	10002: "xxtea encrypt fail.",
	10003: "server inner error",
	10004: "api encrypt fail.",
	10005: "token expired",
	10006: "sign error",
	10007: "address error",
}
