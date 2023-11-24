package main

type OneSideOrderBook struct {
	priceLevelHeap *PriceLevelHeap
}

func (osob *OneSideOrderBook) Add(order *Order) {
	if order.quantity <= 0 {
		panic("order quantity must be positive")
	}

	osob.priceLevelHeap.AddOrder(order)
}

func (osob *OneSideOrderBook) Match(order *Order) ([]Trade, []string) {
	trades := make([]Trade, 0)
	closedOrderIds := make([]string, 0)
	for order.quantity > 0 && osob.priceLevelHeap.IsPriceCrossed(order) {
		makerPriceLevel := osob.priceLevelHeap.GetTopPriceLevel()
		makerOrder := makerPriceLevel.GetOldestOrder()
		tradeQuantity := min(order.quantity, makerOrder.quantity)
		trades = append(trades, Trade{
			makerOrderId: makerOrder.id,
			takerOrderId: order.id,
			quantity:     tradeQuantity,
			price:        makerOrder.price,
		})

		order.quantity -= tradeQuantity
		makerOrder.quantity -= tradeQuantity

		if makerOrder.quantity == 0 {
			closedOrderIds = append(closedOrderIds, makerOrder.id)
			makerPriceLevel.PopOldestOrder()
		}
	}

	return trades, closedOrderIds
}

func (osob *OneSideOrderBook) Cancel(order *Order) {
	osob.priceLevelHeap.CancelOrder(order)
}

func NewAskSideOrderBook() *OneSideOrderBook {
	return &OneSideOrderBook{
		priceLevelHeap: NewAskPriceLevelHeap(),
	}
}

func NewBidSideOrderBook() *OneSideOrderBook {
	return &OneSideOrderBook{
		priceLevelHeap: NewBidPriceLevelHeap(),
	}
}
