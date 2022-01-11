package sign

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_sign(t *testing.T) {
	userAddress := "0x0d50137433c1245673A934D818f9614A72831462"
	jikeID := "5e6bf0c9"
	salt := "4bc8ff1a135e6bf0c99"
	contractAddress := "0xFbFd732617D215252aDE934168e6bdEbD3059af3"
	// user + jikeID + salt + contract
	sig := getSignature(userAddress, jikeID, salt, contractAddress)
	sigStr := fmt.Sprintf("0x%x", sig)
	fmt.Println(fmt.Sprintf("signature: 0x%x, len: %d", sig, len(sig)))
	expectedSig := "0x8047a1e82af2321f1b6e5840126a27e8b965c0b80c855462f680cb32ff0912314ac56b7956b827ad02406aff2c8e53e4aeee54cc4d518af8792f76b3d40dfe971c"
	if !reflect.DeepEqual(expectedSig, sigStr) {
		t.Fatalf("Output not expected.\n\tExpected: %#v\n\tGot: %#v\n", expectedSig, sigStr)
	}
}
