package types

type Config struct {
	Dot Dot `json:"dot"`
}

type Dot struct {
	RpcHost  string `json:"rpc_host"`
	MetaFile string `json:"meta_file"`
}
