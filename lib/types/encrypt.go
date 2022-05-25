package types

type Content struct {
	Data string `json:"data"`
}

type ReqBody struct {
	Content string `json:"content"`
	Token   string `json:"token"`
}

type RespBody struct {
	Code    int    `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}
