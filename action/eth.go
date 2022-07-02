package action

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"wallet-srv/lib"
	"wallet-srv/lib/resp"
	"wallet-srv/lib/tx/eth"
	ctypes "wallet-srv/lib/types"
	"wallet-srv/lib/wallet"
	"wallet-srv/model"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

//支持ETH链/BSC链/HECO链/Matic链/OKX链/ETC链验签
func EthSignAction(c *gin.Context) {
	/*
		   {
			   "from": "",
			   "to": "",
			   "amount": 12313,		//单位为Wei
			   "nonce": 0,
			   "gaslimit": 60000,
			   "gasprice": 2121231, //单位为Wei
			   "contract": "", 		//为空则为ETH转账，否则为对应代币转账
			   "chainid": 1, 		//1为ETH链，56为BSC链，61为ETC链，128为HECO链，137为Matic链, 66为OK链，250 FTM 10 Optimism 592 Astar
		   }
		   **/
	var reqData ctypes.EthTx
	err := lib.DecodeHTTPReqJSON(c, &reqData)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("ETHSignAction: DecodeHTTPReqJSON error", err)
		return
	}

	nonce := uint64(reqData.Nonce)
	value := big.NewInt(reqData.Amount)
	fromAddress := reqData.From
	toAddress := reqData.To
	contractAddress := reqData.Contract
	chainID := big.NewInt(int64(reqData.ChainID))
	gasLimit := uint64(reqData.GasLimit)
	gasPrice := big.NewInt(reqData.GasPrice)
	var data []byte

	if !common.IsHexAddress(fromAddress) || !common.IsHexAddress(toAddress) {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("ETHSignAction ReqData address error")
		return
	}

	privKey, err := getEthPrivateKey(fromAddress)
	if err != nil {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("ETHSignAction: getData error ", err)
		return
	}

	var tx *types.Transaction

	if contractAddress != "" {
		data = getData(toAddress, value)
		tx = types.NewTransaction(nonce, common.HexToAddress(contractAddress), big.NewInt(0), gasLimit, gasPrice, data)
	} else {
		tx = types.NewTransaction(nonce, common.HexToAddress(toAddress), value, gasLimit, gasPrice, data)
	}

	//签名交易
	signTx, err := eth.SignTx(common.HexToAddress(fromAddress), privKey, tx, chainID)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("ETHSignAction SignTxEth error: ", err)
		return
	}

	var buf bytes.Buffer
	if err := signTx.EncodeRLP(&buf); err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("ETHSignAction EncodeRLP error: ", err)
		return
	}

	signTxHex := fmt.Sprintf("0x%x", buf.Bytes())
	glog.Info("EthSignTx: ", signTxHex)
	resp.Ok(c, signTxHex)
}

func getEthPrivateKey(address string) (*ecdsa.PrivateKey, error) {
	cModel, table := wallet.GetDBModel("eth")
	var addrM = &model.CoinAddress{}
	addrM.Address = address
	cModel.GetData(table, addrM)
	if addrM.PrivKey == "" {
		return nil, errors.New("No such address in the database.")
	}

	s, _ := hex.DecodeString(addrM.PrivKey)
	path, _ := wallet.MakeBip44Path(wallet.ETH, 0, 0, 0)
	w, err := wallet.NewEthWalletByPath(path, s, wallet.EthChainId)
	if err != nil {
		return nil, err
	}

	return w.DeriveNativePrivateKey(), nil
}

func getData(toAddress string, amount *big.Int) []byte {

	// transferFnString := []byte("transfer(address,uint256)")
	// hash := sha3.NewLegacyKeccak256()
	// hash.Write(transferFnString)
	// methodId := hash.Sum(nil)[:4]

	methodId := []byte("0xa9059cbb")
	paddedToAddress := common.LeftPadBytes([]byte(toAddress), 32)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodId...)
	data = append(data, paddedToAddress...)
	data = append(data, paddedAmount...)

	return data
}
