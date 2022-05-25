package lib

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
	"wallet-srv/conf"
	"wallet-srv/lib/types"
)

//check auth access
func AuthMachineID() {
	machineID, err := GetMachineID()
	if err != nil || machineID != conf.MACHINE_UUID {
		fmt.Println("Wallet Service No Auth")
		os.Exit(0)
	}
}

//machine uniq id
func GetMachineID() (string, error) {

	mac, err := getMacAddress()
	if err != nil {
		return "", err
	}

	ip, err := getLocalIP()
	if err != nil {
		return "", err
	}

	return strings.ToUpper(Wmd5(fmt.Sprintf("%s%s", mac, ip))), nil
}

// get mac address
func getMacAddress() (string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		panic(err.Error())
	}
	mac, macerr := "", errors.New("system error")
	for i := 0; i < len(netInterfaces); i++ {
		//fmt.Println(netInterfaces[i])
		if (netInterfaces[i].Flags&net.FlagUp) != 0 && (netInterfaces[i].Flags&net.FlagLoopback) == 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				ipnet, ok := address.(*net.IPNet)
				//fmt.Println(ipnet.IP)
				if ok && ipnet.IP.IsGlobalUnicast() {

					mac = netInterfaces[i].HardwareAddr.String()
					return mac, nil
				}
			}
		}
	}
	return mac, macerr
}

//get local ip
func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", nil
	}

	var localIP string
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIP = ipnet.IP.String()
				break
			}

		}
	}
	return localIP, nil
}

//md5
func Wmd5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// rand string by length
func RandStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

//load conf/config.json configure file
func LoadConfig(filename string) (*types.Config, error) {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		// log.Error("Read config file error: ", zap.Error(err))
		return nil, err
	}
	// log.Info(string(content))

	var config types.Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		// log.Error("Unmarshal config file error: ", zap.Error(err))
		return nil, err
	}

	return &config, nil
}
