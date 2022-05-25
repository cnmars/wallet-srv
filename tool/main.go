package main

import (
	"flag"
	"fmt"
	"os"
	"wallet-srv/lib"
	"wallet-srv/lib/pkg/dot"
)

var cmdFlag = flag.NewFlagSet("cmd", flag.ExitOnError)

func main() {

	id := cmdFlag.Bool("id", false, "machine uniq id")
	ik := cmdFlag.Bool("aes-ik", false, "aes encrypt ik length 32 charset")
	iv := cmdFlag.Bool("aes-iv", false, "aes encrypt iv length 16 charset")
	xxkey := cmdFlag.Bool("xxkey", false, "xxteam private")
	dotMeta := cmdFlag.Bool("dot-meta", false, "get dot chain lastest meta")
	cFile := cmdFlag.String("c", "../conf/config.json", "config file")
	randBool := cmdFlag.Bool("randstr", false, "generate rand string")
	randLength := cmdFlag.Int("len", 32, "rand string length")

	cmdFlag.Parse(os.Args[1:])
	cmdFlag.Usage = usage

	switch true {
	case *id:
		machineID, err := lib.GetMachineID()
		if err != nil {
			fmt.Println("Get local machine ID fail.")
			return
		}
		fmt.Printf("MachineID：%s\n\n", machineID)
	case *ik:
		fmt.Printf("AES-IK: %s\n\n", lib.RandStr(32))
	case *iv:
		fmt.Printf("AES-IV: %s\n\n", lib.RandStr(16))
	case *xxkey:
		fmt.Printf("XXTEA-KEY: %s\n\n", lib.RandStr(40))
	case *dotMeta:
		config, err := lib.LoadConfig(*cFile)
		if err != nil {
			fmt.Println("config json file load failed. ../conf/config.json not exist.")
			return
		}
		dot.UpdateMeta(config)

		// meta, err := dot.GetMeta()
		// fmt.Println(meta, " ", err, *cFile)
	case *randBool:
		fmt.Printf("\nRandStr: %s\n\n", lib.RandStr(*randLength))
	default:
		cmdFlag.Usage()
	}

}

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: ./wallet-tool -id -aes-ik -aes-kv -xxkey -dot-meta -c ../conf/config.json

Options:
`)
	cmdFlag.PrintDefaults()

}
