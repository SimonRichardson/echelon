package score

import (
	"time"

	"github.com/SimonRichardson/echelon/internal/permitters"
	r "github.com/SimonRichardson/echelon/echelon-shim/cluster/score"
	"github.com/SimonRichardson/echelon/errors"
	"github.com/SimonRichardson/echelon/internal/typex"
)

func noopTactic([]r.Cluster, func(r.Cluster)) error {
	return nil
}

func nonBlocking(clusters []r.Cluster, fn func(r.Cluster)) error {
	for _, c := range clusters {
		go func(c r.Cluster) {
			fn(c)
		}(c)
	}
	return nil
}

func rateLimited(requestsPerDuration int,
	requestsDuration time.Duration,
) func([]r.Cluster, func(r.Cluster)) error {
	permits := permitters.New(int64(requestsPerDuration), requestsDuration)
	return func(clusters []r.Cluster, fn func(r.Cluster)) error {
		if n := len(clusters); !permits.Allowed(int64(n)) {
			return typex.Errorf(errors.Source, errors.RateLimited,
				"RateLimited: element rate exceeded; request discarded")
		}
		return nonBlocking(clusters, fn)
	}
}
