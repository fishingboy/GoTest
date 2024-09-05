package ecommerce

type Order struct {
	Customer Customer
	Items    []Product
}

func (this *Order) Add(product Product) {
	this.Items = append(this.Items, product)
}

func (this *Order) Total() int {
	total := 0
	for _, item := range this.Items {
		total += item.Price
	}
	return total
}
