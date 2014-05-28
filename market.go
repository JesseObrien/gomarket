package gomarket

import (
	"github.com/garyburd/redigo/redis"
	"math/rand"
)

func Init() {

}

type market struct {
	conn       redis.Conn
	name       string
	buyOrders  map[string][]BuyOrder
	sellOrders map[string][]SellOrder
	orderId    int64
}

func NewMarket() *market {
	ds, err := redis.Dial("tcp", ":6379")

	if err != nil {
		panic(err)
	}

	return &market{
		orderId:    0,
		buyOrders:  make(map[string][]BuyOrder),
		sellOrders: make(map[string][]SellOrder),
		conn:       ds,
	}
}

func (m *market) seedSellOrders(num int64, resource string) {
	for n := int64(0); n < num; n++ {
		m.submitSellOrder(int64(rand.Int63n(100)), resource, rand.Float64()*10.0)
	}
}

func (m *market) seedBuyOrders(num int64, resource string) {

}

func (m *market) IncrOrderId() {
	_, err := m.conn.Do("INCR", "gomarket:globalOrderId")

	if err != nil {
		panic(err)
	}
}

func (m *market) nextOrderId() int64 {
	m.IncrOrderId()

	id, err := redis.Int64(m.conn.Do("GET", "gomarket:globalOrderId"))

	if err != nil {
		panic(err)
	}

	return int64(id)
}

func (m *market) submitMarketBuyOrder(quantity int64, resourceName string) {
	m.buyOrders[resourceName] = append(m.buyOrders[resourceName], NewBuyOrder(quantity, m.nextOrderId()))
}

func (m *market) submitSellOrder(quantity int64, resourceName string, price float64) {

	m.sellOrders[resourceName] = append(m.sellOrders[resourceName], NewSellOrder(quantity, m.nextOrderId()))

}
