package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderBook_PartiallyFilled(t *testing.T) {
	orderBook := NewOrderBook(
		NewOrder(
			Ask,
			"A",
			2,
			15,
		),
		NewOrder(
			Ask,
			"B",
			2,
			10,
		),
		NewOrder(
			Ask,
			"C",
			3,
			10,
		),
		NewOrder(
			Ask,
			"D",
			5,
			20,
		),
	)

	trades, closedOrderIds := orderBook.Add(NewOrder(
		Bid,
		"Z",
		8,
		18))
	assert.Equal(t, []Trade{
		{
			"B",
			"Z",
			2,
			10,
		},
		{
			"C",
			"Z",
			3,
			10,
		},
		{
			"A",
			"Z",
			2,
			15,
		},
	}, trades)

	assert.Equal(t, []string{"B", "C", "A"}, closedOrderIds)
}

func TestOrderBook_FullyFilled(t *testing.T) {
	orderBook := NewOrderBook(
		NewOrder(
			Ask,
			"A",
			2,
			15,
		),
		NewOrder(
			Ask,
			"B",
			2,
			10,
		),
		NewOrder(
			Ask,
			"C",
			3,
			10,
		),
		NewOrder(
			Ask,
			"D",
			5,
			20,
		),
	)

	trades, closedOrderIds := orderBook.Add(NewOrder(
		Bid,
		"Z",
		6,
		18))
	assert.Equal(t, []Trade{
		{
			"B",
			"Z",
			2,
			10,
		},
		{
			"C",
			"Z",
			3,
			10,
		},
		{
			"A",
			"Z",
			1,
			15,
		},
	}, trades)

	assert.Equal(t, []string{"B", "C", "Z"}, closedOrderIds)
}
