package bufiopkg

import (
	"fmt"
	"os"
	"testing"
)

func TestReadFrom(t *testing.T) {
	fmt.Println("请输入不多于9个字符，以回车结束：")
	data, err := ReadFrom(os.Stdin, 11)
	if err != nil {
		fmt.Println(data)
	}
}
