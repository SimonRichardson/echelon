package request

import (
	"time"

	"github.com/SimonRichardson/echelon/internal/errors"
	"github.com/SimonRichardson/echelon/internal/permitters"
	"github.com/SimonRichardson/echelon/internal/typex"
)

func noopTactic([]Cluster, func(int, Cluster)) error {
	return nil
}

func nonBlocking(clusters []Cluster, fn func(int, Cluster)) error {
	for k, c := range clusters {
		go func(k int, c Cluster) {
			fn(k, c)
		}(k, c)
	}
	return nil
}

func rateLimited(requestPerDuration int,
	requestsDuration time.Duration,
) func([]Cluster, func(int, Cluster)) error {
	permits := permitters.New(int64(requestPerDuration), requestsDuration)
	return func(clusters []Cluster, fn func(int, Cluster)) error {
		if n := len(clusters); !permits.Allowed(int64(n)) {
			return typex.Errorf(errors.Source, errors.RateLimited,
				"RateLimited: element rate exceeded; request discarded")
		}
		return nonBlocking(clusters, fn)
	}
}
