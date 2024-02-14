package main

import "fmt"

type item struct {
	productID int
	qtde      int
	price     float64
}

type order struct {
	userID int
	itens  []item
}

func (o order) totalValue() float64 {
	total := 0.0
	for _, item := range o.itens {
		total += item.price * float64(item.qtde)
	}
	return total
}

func main() {
	order := order{
		userID: 1,
		itens: []item{
			item{1, 2, 12.10},
			item{2, 1, 23.49},
			item{11, 100, 3.13},
		},
	}

	fmt.Printf("Valor total do pedido é R$%.2f", order.totalValue())
}
