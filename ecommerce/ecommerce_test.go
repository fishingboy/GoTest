package ecommerce

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrder_Total(t *testing.T) {
	customer := Customer{Name: "leo", Address: "大同路三段"}
	product1 := Product{Name: "Apple Watch", Price: 13000}
	product2 := Product{Name: "iPhone", Price: 20000}
	order := Order{Customer: customer}
	order.Add(product1)
	order.Add(product2)
	total := order.Total()
	assert.Equal(t, 33000, total)
}
