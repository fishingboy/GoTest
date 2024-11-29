package mix

import (
	"fmt"
)

type Person struct {
	age int32
	car *Car
}

type Car struct {
	speed int32
}

func (this *Car) run() {
	fmt.Println("run!!!")
}
