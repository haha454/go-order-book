package main

import "container/list"

// Order cancellation removes the order at its price level
type PriceLevel struct {
	price             int
	orders            *list.List
	elementPerOrderId map[string]*list.Element
}

func NewPriceLevel(price int) *PriceLevel {
	return &PriceLevel{price: price, orders: list.New(), elementPerOrderId: make(map[string]*list.Element)}
}

func (priceLevel *PriceLevel) Add(order *Order) {
	if priceLevel.price != order.price {
		panic("shouldn't reach here")
	}

	if _, exist := priceLevel.elementPerOrderId[order.id]; exist {
		panic("duplicate order ID")
	}

	priceLevel.elementPerOrderId[order.id] = priceLevel.orders.PushBack(order)
}

func (priceLevel *PriceLevel) Cancel(id string) {
	if _, exist := priceLevel.elementPerOrderId[id]; !exist {
		panic("non-existent order ID")
	}

	priceLevel.orders.Remove(priceLevel.elementPerOrderId[id])
	delete(priceLevel.elementPerOrderId, id)
}

// nil if no orders at this price level
// Might return an order with zero quantity
// Client needs to call PopOldestOrder until it sees a non-zero quantity Order or nil, which means the price level is empty
func (priceLevel *PriceLevel) GetOldestOrder() *Order {
	if priceLevel.orders.Front() == nil {
		return nil
	}
	return priceLevel.orders.Front().Value.(*Order)
}

// nil if no orders at this price level
func (priceLevel *PriceLevel) PopOldestOrder() *Order {
	if priceLevel.orders.Front() == nil {
		return nil
	}
	order := priceLevel.orders.Remove(priceLevel.orders.Front()).(*Order)
	delete(priceLevel.elementPerOrderId, order.id)
	return order
}
