package ch0

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestArray(t *testing.T) {
	list1 := make([]int64, 0)
	list2 := make([]int64, 10)
	s := 10
	for i := 0; i < s; i++ {
		n := rand.Int63n(5000 * 10000)
		list1 = append(list1, n)
		list2 = append(list2, n)
	}
	fmt.Println(list1)
	fmt.Println(list2)
}

type SDKVer string

func TestJWT(t *testing.T) {
	v1 := SDKVer("1.0.0")
	v2 := SDKVer("1.0.0")
	t.Log(v1 == v2)
}
