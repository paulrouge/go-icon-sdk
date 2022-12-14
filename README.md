# go-icon-sdk

The Icon SDK for Go is a library for building applications on the ICON network.

## Create Client
In src/main.go in the main function:

1. Set the node you want to connect to globally. You can add networks in the networks/networks.go file.
```go
// Lisbon Testnet
networks.SetActiveNetwork(networks.Lisbon())

// Mainnet
networks.SetActiveNetwork(networks.Mainnet())
```

2. Create client
```go
Client := client.NewClientV3(networks.GetActiveNetwork().URL)
```

We can now call several functions on the client. For example, we can get the balance of an address like this:

```go
// declare an AddressParam
var adr v3.AddressParam 

// set the address to the .Address field
adr.Address = jsonrpc.Address("hx9c13cd371aed69c79870b3a3f7492c10122f0315")

// get the balance of the address
balance, _ := Client.GetBalance(&adr)

// print the balance using util.HexToBigInt()
fmt.Println(util.HexToBigInt(string(*balance)))
```

[Click here to see all the available methods on the created Client](https://pkg.go.dev/github.com/icon-project/goloop@v1.2.14/client#NewClientV3)


## Create Wallet
When creating a new wallet it is automatically __saved as a keystore file.__ Call the function below with the _"path/filename"_. The password is used to encrypt the keystore file, don't forget it!

```go
wallet.CreateNewWalletAndKeystore("../mywallets/keystore01", "password")
```

## Load Wallet
When loading a wallet you need to provide the path to the keystore file and the password to decrypt the keystore file.

```go
Wallet := wallet.LoadWalletFromKeystore("../mywallets/keystore01", "password")
```
__Note:__ To prevent confusing between the created wallet instance and the wallet-package we name the wallet that we load "Wallet" (so with a capital W, instead of the package name).

## Send ICX
Use the TransferICXBuilder to get a transaction object. The address should be a string and the amount must be converted to a big.Int before sending it to the builder. We do this by using the "util.ICXToLoop()" function.


```go
// set address & amount of ICX to sent
address := "hx0000000000000000000000000000000000000000" // must be a string
amount := 1 // can also be a string "1" or a float 1.0

// convert amount of icx to loop in big.Int
bn := util.ICXToLoop(amount)

// create transaction object
txobject := transactions.TransferICXBuilder(address, bn)

// we need to have a wallet loaded to sign the transaction
Wallet := wallet.LoadWallet("../mywallets/keystore01", "password")

// sign & send the transaction
tx, err := Client.SendTransaction(Wallet, txobject)
if err != nil {
    fmt.Println(err)
}

// print the transaction hash
fmt.Println(*tx)
```

## Call a Smart Contract on the ICON Blockchain (read-only)
Use the CallBuilder to get a call-object. The Callbuilder takes in the address of the smart contract as a string, the name of the method you want to call (also as a string) and a params object. If the method you want to call does not take any parameters you can just pass in a empty object.

1. Call a method with no parameters

```go
// set address
contractAddress := "cx33a937d7ab021eab50a7b729c4de9c10a77d51bd"

// set the method to call (there is a method on this contract called "name")
method := "name" 

// create call object with params as nil
callObject := transactions.CallBuilder(contractAddress, method, nil)

// make the call
response, err := Client.Call(callObject) 
if err != nil {
    fmt.Println(err)
}

// print the response
fmt.Println(response) // "Art Gallery"
```

2. Call a method with parameters

```go
// set address
contractAddress := "cx33a937d7ab021eab50a7b729c4de9c10a77d51bd"

// this is the method takes in a parameter
method := "getNFTPrice" 

// the parameter _tokenId is set to 0x2
params := map[string]interface{}{
    "_tokenId": "0x2", 
}

// create call object
callObject := transactions.CallBuilder(contractAddress, method, params)

// make the call
response, err := Client.Call(callObject)
if err != nil {
    fmt.Println(err)
}

// the response is a string, we need to convert it to a hex
hex := jsonrpc.HexInt(response.(string))

// and then convert it to a bigInt
bn := util.HexToBigInt(hex)

// and finally print it
fmt.Println(bn) 
```

## Change a value in a Smart Contract on the ICON Blockchain
When you want to change a value on a smart contract you need to use the "SendTransaction" function. This function takes in a wallet, a transaction object and a stepLimit. The stepLimit is the maximum amount of steps that the transaction can use. The stepLimit is calculated by the ICON network and is returned in the response of the transaction. If you want to be sure that your transaction is executed you can set the stepLimit to a very high number. 

Here we first call the current value of the 'name' variable on the contract, and then change it.

```go
// set the contract address
contractAddress := "cx2b60e6e094df34a0d7c05b5ff5cb6758aba7e83e"

// this address has a method called name that returns the current "name" value of the contract
method := "name"

// we only read the contract, so we don't need to sign the tx and can use the CallBuilder
callObject := transactions.CallBuilder(contractAddress, method, nil)

// send the call
res, _ := Client.Call(callObject)

fmt.Println(res) // Returns the current value of 'name' on the contract.

//////////////////////////////// NOW WE WILL CHANGE THE VALUE ///////////////////////////////////

// the method we want to call is called "setName"
method = "setName"

// the params for the method,
params := map[string]interface{}{
    "name": "Satoshi",
}

// this transaction / method call does not require payment so we can set the value to 0,
value := util.HexToBigInt("0x0")

// We need to sign the tx, so we use the TransactionBuilder. 
tx := transactions.TransactionBuilder(Wallet.Address(), contractAddress, method, params, value)

// sign the tx
hash, err := Client.SendTransaction(Wallet, tx)
if err != nil {
    fmt.Println(err)
}

fmt.Println(*hash) // Returns the hash of the tx.
```
Run the first part of the code again or [check the contract on the tracker](https://lisbon.tracker.solidwallet.io/contract/cx2b60e6e094df34a0d7c05b5ff5cb6758aba7e83e#readcontract) to see if the value has changed.