package sign

import (
	"fmt"
	"github.com/NaturalSelectionLabs/bridge-utils/crypto/secp256k1"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func getSignature(userAddr string, id string, salt string, contractAddr string) []byte {
	kp, _ := secp256k1.NewKeypairFromString("14653d5667a235dd5acc4db4ececc1c734b775cf2753ec7b3d2e39921ddc5a6f")

	// hash
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"address", "string", "string", "address"},
		// values
		[]interface{}{
			ethcommon.HexToAddress(userAddr),
			id,
			salt,
			ethcommon.HexToAddress(contractAddr),
		},
	)

	hashEth := solsha3.SoliditySHA3WithPrefix(hash)
	fmt.Println("hash: ", fmt.Sprintf("0x%x", hash))
	//fmt.Println("hashEth: ", fmt.Sprintf("0x%x", hashEth))

	sig, err := crypto.Sign(hashEth, kp.PrivateKey())
	if err != nil {
		fmt.Println("Failed to generate withdraw signature", "err", err)
	}
	// v + 27
	sig[64] = sig[64] + 27
	return sig
}
