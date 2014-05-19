package gomarket 

import (
	"testing"
)

func TestAddBuyOrder(t *testing.T) {
	m := NewMarket()
	m.submitMarketSellOrder(20, "coal")

	l := len(m.sellOrders["coal"])

	if l != 1 {
		t.Errorf("Number of coal orders: %v, want %v", l, 1)
	}

	o := m.sellOrders["coal"][0]

	if o.getQuantity() != 20 {
		t.Errorf("Quantity: $v, want %v", o.getQuantity(), 20)
	}

}
