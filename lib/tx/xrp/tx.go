package xrp

import (
	"encoding/hex"
	"strings"
	"wallet-srv/lib/pkg/btcec"
	"wallet-srv/lib/pkg/xrp/crypto"
	"wallet-srv/lib/pkg/xrp/data"

	"github.com/golang/glog"
)

type ecdsaKey struct {
	*btcec.PrivateKey
}

// Account is a Ripple account
type Account struct {
	AccountId    string   `json:"account_id"`
	PublicKey    string   `json:"public_key"`
	PrivateKey   string   `json:"private_key"`
	Secret       string   `json:"secret"`
	TxSpendLimit string   `json:"tx_spend_limit"`
	Whitelist    []string `json:"whitelist"`
	Blacklist    []string `json:"blacklist"`
}

func NewAccount(seed string) (*Account, error) {
	s, err := hex.DecodeString(seed)
	if err != nil {
		glog.Info("NewAccount hex.DecodeString err ", err)
		return nil, err
	}

	ecdsaKey, err := crypto.NewECDSAKey(s)
	if err != nil {
		glog.Info("NewAccount crypto.NewECDSAKey err ", err)
		return nil, err
	}

	keySequenceZero := uint32(0)

	publicKeyHash, err := crypto.AccountPublicKey(ecdsaKey, &keySequenceZero)
	if err != nil {
		glog.Info("NewAccount crypto.AccountPublicKey err ", err)
		return nil, err
	}

	privateKeyHash, err := crypto.AccountPrivateKey(ecdsaKey, &keySequenceZero)
	if err != nil {
		glog.Info("NewAccount crypto.AccountPrivateKey err ", err)
		return nil, err
	}

	// Create a Ripple account id from the public key
	accountIdHash, err := crypto.AccountId(ecdsaKey, &keySequenceZero)
	if err != nil {
		glog.Info("NewAccount crypto.AccountId err ", err)
		return nil, err
	}

	var blacklist []string
	var whitelist []string

	return &Account{
		AccountId:    accountIdHash.String(),
		PublicKey:    publicKeyHash.String(),
		PrivateKey:   privateKeyHash.String(),
		Secret:       seed,
		TxSpendLimit: "",
		Whitelist:    whitelist,
		Blacklist:    blacklist,
	}, nil
}

func NewTxSign(privKey, from, to, amount, assetCode, assetIssuer string, feeAmount int64, fromSeq *uint32) (*data.Payment, error) {

	tx, err := createPaymentTransaction(from, to, amount, assetCode, assetIssuer, feeAmount)
	if err != nil {
		glog.Info("NewTxSign createPaymentTransaction err: ", err)
		return nil, err
	}

	fromAccount, err := NewAccount(privKey)
	if err != nil {
		glog.Info("NewTxSign NewAccount err: ", err)
		return nil, err
	}
	txSign, err := signPaymentTransaction(fromAccount, tx, fromSeq)
	if err != nil {
		glog.Info("NewTxSign signPaymentTransaction err: ", err)
		return nil, err
	}

	return txSign, nil
}

func createPaymentTransaction(from string, to string, amount string, assetCode string, assetIssuer string, feeAmount int64) (*data.Payment, error) {
	src, err := data.NewAccountFromAddress(from)
	if err != nil {
		glog.Info("createPaymentTransaction err ", err)
		return nil, err
	}

	dest, err := data.NewAccountFromAddress(to)
	if err != nil {
		glog.Info("createPaymentTransaction err ", err)
		return nil, err
	}

	// Convert the amount into an object
	var amountObj *data.Amount
	if strings.EqualFold(assetCode, "native") {
		amountObj, err = data.NewAmount(amount + "/XRP")
		if err != nil {
			glog.Info("createPaymentTransaction err ", err)
			return nil, err
		}
	} else {
		amountObj, err = data.NewAmount(amount + "/" + assetCode + "/" + assetIssuer)
		if err != nil {
			glog.Info("createPaymentTransaction err ", err)
			return nil, err
		}
	}

	// Create payment
	payment := &data.Payment{
		Destination: *dest,
		Amount:      *amountObj,
	}
	payment.TransactionType = data.PAYMENT

	payment.Flags = new(data.TransactionFlag)

	fee, err := data.NewNativeValue(feeAmount)
	base := payment.GetBase()
	base.Fee = *fee
	//base.Memos = new data.Memo{}
	base.Account = *src

	return payment, nil
}

func signPaymentTransaction(account *Account, paymentTx *data.Payment, sequence *uint32) (*data.Payment, error) {
	seed, _ := hex.DecodeString(account.Secret)
	key, err := crypto.NewECDSAKey(seed)

	base := paymentTx.GetBase()
	base.Sequence = *sequence

	// Sign the payment transaction
	keySequence := uint32(0)
	err = data.Sign(paymentTx, key, &keySequence)
	if err != nil {
		return nil, err
	}

	//verfiy sign
	if ok, err := data.CheckSignature(paymentTx); !ok {
		return nil, err
	}

	return paymentTx, nil
}
