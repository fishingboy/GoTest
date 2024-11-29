package tests

import (
	"fmt"
	"sort"
	"testing"
)

type Boy struct {
	Age  int32
	Name string
}

func TestSort(t *testing.T) {
	arr := []Boy{
		{Age: 1, Name: "CCC"},
		{Age: 100, Name: "AAA"},
		{Age: 50, Name: "BBB"},
		{Age: 10, Name: "XXX"},
		{Age: 990, Name: "ZZZ"},
	}
	fmt.Println(arr)

	sort.Slice(arr, func(l, r int) bool {
		// 如果判斷式成立，就不進行交換
		// 所以如果判斷試是 < ，就代表從小到大排序
		// 所以如果判斷試是 > ，就代表從大到小排序
		return arr[l].Age < arr[r].Age
	})

	fmt.Println(arr)

	sort.Slice(arr, func(l, r int) bool {
		// 如果判斷式成立，就不進行交換
		// 所以如果判斷試是 < ，就代表從小到大排序
		// 所以如果判斷試是 > ，就代表從大到小排序
		return arr[l].Name > arr[r].Name
	})

	fmt.Println(arr)
}
