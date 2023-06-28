package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type TransLogV1 struct {
}

type RecordPaymentRequest struct {
}

type Fee struct {
	Amount int64
}

func main() {
	str := "1hello world"
	hash := md5.Sum([]byte(str))
	hashString := hex.EncodeToString(hash[:])
	fmt.Println(hashString)

	for _, f := range []func(v1 TransLogV1, v2 *RecordPaymentRequest) error{} {
		fmt.Println(f)
	}

	var fees []*Fee
	for _, f := range fees {
		fmt.Println(f)
	}

}

// 5eb63bbbe01eeed093cb22bb8f5acdc3
// 5fcc6da514a1d1fc0656ac2ece36d8e7
// 118e57b1eecd61e54b502055d82aa57f
