package tests

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringBuilder(t *testing.T) {
	builder := &strings.Builder{}
	builder.Write([]byte{48})
	builder.WriteString("1")
	builder.WriteString("2")
	builder.WriteString("3")
	fmt.Println(builder)
}
