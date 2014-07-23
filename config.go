package gomarket

func redisKey(s string) string {
	return "gomarket:" + s
}
