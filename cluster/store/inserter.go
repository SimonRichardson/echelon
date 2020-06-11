package store

import (
	"time"

	t "github.com/SimonRichardson/echelon/cluster"
	s "github.com/SimonRichardson/echelon/selectors"
	"github.com/garyburd/redigo/redis"
)

const (
	defaultFieldExists    = 0
	defaultFieldInsertion = 1
)

func insertion(conn redis.Conn, members []s.KeyFieldScoreTxnValue, sizeExpiry s.SizeExpiry) ([]s.KeyCount, error) {
	var (
		now    = time.Now()
		expiry = now.Add(sizeExpiry.Expiry).UnixNano()
	)

	for _, member := range members {
		if err := sendInsertScript(conn,
			member.Key,
			member.Field,
			member.Score,
			expiry,
			member.Txn,
			member.Value,
		); err != nil {
			return generateResult(members, 0), err
		}
	}

	if err := conn.Flush(); err != nil {
		return generateResult(members, 0), err
	}

	if !defaultVerifyResults {
		return generateResult(members, 1), nil
	}

	result := make([]s.KeyCount, 0, len(members))

	for _, m := range members {
		res, err := redis.Int(conn.Receive())
		if err != nil {
			return result, err
		}

		if res == defaultFieldExists || res == defaultFieldInsertion {
			result = append(result, s.KeyCount{Key: m.Key, Count: 1})
		}
	}

	if len(result) < len(members) {
		return result, t.ErrPartialInsertions
	}

	return result, nil
}
