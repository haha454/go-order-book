package main

import "container/heap"

type PriceLevelHeap struct {
	priceLevels                  []*PriceLevel
	priceLevelPerPrice           map[int]*PriceLevel
	priceLevelComparator         priceLevelComparator
	orderPriceCrossingComparator orderPriceCrossingComparator
}

type orderPriceCrossingComparator func(order *Order, priceLevel *PriceLevel) bool

func AskPriceLevelCrossed(order *Order, priceLevel *PriceLevel) bool {
	return order.price >= priceLevel.price
}

func BidPriceLevelCrossed(order *Order, priceLevel *PriceLevel) bool {
	return order.price <= priceLevel.price
}

type priceLevelComparator func(a, b *PriceLevel) bool

func AskPriceLevelHeapOrdered(a, b *PriceLevel) bool {
	return a.price < b.price
}

func BidPriceLevelHeapOrdered(a, b *PriceLevel) bool {
	return a.price > b.price
}

func (plh *PriceLevelHeap) IsPriceCrossed(order *Order) bool {
	if plh.GetTopPriceLevel() == nil {
		return false
	}
	return plh.orderPriceCrossingComparator(order, plh.GetTopPriceLevel())
}

func (plh *PriceLevelHeap) AddOrder(order *Order) {
	if _, exist := plh.priceLevelPerPrice[order.price]; !exist {
		plh.priceLevelPerPrice[order.price] = NewPriceLevel(order.price)
	}

	priceLevel := plh.priceLevelPerPrice[order.price]
	priceLevel.Add(order)
	plh.addPriceLevel(priceLevel)
}

func (plh *PriceLevelHeap) CancelOrder(order *Order) {
	if _, exist := plh.priceLevelPerPrice[order.price]; !exist {
		panic("non-existent order")
	}

	plh.priceLevelPerPrice[order.price].Cancel(order.id)
}

func (plh *PriceLevelHeap) addPriceLevel(priceLevel *PriceLevel) {
	heap.Push(plh, priceLevel)
}

// nil if the heap is empty
// ALways return a nil or a price level with at least one order
func (plh *PriceLevelHeap) GetTopPriceLevel() *PriceLevel {
	for len(plh.priceLevels) > 0 && plh.priceLevels[0].GetOldestOrder() == nil {
		plh.popPriceLevel()
	}

	if len(plh.priceLevels) == 0 {
		return nil
	}

	return plh.priceLevels[0]
}

// nil if the heap is empty
func (plh *PriceLevelHeap) popPriceLevel() *PriceLevel {
	if len(plh.priceLevels) == 0 {
		return nil
	}
	priceLevel := heap.Pop(plh).(*PriceLevel)
	delete(plh.priceLevelPerPrice, priceLevel.price)
	return priceLevel
}

func NewAskPriceLevelHeap() *PriceLevelHeap {
	return &PriceLevelHeap{
		priceLevels:                  make([]*PriceLevel, 0),
		priceLevelPerPrice:           make(map[int]*PriceLevel),
		priceLevelComparator:         AskPriceLevelHeapOrdered,
		orderPriceCrossingComparator: AskPriceLevelCrossed,
	}
}

func NewBidPriceLevelHeap() *PriceLevelHeap {
	return &PriceLevelHeap{
		priceLevels:                  make([]*PriceLevel, 0),
		priceLevelPerPrice:           make(map[int]*PriceLevel),
		priceLevelComparator:         BidPriceLevelHeapOrdered,
		orderPriceCrossingComparator: BidPriceLevelCrossed,
	}
}

func (plh *PriceLevelHeap) Len() int {
	return len(plh.priceLevels)
}

func (plh *PriceLevelHeap) Less(i, j int) bool {
	return plh.priceLevelComparator(plh.priceLevels[i], plh.priceLevels[j])
}

func (plh *PriceLevelHeap) Swap(i, j int) {
	plh.priceLevels[i], plh.priceLevels[j] = plh.priceLevels[j], plh.priceLevels[i]
}

func (plh *PriceLevelHeap) Push(x any) {
	plh.priceLevels = append(plh.priceLevels, x.(*PriceLevel))
}

func (plh *PriceLevelHeap) Pop() any {
	length := len(plh.priceLevels)
	result := plh.priceLevels[length-1]
	plh.priceLevels = plh.priceLevels[:length-1]
	return result
}
