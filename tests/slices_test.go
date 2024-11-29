package tests

import (
	"fmt"
	"testing"
)

func TestSlices(t *testing.T) {
	arr := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	start := 1
	size := 5
	fmt.Println(arr[start : start+size])
}
