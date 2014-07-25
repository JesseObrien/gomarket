package gomarket

import (
	"fmt"
	"testing"
)

func newMarketSeederFactory() *marketseeder {
	m := NewMarketSeeder()
	m.seedSymbols()
	m.seedSellOrders(2, "jobr")
	m.seedBuyOrders(2, "jobr")
	return m
}

func TestAddBuyOrder(t *testing.T) {
	m := newMarketSeederFactory()

	l := len(m.market.sellOrders["JOBR"])

	if l != 2 {
		t.Errorf("Number of sell orders: %v, want %v", l, 1)
	}

	o := m.market.sellOrders["JOBR"][0]

	fmt.Println(o.Price)
	if o.Quantity < 0 {
		t.Errorf("Quantity: $v, want %v", o.Quantity, "non-zero")
	}

}
