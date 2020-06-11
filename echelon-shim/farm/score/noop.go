package score

import (
	"time"

	bs "github.com/SimonRichardson/echelon/internal/selectors"
	t "github.com/SimonRichardson/echelon/echelon-shim/selectors"
)

// NoopIncrementer defines a incrementer that performs no operations, but attempts to
// provide "canned" data that will allow the application to still execute.
func NoopIncrementer(f *Farm, t Tactic) t.Incrementer { return noop{f} }

type noop struct {
	*Farm
}

func (n noop) Increment(bs.Key, time.Time) (int, error) {
	return -1, nil
}
