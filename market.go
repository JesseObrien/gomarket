package gomarket

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

func Init() {

}

func NewRedisConnection() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")

	if err != nil {
		log.Fatal(err)
	}

	return c
}

type market struct {
	redis      redis.Conn
	name       string
	buyOrders  map[string][]BuyOrder
	sellOrders map[string][]SellOrder
	orderId    int64
}

func NewMarket() *market {
	return &market{
		orderId:    0,
		buyOrders:  make(map[string][]BuyOrder),
		sellOrders: make(map[string][]SellOrder),
		redis:      NewRedisConnection(),
	}
}

func (m *market) ListSymbol(s Symbol) {

	// Add the symbol to the symbols set
	m.redis.Send("SADD", redisKey("symbols"), s.Name)

	// Add the symbol object to a hash
	m.redis.Send("HMSET", redis.Args{}.Add(redisKey("symbols:"+s.Name)).AddFlat(&s))

	m.redis.Flush()

}

func (m *market) DelistSymbol(s Symbol) {

}

func (m *market) NextTransactionId() int64 {

	id, err := redis.Int64(m.redis.Do("INCR", redisKey("globalTransactionId")))

	if err != nil {
		log.Fatal(err)
	}

	return id
}

func (m *market) submitMarketBuyOrder(quantity int64, symbol string) {

	bo := NewBuyOrder(symbol, quantity)

	m.buyOrders[symbol] = append(m.buyOrders[bo.GetSymbol()], bo)
}

func (m *market) submitSellOrder(quantity int64, symbol string, price int64) {

	// Build a new sell order and then append it to the sell orders list
	so := NewSellOrder(symbol, quantity, price)

	m.sellOrders[so.GetSymbol()] = append(m.sellOrders[so.GetSymbol()], so)

	// Record sell order in redis hash

	if err := so.Record(); err != nil {
		log.Fatal(err)
	}
}
