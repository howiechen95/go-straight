package bigpkg

import (
	"fmt"
	"math"
	"testing"
)

func TestNewFromInt64(t *testing.T) {
	NewFromInt64(1999)
	NewFromInt64(math.MaxInt64)
}

func TestNewFromBytes(t *testing.T) {
	str := "12345678"
	NewFromBytes([]byte(str))
}

func TestNewFromString(t *testing.T) {
	fmt.Println(NewFromString("12345678901234567890", 10))
	fmt.Println(NewFromString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16))
}
