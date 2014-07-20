package gomarket

import (
	"github.com/garyburd/redigo/redis"
	"math/rand"
	"strconv"
	"strings"
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

func (m *market) IncrSymbolOrderId(symbol string) int64 {
	m.conn.Send("INCR", "gomarket:"+symbol+":uOrderId")
	m.conn.Send("GET", "gomarket:"+symbol+":uOrderId")

	m.conn.Flush()
	id, err := redis.Int64(m.conn.Receive())

	if err != nil {
		panic(err)
	}

	return id
}

func (m *market) submitMarketBuyOrder(quantity int64, symbol string) {

	s := strings.ToUpper(symbol)
	orderid := m.IncrSymbolOrderId(s)

	m.buyOrders[symbol] = append(m.buyOrders[symbol], NewBuyOrder(s, quantity, orderid))
}

func (m *market) submitSellOrder(quantity int64, symbol string, price float64) {

	// Straighte up the symbol
	s := strings.ToUpper(symbol)
	// Get a new order id
	orderid := m.IncrSymbolOrderId(s)

	// Build a new sell order and then append it to the sell orders list
	so := NewSellOrder(s, quantity, orderid)
	m.sellOrders[symbol] = append(m.sellOrders[s], so)

	orderIdStr := strconv.FormatInt(orderid, 10)

	// Record sell order in redis hash
	m.conn.Send("LPUSH", "gomarket:sellorders:"+s, orderIdStr)
	m.conn.Send("HMSET", redis.Args{}.Add("gomarket:sellorders:"+s+":"+orderIdStr).AddFlat(&so)...)

	m.conn.Flush()
	_, err := m.conn.Receive()

	if err != nil {
		panic(err)
	}

}
