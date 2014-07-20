package gomarket

/**
*
* gomarket:orders:marketBuy
* gomarket:orders:marketSell
* gomarket:orders:limitBuy
* gomarket:orders:limitSell
*
*
*
*
*
*
 */
import (
	"time"
)

type order interface {
	execute() bool
	cancel() bool
}

type Order struct {
	Symbol   string
	Quantity int64
	OrderId  int64
	Price    float64
	Time     time.Time
}

type SellOrder struct {
	Order
}

type BuyOrder struct {
	Order
}

func NewSellOrder(s string, q int64, oid int64) SellOrder {
	t := time.Now()
	return SellOrder{Order{Symbol: s, Quantity: q, OrderId: oid, Time: t.UTC()}}
}

func NewBuyOrder(s string, q int64, oid int64) BuyOrder {
	t := time.Now()
	return BuyOrder{Order{Symbol: s, Quantity: q, OrderId: oid, Time: t.UTC()}}
}

func (m SellOrder) cancel() bool {
	return true
}

func (m SellOrder) execute() bool {
	return true
}
