package mix

import (
	"testing"
)

func TestStart(t *testing.T) {
	person := Person{
		age: 11,
		car: &Car{speed: 100},
	}
	person.car.run()
}
