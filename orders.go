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

type order interface {
	execute() bool
	cancel() bool
}

type Order struct {
	Quantity int64
	OrderId  int64
	Price    float64
}

type SellOrder struct {
	Order
}

type BuyOrder struct {
	Order
}

func NewSellOrder(q int64, o int64) SellOrder {
	return SellOrder{Order{Quantity: q, OrderId: o}}
}

func NewBuyOrder(q int64, o int64) BuyOrder {
	return BuyOrder{Order{Quantity: q, OrderId: o}}
}

func (m SellOrder) cancel() bool {
	return true
}

func (m SellOrder) execute() bool {
	return true
}
