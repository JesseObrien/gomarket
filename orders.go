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
	getOrderId() int64
	getQuantity() int32
}

type MarketSellOrder struct {
	quantity int32
	orderId int64
}

func NewMarketSellOrder(q int32, o int64) order {
	return &MarketSellOrder{quantity: q, orderId: o}
}

func (m MarketSellOrder) cancel() bool {
	return true
}

func (m MarketSellOrder) execute() bool {

	return true
}

func (m MarketSellOrder) getOrderId() int64 {
	return m.orderId
}

func (m MarketSellOrder) getQuantity() int32 {
	return m.quantity
}

