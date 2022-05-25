package main

import (
	"flag"
	"fmt"
	"os"
	"wallet-srv/lib"
	"wallet-srv/lib/wallet"
	"wallet-srv/model"

	"github.com/golang/glog"
)

var (
	AddrGenerateMap = map[string]func() (string, string, error){
		"btc":  wallet.GetBtcAddress,
		"bch":  wallet.GetBchAddress,
		"doge": wallet.GetDogeAddress,
		"dash": wallet.GetDashAddress,
		"ltc":  wallet.GetLtcAddress,
		"qtum": wallet.GetQtumAddress,
		"eth":  wallet.GetEthAddress,
		"trx":  wallet.GetTrxAddress,
		"sol":  wallet.GetSolAddress,
		"fil":  wallet.GetFilAddress,
		"ada":  wallet.GetAdaAddress,
		"xrp":  wallet.GetXrpAddress,
		"luna": wallet.GetLunaAddress,
		"near": wallet.GetNearAddress,
		"dot":  wallet.GetDotAddress,
		"bsv":  wallet.GetBsvAddress,
		"avax": wallet.GetAvaxAddress,
	}
)

var cmdFlag = flag.NewFlagSet("cmd", flag.ExitOnError)

func init() {
	flag.Lookup("alsologtostderr").Value.Set("true")
	flag.Lookup("log_dir").Value.Set("../logs/address")
	flag.Lookup("v").Value.Set("0")
}

func main() {

	defer glog.Flush()

	lib.AuthMachineID()

	num := cmdFlag.Int("num", 10, "generate address numbers.")
	loop := cmdFlag.Bool("loop", false, "loop create address")
	coin := cmdFlag.String("coin", "", "-coin=btc|bch|bsv|dash|doge|ltc|qtum|eth|trx|sol|fil|ada|xrp|luna|near|dot|avax")
	cmdFlag.Parse(os.Args[1:])
	cmdFlag.Usage = usage

	if *coin == "" || AddrGenerateMap[*coin] == nil {
		cmdFlag.Usage()
		return
	}

	cModel, table := wallet.GetDBModel(*coin)
	if *loop {
		freeNum := cModel.CountData(table, 0)
		if freeNum > int64(*num) {
			glog.Infof("[%s] free address total: %d", *coin, freeNum)
			return
		}
	}

	for i := 0; i < *num; i++ {
		//address
		privKey, address, err := AddrGenerateMap[*coin]()
		if err != nil {
			glog.Errorf("[%s] create address failed. %v", *coin, err)
			continue
		}

		//seed and address insert db table
		//glog.Info("---", address, " ", privKey, " ", err)
		var data = &model.CoinAddress{Address: address, PrivKey: privKey, IsUsed: 0, Multi: 1}

		//secord encrypt privkey
		err = model.DBEncrypt(data)
		if err != nil {
			glog.Error(err)
			continue
		}

		cModel.AddData(table, data)
	}

	glog.Infof("[%s] %d address generate finish.", *coin, *num)

}

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: ./wallet -num 100 -coin btc

Options:
`)
	cmdFlag.PrintDefaults()
}
