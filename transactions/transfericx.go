package transactions

import (
	"github.com/icon-project/goloop/server/jsonrpc"
	"github.com/icon-project/goloop/server/v3"
	"paulrouge/go-icon-sdk/networks"
	"math/big"
	"paulrouge/go-icon-sdk/util"
)

// amount is number of icx as a string
func TransferICXBuilder(to string, amount *big.Int) *v3.TransactionParam {
	// convert to to jsonrpc.Address
	toAddress := jsonrpc.Address(to)
	
	txParams := v3.TransactionParam{
		FromAddress: "hx9c13cd371aed69c79870b3a3f7492c10122f0315",
		ToAddress: toAddress,
		Value: util.BigIntToHex(amount),
		StepLimit: "0xf4240",
		NetworkID: networks.GetActiveNetwork().NID,
		Nonce: "0x1",
		Version: "0x3",
		Timestamp: "0x",
	}

	return &txParams
}