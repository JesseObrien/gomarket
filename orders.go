package gomarket

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"strconv"
	"strings"
	"time"
)

type order interface {
	execute() bool
	cancel() bool
	record() bool
	GetSymbol() string
	GetPrice() float64
}

type Order struct {
	Symbol   string
	Quantity int64
	OrderId  int64
	Price    int64
	Time     time.Time
}

type SellOrder struct {
	Order
}

type BuyOrder struct {
	Order
}

func NewSellOrder(s string, q int64, price int64) SellOrder {
	t := time.Now()
	oid := NextOrderId(s)

	return SellOrder{Order{Symbol: s, Quantity: q, OrderId: oid, Time: t.UTC(), Price: price}}
}

func NewBuyOrder(s string, q int64) BuyOrder {
	t := time.Now()
	oid := NextOrderId(s)
	return BuyOrder{Order{Symbol: s, Quantity: q, OrderId: oid, Time: t.UTC()}}
}

func NextOrderId(symbol string) int64 {

	r := NewRedisConnection()

	id, err := redis.Int64(r.Do("INCR", redisKey(symbol+":uOrderId")))

	if err != nil {
		log.Fatal(err)
	}

	return id
}

func (o *Order) GetSymbol() string {
	return strings.ToUpper(o.Symbol)
}

func (s *SellOrder) cancel() bool {
	return true
}

func (s *SellOrder) execute() bool {
	return true
}

func (s *SellOrder) Record() error {
	r := NewRedisConnection()

	orderIdStr := strconv.FormatInt(s.OrderId, 10)

	// Add the order id to the sell orders for this symbol
	r.Send("LPUSH", redisKey(s.GetSymbol()+":orders:sell"), orderIdStr)

	// Push the order details into a hash
	r.Send("HMSET", redis.Args{}.Add(redisKey(s.GetSymbol()+":orders:sell:"+orderIdStr)).AddFlat(s))

	r.Flush()

	if _, err := r.Receive(); err != nil {
		return err
	}

	return nil
}
