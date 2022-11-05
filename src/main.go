package main

import (
	"fmt"
	"github.com/icon-project/goloop/client"
	"github.com/icon-project/goloop/server/jsonrpc"
	"paulrouge/go-icon-sdk/networks"
	"paulrouge/go-icon-sdk/transactions"
	"paulrouge/go-icon-sdk/util"
	"paulrouge/go-icon-sdk/wallet"
	v3 "github.com/icon-project/goloop/server/v3"
)


func main() {
	fmt.Println("Hello, world!")
	
	// set the active network globally (this way we can reuse the network id in the tx builders)
	networks.SetActiveNetwork(networks.Lisbon())
	

	Client := client.NewClientV3(networks.GetActiveNetwork().URL)
	_ = Client

	Wallet := wallet.LoadWallet("../mywallets/keystore.json", "joejoe")

	_ = Wallet

	// bn := util.ICXToLoop(0.1)
	// txobject := transactions.TransferICXBuilder("hx9c13cd371aed69c79870b3a3f7492c10122f0315", bn)

	// tx, err := Client.SendTransaction(Wallet, txobject)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // tx to string
	// fmt.Println(string(*tx))

	// a := "cx33a937d7ab021eab50a7b729c4de9c10a77d51bd"
	// method := "getNFTPrice"
	// params := map[string]interface{}{
	// 	"_tokenId": "0x2",
	// }

	// callObject := transactions.CallBuilder(a,method, params)

	// response, err := Client.Call(callObject)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	
	// // hex := jsonrpc.HexInt(response.(string))
	// bn := util.HexToBigInt(response.(string))
	
	// fmt.Println(bn) // is BigInt with 18 decimals
	
	// declare an AddressParam
	var adr v3.AddressParam 
	
	// set the address to the .Address field
	adr.Address = jsonrpc.Address("hx9c13cd371aed69c79870b3a3f7492c10122f0315")
	
	// get the balance of the address
	balance, _ := Client.GetBalance(&adr)

	// print the balance using util.HexToBigInt()
	fmt.Println(util.HexToBigInt(string(*balance)))

}