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
	"github.com/garyburd/redigo/redis"
	"strconv"
	"strings"
	"time"
)

type order interface {
	execute() bool
	cancel() bool
	record() bool
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

func NewSellOrder(s string, q int64) SellOrder {
	t := time.Now()
	oid := NextOrderId(s)
	return SellOrder{Order{Symbol: s, Quantity: q, OrderId: oid, Time: t.UTC()}}
}

func NewBuyOrder(s string, q int64) BuyOrder {
	t := time.Now()
	oid := NextOrderId(s)
	return BuyOrder{Order{Symbol: s, Quantity: q, OrderId: oid, Time: t.UTC()}}
}

func NextOrderId(symbol string) int64 {

	r := NewRedisConnection()

	r.Send("INCR", redisKey(symbol+":uOrderId"))
	r.Send("GET", redisKey(symbol+":uOrderId"))
	r.Flush()

	id, err := redis.Int64(r.Receive())

	if err != nil {
		panic(err)
	}

	return id
}

func (o *Order) GetSymbol() string {
	// s := strings.ToUpper(symbol)
	return strings.ToUpper(o.Symbol)
}

func (s *SellOrder) cancel() bool {
	return true
}

func (s *SellOrder) execute() bool {
	return true
}

func (s *SellOrder) Record() bool {
	r := NewRedisConnection()

	orderId := NextOrderId(s.GetSymbol())

	orderIdStr := strconv.FormatInt(orderId, 10)

	r.Send("LPUSH", redisKey("sellorders:"+s.GetSymbol()), orderIdStr)
	r.Send("HMSET", redis.Args{}.Add(redisKey("sellorders:"+s.Symbol+":"+orderIdStr)).AddFlat(s))

	return true
}
