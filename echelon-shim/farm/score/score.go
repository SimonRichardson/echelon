package score

import (
	"time"

	bs "github.com/SimonRichardson/echelon/internal/selectors"
	st "github.com/SimonRichardson/echelon/echelon-shim/selectors"
)

// IncrementFromTime defines a strategy to write to all the cluster and then
// wait for all the cluster items to respond before continuing onwards.
func IncrementFromTime(f *Farm, t Tactic) st.Incrementer {
	return incrementFromTime{f, t}
}

type incrementFromTime struct {
	*Farm
	tactic Tactic
}

func (w incrementFromTime) Increment(key bs.Key, t time.Time) (int, error) {
	now := time.Now()
	return scoredValue(now) - scoredValue(t), nil
}

func scoredValue(t time.Time) int {
	return int(t.Unix()*1e7 + int64(t.Nanosecond())/100)
}
