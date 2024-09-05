package tests

import (
	"fmt"
	"testing"
)

// Engine 定義引擎結構體
type Engine struct {
	Power int
}

// Car 定義汽車結構體，使用組合包含 Engine
type Car struct {
	Brand  string
	Engine // 包含 Engine 結構體
}

// Start 方法屬於 Car 結構體，但可以訪問 Engine 的屬性
func (c *Car) Start() {
	// 上層可以訪問下層的屬性耶，但反過來不行！！
	fmt.Printf("%s starts with %d horsepower\n", c.Brand, c.Power)
}

func TestStart(t *testing.T) {
	myCar := Car{
		Brand: "Toyota",
		Engine: Engine{
			Power: 250,
		},
	}

	myCar.Start() // 調用 Car 結構體的 Start 方法
}
