package commend

import (
	"fmt"
	"github.com/fishingboy/GoTest/Other"
)

func Scope() {
	fmt.Println("here is scope.go")
	Version()
	other.Other()
	fmt.Println(VERSION)
	fmt.Println(name)
}
