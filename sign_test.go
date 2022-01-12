package sign

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

func Test_sign(t *testing.T) {
	userAddress := "0x170A213b5E8c95D3A8ae4C11569B60c2Ff782ab5"
	jikeID := "5e6bf0c9"
	salt := big.NewInt(123456789)
	contractAddress := "0xBa120cE9dB8BC5aD7bf5cfAbaf133A6F9423CFba"
	// user + jikeID + salt + contract
	sig := fmt.Sprintf("0x%x", getSignature(userAddress, jikeID, salt, contractAddress))
	fmt.Println(fmt.Sprintf("userAddress: %s, jikeId: %s, salt: %d, signature: %s", userAddress, jikeID, salt, sig))
	expectedSig := "0x3968601436516991a0aa9a46f9484294bcffee4d8228d9b90660dee64dfed3c27b0634885be4f4b07a975085f9996dcbeb13365ef73412203f75447d8dbd480d1c"
	if !reflect.DeepEqual(expectedSig, sig) {
		t.Fatalf("Output not expected.\n\tExpected: %#v\n\tGot: %#v\n", expectedSig, sig)
	}
}
