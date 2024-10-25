package tests

import (
	"fmt"
	"testing"
)

type Dog[Q any] struct {
	weight  int32
	name    string
	request *Q
}

type x struct {
	id int32
}

type y struct {
	id float32
}

func TestDog(t *testing.T) {
	dog1 := Dog[x]{
		weight: 100,
		name:   "lucky",
	}
	dog2 := Dog[y]{
		weight: 200,
		name:   "sun",
	}
	fmt.Println(dog1, dog2)
}
