package dot

import (
	"fmt"
	"io/ioutil"
	"wallet-srv/conf"
	"wallet-srv/lib"
	gsrpc "wallet-srv/lib/pkg/dot/gsrpc"
	dtypes "wallet-srv/lib/pkg/dot/types"
	"wallet-srv/lib/types"

	"github.com/golang/glog"
)

func UpdateMeta(c *types.Config) {

	api, err := gsrpc.NewSubstrateAPI(c.Dot.RpcHost)
	if err != nil {
		glog.Info("gsrpc.NewSubstrateAPI: ", err)
		return
	}
	meta, err := api.RPC.State.GetMetadataLatest()
	metaBuf, err := dtypes.EncodeToHexString(&meta)
	if err != nil {
		glog.Info("json.Marshal(meta): ", err)
		return
	}

	err = ioutil.WriteFile(c.Dot.MetaFile, []byte(metaBuf), 0644)
	if err != nil {
		glog.Info("dot meta data ioutil.WriteFile: ", err)
		return
	}
	glog.Info("Dot GetMetadataLatest update finished.")
	return
}

//load meta cache
func GetMeta() (*dtypes.Metadata, error) {
	config, err := lib.LoadConfig(conf.DEFAULT_CONFIG)
	if err != nil {
		return nil, err
	}
	meta, err := ioutil.ReadFile(config.Dot.MetaFile)
	if err != nil {
		return nil, err
	}

	var metaData dtypes.Metadata
	err = dtypes.DecodeFromBytes(meta, &metaData)
	fmt.Println("metaData: ", metaData)
	if err != nil {
		return nil, err
	}

	return &metaData, nil
}
