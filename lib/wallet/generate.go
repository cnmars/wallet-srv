package wallet

import (
	"encoding/hex"
	"fmt"
)

// BTC
func GetBtcAddress() (string, string, error) {
	hdw, _ := NewHDWallet(SymbolBtc, BTC)

	w, err := hdw.NewWallet(0, 0, 0)
	if err != nil {
		return "", "", err
	}

	return hex.EncodeToString(hdw.seed), w.DeriveAddress(), nil
}

// BCH
func GetBchAddress() (string, string, error) {
	hd, _ := NewHDWallet(SymbolBch, BCH)

	w, err := hd.NewWallet(0, 0, 0)
	if err != nil {
		return "", "", err
	}

	address, err := GetBitCoinCashAddress(w.DeriveAddress())
	if err != nil {
		return "", "", err
	}

	return hex.EncodeToString(hd.seed), address, nil
}

// BSV
func GetBsvAddress() (string, string, error) {
	hd, _ := NewHDWallet(SymbolBSV, BSV)

	w, err := hd.NewWallet(0, 0, 0)
	if err != nil {
		return "", "", err
	}

	return hex.EncodeToString(hd.seed), w.DeriveAddress(), nil
}

// LTC
func GetLtcAddress() (string, string, error) {
	hdw, _ := NewHDWallet(SymbolLtc, LTC)

	w, err := hdw.NewWallet(0, 0, 0)
	if err != nil {
		return "", "", err
	}

	return hex.EncodeToString(hdw.seed), w.DeriveAddress(), nil
}

// DASH
func GetDashAddress() (string, string, error) {
	hdw, _ := NewHDWallet(SymbolDash, DASH)

	w, err := hdw.NewWallet(0, 0, 0)
	if err != nil {
		return "", "", err
	}

	return hex.EncodeToString(hdw.seed), w.DeriveAddress(), nil
}

// DOGE
func GetDogeAddress() (string, string, error) {
	hdw, _ := NewHDWallet(SymbolDoge, DOGE)

	w, err := hdw.NewWallet(0, 0, 0)
	if err != nil {
		return "", "", err
	}

	return hex.EncodeToString(hdw.seed), w.DeriveAddress(), nil
}

// QTUM
func GetQtumAddress() (string, string, error) {
	hdw, _ := NewHDWallet(SymbolQTUM, QTUM)

	w, err := hdw.NewWallet(0, 0, 0)
	if err != nil {
		return "", "", err
	}

	return hex.EncodeToString(hdw.seed), w.DeriveAddress(), nil
}

//ETH BSC HT Matic ETC
func GetEthAddress() (string, string, error) {
	hdw, _ := NewHDWallet(SymbolEth, ETH)

	w, err := hdw.NewWallet(0, 0, 0)
	if err != nil {
		return "", "", err
	}

	return hex.EncodeToString(hdw.seed), w.DeriveAddress(), nil
}

//TRX
func GetTrxAddress() (string, string, error) {
	hdw, _ := NewHDWallet(SymbolTrx, TRX)

	w, err := hdw.NewWallet(0, 0, 0)
	if err != nil {
		return "", "", err
	}

	return hex.EncodeToString(hdw.seed), w.DeriveAddress(), nil
}

//SOL
func GetSolAddress() (string, string, error) {

	hdw, _ := NewHDWallet(SymbolSol, SOL)

	w, err := hdw.NewWallet(0, 0, 0)
	if err != nil {
		return "", "", err
	}

	return hex.EncodeToString(hdw.seed), w.DeriveAddress(), nil
}

//FIL
func GetFilAddress() (string, string, error) {

	hdw, _ := NewHDWallet(SymbolFil, FIL)
	w, err := hdw.NewWallet(0, 0, 0)
	if err != nil {
		return "", "", err
	}

	addr := w.DeriveAddress()
	if addr == "" {
		return "", "", fmt.Errorf("address create failed.")
	}

	return hex.EncodeToString(hdw.seed), w.DeriveAddress(), nil
}

//Ada
func GetAdaAddress() (string, string, error) {

	hdw, _ := NewHDWallet(SymbolAda, ADA)
	w, err := hdw.NewWallet(0, 0, 0)
	if err != nil {
		return "", "", err
	}

	addr := w.DeriveAddress()
	if addr == "" {
		return "", "", fmt.Errorf("address create failed.")
	}

	return hex.EncodeToString(hdw.seed), w.DeriveAddress(), nil
}

//XRP
func GetXrpAddress() (string, string, error) {

	hdw, _ := NewHDWallet(SymbolXrp, XRP)
	w, err := hdw.NewWallet(0, 0, 0)
	if err != nil {
		return "", "", err
	}

	addr := w.DeriveAddress()
	if addr == "" {
		return "", "", fmt.Errorf("address create failed.")
	}

	return hex.EncodeToString(hdw.seed), w.DeriveAddress(), nil
}

//LUNA
func GetLunaAddress() (string, string, error) {

	hdw, _ := NewHDWallet(SymbolLuna, LUNA)
	w, err := hdw.NewWallet(0, 0, 0)
	if err != nil {
		return "", "", err
	}

	addr := w.DeriveAddress()
	if addr == "" {
		return "", "", fmt.Errorf("address create failed.")
	}

	return hex.EncodeToString(hdw.seed), w.DeriveAddress(), nil
}

//XMR
func GetXmrAddress() (string, string, error) {

	hdw, _ := NewHDWallet(SymbolXmr, XMR)
	w, err := hdw.NewWallet(0, 0, 0)
	if err != nil {
		return "", "", err
	}

	addr := w.DeriveAddress()
	if addr == "" {
		return "", "", fmt.Errorf("address create failed.")
	}

	return hex.EncodeToString(hdw.seed), w.DeriveAddress(), nil
}

//NEAR
func GetNearAddress() (string, string, error) {

	hdw, _ := NewHDWallet(SymbolNear, NEAR)
	w, err := hdw.NewWallet(0, 0, 0)
	if err != nil {
		return "", "", err
	}

	addr := w.DeriveAddress()
	if addr == "" {
		return "", "", fmt.Errorf("address create failed.")
	}

	return hex.EncodeToString(hdw.seed), w.DeriveAddress(), nil
}

//DOT
func GetDotAddress() (string, string, error) {

	hdw, _ := NewHDWallet(SymbolDot, DOT)
	w, err := hdw.NewWallet(0, 0, 0)
	if err != nil {
		return "", "", err
	}

	addr := w.DeriveAddress()
	if addr == "" {
		return "", "", fmt.Errorf("address create failed.")
	}

	return hex.EncodeToString(hdw.seed), w.DeriveAddress(), nil
}

//AVAX
func GetAvaxAddress() (string, string, error) {

	hdw, _ := NewHDWallet(SymbolAvax, AVAX)
	w, err := hdw.NewWallet(0, 0, 0)
	if err != nil {
		return "", "", err
	}

	addr := w.DeriveAddress()
	if addr == "" {
		return "", "", fmt.Errorf("address create failed.")
	}

	return hex.EncodeToString(hdw.seed), w.DeriveAddress(), nil
}

//Flow
func GetFlowAddress() (string, string, error) {

	hdw, _ := NewHDWallet(SymbolFlow, FLOW)
	w, err := hdw.NewWallet(0, 0, 0)
	if err != nil {
		return "", "", err
	}

	addr := w.DeriveAddress()
	if addr == "" {
		return "", "", fmt.Errorf("address create failed.")
	}

	return hex.EncodeToString(hdw.seed), w.DeriveAddress(), nil
}
