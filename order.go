package main

type Order struct {
	side     Side
	id       string
	quantity int
	price    int
}

func NewOrder(side Side, id string, quantity int, price int) *Order {
	return &Order{side: side, id: id, quantity: quantity, price: price}
}

func NewBidOrder(id string, quantity int, price int) *Order {
	return NewOrder(Bid, id, quantity, price)
}

func NewAskOrder(id string, quantity int, price int) *Order {
	return NewOrder(Ask, id, quantity, price)
}
