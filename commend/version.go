package commend

import (
	"fmt"
)

const VERSION = "1.0.0"

func Version() {
	fmt.Println(GetVersion())
}

func GetVersion() string {
	return VERSION
}
