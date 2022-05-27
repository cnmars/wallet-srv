package flow

import (
	"github.com/golang/glog"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
)

//FlowToken transfer script
const TOKEN_TRANSFER_CADENCE_SCRIPT = `
import FungibleToken from 0xf233dcee88fe0abe
import FlowToken from 0x1654653399040a61

transaction(amount: UFix64, to: Address) {

    let sentVault: @FungibleToken.Vault

    prepare(signer: AuthAccount) {
        let vaultRef = signer.borrow<&FlowToken.Vault>(from: /storage/flowTokenVault)
			?? panic("Could not borrow reference to the owner''s Vault!")

        self.sentVault <- vaultRef.withdraw(amount: amount)
    }

    execute {
        let receiverRef =  getAccount(to)
            .getCapability(/public/flowTokenReceiver)
            .borrow<&{FungibleToken.Receiver}>()
			?? panic("Could not borrow receiver reference to the recipient''s Vault")

        receiverRef.deposit(from: <-self.sentVault)
    }
}`

type Tx struct {
	Signer flow.Address
	Sender flow.Address
	Player flow.Address
	Tx     *flow.Transaction
}

func NewTx() *Tx {
	tx := *flow.NewTransaction()
	return &Tx{
		Signer: flow.Address{},
		Sender: flow.Address{},
		Player: flow.Address{},
		Tx:     &tx,
	}
}

func (tx *Tx) SetSender(addr flow.Address) {
	tx.Sender = addr
}

func (tx *Tx) SetParam(amt string, to string) {
	amount, err := cadence.NewUFix64(amt)
	if err != nil {
		glog.Error("SetParam amt error. ", err)
		return
	}

	recipient := cadence.NewAddress(flow.HexToAddress(to))
	err = tx.Tx.AddArgument(amount)
	if err != nil {
		glog.Error("AddArgument to error. ", err)
		return
	}

	err = tx.Tx.AddArgument(recipient)
	if err != nil {
		glog.Error("AddArgument recipient error. ", err)
		return
	}
}

func (tx *Tx) SetTransferScript() {
	tx.Tx.SetScript([]byte(TOKEN_TRANSFER_CADENCE_SCRIPT))
}

func (tx *Tx) SetFeeLimit(fee uint64) {
	tx.Tx.SetGasLimit(fee)
}

func (tx *Tx) SetPlayer(play flow.Address) {
	tx.Tx.SetPayer(play)
}

func (tx *Tx) SetAuthorizer(addr flow.Address) {
	tx.Tx.AddAuthorizer(addr)
}

func (tx *Tx) SetProposalKey(addr flow.Address, keyIndex int, seqNum uint64) {
	tx.Tx.SetProposalKey(addr, keyIndex, seqNum)
}

func (tx *Tx) SetBlockHash(blockHash string) {
	tx.Tx.SetReferenceBlockID(flow.HexToID(blockHash))
}

func (tx *Tx) Sign(senderPrivKey crypto.PrivateKey, playerPrivKey crypto.PrivateKey) error {
	senderKey := flow.NewAccountKey().
		SetPublicKey(senderPrivKey.PublicKey()).
		SetSigAlgo(senderPrivKey.Algorithm()).
		SetHashAlgo(crypto.SHA3_256).
		SetWeight(flow.AccountKeyWeightThreshold)

	sendSigner, err := crypto.NewInMemorySigner(senderPrivKey, senderKey.HashAlgo)
	if err != nil {
		return err
	}

	playerKey := flow.NewAccountKey().
		SetPublicKey(playerPrivKey.PublicKey()).
		SetSigAlgo(playerPrivKey.Algorithm()).
		SetHashAlgo(crypto.SHA3_256).
		SetWeight(flow.AccountKeyWeightThreshold)

	playerSigner, err := crypto.NewInMemorySigner(playerPrivKey, playerKey.HashAlgo)
	if err != nil {
		return err
	}

	//sender sign SignPayload
	err = tx.Tx.SignPayload(tx.Sender, 0, sendSigner)
	if err != nil {
		return err
	}

	//player sign SignEnvelope
	err = tx.Tx.SignEnvelope(tx.Player, 0, playerSigner)
	if err != nil {
		return err
	}

	return nil
}

/**

// Replace with script above
const transferScript string = TOKEN_TRANSFER_CADENCE_SCRIPT

var (
    senderAddress    flow.Address
    senderAccountKey flow.AccountKey
    senderPrivateKey crypto.PrivateKey
)

func main() {
    tx := flow.NewTransaction().
        SetScript([]byte(transferScript)).
        SetGasLimit(100).
        SetPayer(senderAddress).
        SetAuthorizer(senderAddress).
        SetProposalKey(senderAddress, senderAccountKey.Index, senderAccountKey.SequenceNumber)

    amount, err := cadence.NewUFix64("123.4")
    if err != nil {
        panic(err)
    }

    recipient := cadence.NewAddress(flow.HexToAddress("0xabc..."))

    err = tx.AddArgument(amount)
    if err != nil {
        panic(err)
    }

    err = tx.AddArgument(recipient)
    if err != nil {
        panic(err)
    }
}


single party, single signuature

	referenceBlockID := examples.GetReferenceBlockId(flowClient)
	tx := flow.NewTransaction().
		SetScript([]byte(`
            transaction {
                prepare(signer: AuthAccount) { log(signer.address) }
            }
        `)).
		SetGasLimit(100).
		SetProposalKey(account1.Address, account1.Keys[0].Index, account1.Keys[0].SequenceNumber).
		SetReferenceBlockID(referenceBlockID).
		SetPayer(account1.Address).
		AddAuthorizer(account1.Address)

	// account 1 signs the envelope with key 1
	err = tx.SignEnvelope(account1.Address, account1.Keys[0].Index, key1Signer)

*/
