package stringpkg

import (
	"fmt"
	"testing"
)

func TestStr2Int(t *testing.T) {
	fmt.Println(Str2Int("123"))
	fmt.Println(Str2Int("123.324"))
}
