package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPriceLevel_GetOldestOrder(t *testing.T) {
	t.Run("Add two orders -> GetOldestOrder should return the first one", func(t *testing.T) {
		priceLevel := NewPriceLevel(0xA)
		firstOrder := NewBidOrder("1", 0xB, 0xA)
		priceLevel.Add(firstOrder)
		priceLevel.Add(NewBidOrder("2", 0xB, 0xA))
		assert.Equal(t, firstOrder, priceLevel.GetOldestOrder())
	})

	t.Run("GetOldestOrder should return nil if no orders are at the level", func(t *testing.T) {
		priceLevel := NewPriceLevel(0xA)
		assert.Nil(t, priceLevel.GetOldestOrder())
	})

	t.Run("Add two orders and Cancel the first -> GetOldestOrder should return the second", func(t *testing.T) {
		priceLevel := NewPriceLevel(0xA)
		firstOrder := NewBidOrder("1", 0xB, 0xA)
		priceLevel.Add(firstOrder)
		secondOrder := NewBidOrder("2", 0xB, 0xA)
		priceLevel.Add(secondOrder)
		priceLevel.Cancel("1")
		assert.Equal(t, secondOrder, priceLevel.GetOldestOrder())
	})

	t.Run("Add two orders and PopOldestOrder -> GetOldestOrder should return the second", func(t *testing.T) {
		priceLevel := NewPriceLevel(0xA)
		firstOrder := NewBidOrder("1", 0xB, 0xA)
		priceLevel.Add(firstOrder)
		secondOrder := NewBidOrder("2", 0xB, 0xA)
		priceLevel.Add(secondOrder)
		_ = priceLevel.PopOldestOrder()
		assert.Equal(t, secondOrder, priceLevel.GetOldestOrder())
	})
}

func TestPriceLevel_PopOldestOrder(t *testing.T) {
	t.Run("Add two orders -> PopOldestOrder twice should return the first and the second", func(t *testing.T) {
		priceLevel := NewPriceLevel(0xA)
		firstOrder := NewBidOrder("1", 0xB, 0xA)
		priceLevel.Add(firstOrder)
		secondOrder := NewBidOrder("2", 0xB, 0xA)
		priceLevel.Add(secondOrder)
		assert.Equal(t, firstOrder, priceLevel.PopOldestOrder())
		assert.Equal(t, secondOrder, priceLevel.PopOldestOrder())
	})

	t.Run("PopOldestOrder should return nil if no orders are at the level", func(t *testing.T) {
		priceLevel := NewPriceLevel(0xA)
		assert.Nil(t, priceLevel.PopOldestOrder())
	})

	t.Run("Add two orders and Cancel the first -> PopOldestOrder should return the second", func(t *testing.T) {
		priceLevel := NewPriceLevel(0xA)
		firstOrder := NewBidOrder("1", 0xB, 0xA)
		priceLevel.Add(firstOrder)
		secondOrder := NewBidOrder("2", 0xB, 0xA)
		priceLevel.Add(secondOrder)
		priceLevel.Cancel("1")
		assert.Equal(t, secondOrder, priceLevel.PopOldestOrder())
	})
}
