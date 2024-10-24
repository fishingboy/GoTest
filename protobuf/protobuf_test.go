package protobuf

import (
	"fmt"
	"github.com/yinweli/Mizugo/mizugos/cryptos"
	"testing"
)

func TestString(t *testing.T) {
	str := "abc123"
	str2 := []byte(str)
	fmt.Println(str, str2)
	fmt.Printf("%s %s\n", str, str2)
}

func TestBase64(t *testing.T) {
	coder := cryptos.NewBase64()
	str := []byte("abc123")
	response, err := coder.Encode(str)
	fmt.Printf("response: %s", response)
	if err != nil {
		fmt.Println("error:", err)
	}
}
