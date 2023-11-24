package main

type OrderBook struct {
	ask *OneSideOrderBook
	bid *OneSideOrderBook
	// stores all orders even if the order has been cancelled or fully filled
	orderPerOrderId map[string]*Order
}

func (ob *OrderBook) Add(order *Order) (trades []Trade, closedOrderIds []string) {
	if _, exist := ob.orderPerOrderId[order.id]; exist {
		panic("duplicate order ID")
	}

	switch order.side {
	case Bid:
		trades, closedOrderIds = ob.ask.Match(order)
		if order.quantity > 0 {
			ob.bid.Add(order)
		}
	case Ask:
		trades, closedOrderIds = ob.bid.Match(order)
		if order.quantity > 0 {
			ob.ask.Add(order)
		}
	}

	ob.orderPerOrderId[order.id] = order
	if order.quantity == 0 {
		closedOrderIds = append(closedOrderIds, order.id)
	}
	return trades, closedOrderIds
}

func (ob *OrderBook) Cancel(id string) {
	if _, exist := ob.orderPerOrderId[id]; !exist {
		panic("non-existent order ID")
	}

	switch order := ob.orderPerOrderId[id]; order.side {
	case Bid:
		ob.bid.Cancel(order)
	case Ask:
		ob.ask.Cancel(order)
	}
}

func NewOrderBook(orders ...*Order) *OrderBook {
	orderBook := &OrderBook{ask: NewAskSideOrderBook(), bid: NewBidSideOrderBook(), orderPerOrderId: make(map[string]*Order)}
	for _, order := range orders {
		orderBook.Add(order)
	}

	return orderBook
}
