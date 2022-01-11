package main

import (
	"fmt"
	"github.com/NaturalSelectionLabs/bridge-utils/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func getSignature() []byte {
	kp, _ := secp256k1.NewKeypairFromString("b0897a82b80437060eb73147b77cba0c5f2ba182637f2be083784d0a111fc6b3")

	// hash
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"string"},
		//[]string{"string", "uint256", "address", "address", "uint32"},

		// values
		[]interface{}{
			"withdrawERC20",
		},
	)

	hashEth := solsha3.SoliditySHA3WithPrefix(hash)
	fmt.Println("hash: ", fmt.Sprintf("0x%x", hash))
	fmt.Println("hashEth: ", fmt.Sprintf("0x%x", hashEth))

	sig, err := crypto.Sign(hashEth, kp.PrivateKey())
	if err != nil {
		fmt.Println("Failed to generate withdraw signature", "err", err)
	}
	return sig
}

func main() {
	sig := getSignature()
	fmt.Println(fmt.Sprintf("sig: 0x%x", sig))
}
