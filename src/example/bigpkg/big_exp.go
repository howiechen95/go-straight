package bigpkg

import (
	"encoding/base64"
	"fmt"
	"math/big"
)

func NewFromInt64(i int64) {
	num := big.NewInt(i)
	fmt.Printf("Big int: %v\n", num)
}

func NewFromBytes(b []byte) {
	z := new(big.Int)
	z.SetBytes(b)
	fmt.Printf("Big int: %v\n", z)
}

func Base64ToInt(s string) (*big.Int, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	i := new(big.Int)
	i.SetBytes(data)
	return i, nil
}

func NewFromString(s string, base int) *big.Int {
	n := new(big.Int)
	n, ok := n.SetString(s, base)
	if !ok {
		fmt.Println("SetString: error")
		return nil
	}
	//fmt.Println(n)
	return n
}
