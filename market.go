package gomarket 

import (
	"github.com/garyburd/redigo/redis"
)

type market struct {
	conn redis.Conn
	name string
	buyOrders map[string][]order
	sellOrders map[string][]order
	orderId int64
}

func NewSeededMarket() *market {
	m := NewMarket()
	return m
}

func NewMarket() *market{
	ds, err := redis.Dial("tcp", ":6379")

	if err != nil {
		panic(err)
	}

	return &market{
		orderId: 0, 
		buyOrders: make(map[string][]order),
		sellOrders: make(map[string][]order),
		conn: ds,
	}
}

func (m *market) IncrOrderId() {
	_, err := m.conn.Do("INCR", "gomarket:globalOrderId")

	if err != nil {
		panic(err)
	}
}

func (m *market) NextOrderId() int64 {
	m.IncrOrderId()

	id, err := redis.Int64(m.conn.Do("GET", "gomarket:globalOrderId"))

	if err != nil {
		panic(err)
	}

	return id
}

func (m *market) submitMarketBuyOrder(quantity int32, resourceName string) {
	
	m.buyOrders[resourceName] = append(m.buyOrders[resourceName], NewMarketSellOrder(quantity, m.NextOrderId()))

} 

func (m *market) submitMarketSellOrder(quantity int32, resourceName string) {

	m.sellOrders[resourceName] = append(m.sellOrders[resourceName], NewMarketSellOrder(quantity, m.NextOrderId()))

}


