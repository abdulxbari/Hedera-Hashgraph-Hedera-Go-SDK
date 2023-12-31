package main

import (
	"fmt"
	"os"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

func main() {
	// client := hedera.ClientForTestnet()
	client := hedera.ClientForPreviewnet()
	myAccountId, err := hedera.AccountIDFromString(os.Getenv("OPERATOR_ID"))
	if err != nil {
		panic(err)
	}

	myPrivateKey, err := hedera.PrivateKeyFromString(os.Getenv("OPERATOR_KEY"))
	if err != nil {
		panic(err)
	}

	client.SetOperator(myAccountId, myPrivateKey)

	// ## Example
	// Create a ECDSA private key
	// Extract the ECDSA public key public key
	// Extract the Ethereum public address
	// Use the `AccountCreateTransaction` and populate `setAlias(evmAddress)` field with the Ethereum public address
	// Sign the `AccountCreateTransaction` transaction with the new private key
	// Get the `AccountInfo` on the new account and show that the account has contractAccountId

	// Create a ECDSA private key
	privateKey, err := hedera.PrivateKeyGenerateEcdsa()
	if err != nil {
		println(err.Error())
		return
	}
	// Extract the ECDSA public key public key
	publicKey := privateKey.PublicKey()
	// Extract the Ethereum public address
	evmAddress := publicKey.ToEvmAddress()

	// Use the `AccountCreateTransaction` and set the EVM address field to the Ethereum public address
	response, err := hedera.NewAccountCreateTransaction().SetInitialBalance(hedera.HbarFromTinybar(100)).
		SetKey(myPrivateKey).SetAlias(evmAddress).Sign(privateKey).Execute(client)
	if err != nil {
		println(err.Error())
		return
	}

	transactionReceipt, err := response.GetReceipt(client)
	if err != nil {
		println(err.Error(), ": error getting receipt}")
		return
	}

	newAccountId := *transactionReceipt.AccountID

	// Get the `AccountInfo` on the new account and show that the account has contractAccountId
	info, err := hedera.NewAccountInfoQuery().SetAccountID(newAccountId).Execute(client)
	if err != nil {
		println(err.Error())
		return
	}
	// Verify account is created with the provided EVM address
	fmt.Println(info.ContractAccountID == evmAddress)
	// Verify the account Id is the same from the create account transaction
	fmt.Println(info.AccountID.String() == newAccountId.String())
}
