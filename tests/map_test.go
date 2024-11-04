package tests

import (
	"fmt"
	"testing"
)

// 定義結構 Item
type Item struct {
	ID    int
	Count int
}

func TestMap(t *testing.T) {
	// 初始化一個 map 來存儲 Item
	// 一定要宣告是指標，才有辦法改值
	idleItem := make(map[int]*Item)

	// 假設有一些假想的 ItemID 和數量
	itemIDs := []int{1, 2, 3}
	totalCounts := []int{5, 10, 8}

	// 將每個 ItemID 初始化到 map 中
	for i, id := range itemIDs {
		// 一定要宣告是指標，才有辦法改值
		idleItem[id] = &Item{ID: id, Count: totalCounts[i]}
	}

	// 模擬增加計數
	for id := range idleItem {
		// 一定要宣告是指標，才有辦法改值
		idleItem[id].Count++
	}

	// 打印結果
	for id, item := range idleItem {
		fmt.Printf("ItemID: %d, Count: %d\n", id, item.Count)
	}
}

func TestMapCannotAssign(t *testing.T) {
	// 初始化一個 map 來存儲 Item
	// 一定要宣告是指標，才有辦法改值
	idleItem := make(map[int]Item)

	// 假設有一些假想的 ItemID 和數量
	itemIDs := []int{1, 2, 3}
	totalCounts := []int{5, 10, 8}

	// 將每個 ItemID 初始化到 map 中
	for i, id := range itemIDs {
		// 一定要宣告是指標，才有辦法改值
		idleItem[id] = Item{ID: id, Count: totalCounts[i]}
	}

	// 模擬增加計數
	for id := range idleItem {
		// 一定要宣告是指標，才有辦法改值
		// 這樣的話就一定會出錯
		idleItem[id].Count++
	}

	// 打印結果
	for id, item := range idleItem {
		fmt.Printf("ItemID: %d, Count: %d\n", id, item.Count)
	}
}
