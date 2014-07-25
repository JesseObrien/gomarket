package gomarket

import (
	"github.com/garyburd/redigo/redis"
	"github.com/manveru/faker"
	"math/rand"
	"strings"
	"time"
)

type marketseeder struct {
	conn   redis.Conn
	faker  *faker.Faker
	market *market
}

func NewMarketSeeder() *marketseeder {

	f, _ := faker.New("en")
	m := NewMarket()

	return &marketseeder{
		conn:   NewRedisConnection(),
		faker:  f,
		market: m,
	}
}

func (ms *marketseeder) seedSymbols() {

	maxPrice := 400000 // $4,000.00
	minPrice := 100    // $1.00

	symbol := ms.makeSymbol()
	rand.Seed(time.Now().Unix())
	price := rand.Intn(maxPrice-minPrice) + minPrice

	if _, err := ms.conn.Do("SET", "gomarket:"+symbol+":price", price); err != nil {
		panic(err)
	}

}

func (ms *marketseeder) seedBuyOrders(num int64, resource string) {

}

func (ms *marketseeder) seedSellOrders(num int64, resource string) {
	for n := int64(0); n < num; n++ {
		ms.market.submitSellOrder(int64(rand.Int63n(100)), resource, rand.Int63()*10.0)
	}
}

func (ms *marketseeder) makeSymbol() string {

	// @TODO Make this return only alphabet characters
	s := ms.faker.Characters(4)

	return strings.ToUpper(s)
}
