package gomarket

import (
	"fmt"
	"testing"
)

func newMarketFactory() *market {
	m := NewMarket()
	m.seedSellOrders(2, "JOBR")
	m.seedBuyOrders(2, "JOBR")
	return m
}

func TestAddBuyOrder(t *testing.T) {
	m := newMarketFactory()

	l := len(m.sellOrders["JOBR"])

	if l != 2 {
		t.Errorf("Number of sell orders: %v, want %v", l, 1)
	}

	o := m.sellOrders["JOBR"][0]

	fmt.Println(o.Price)
	if o.Quantity < 0 {
		t.Errorf("Quantity: $v, want %v", o.Quantity, "non-zero")
	}

}
