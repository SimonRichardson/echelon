package score

import (
	bs "github.com/SimonRichardson/echelon/internal/selectors"
	s "github.com/SimonRichardson/echelon/selectors"
	"github.com/garyburd/redigo/redis"
)

const (
	prefix    = "i:"
	prefixLen = len(prefix)
)

func increment(conn redis.Conn, key bs.Key) (s.KeyCount, error) {
	res, err := redis.Int(conn.Do("INCR", prefix+key.String()))
	return s.KeyCount{
		Key:   key,
		Count: res,
	}, err
}
