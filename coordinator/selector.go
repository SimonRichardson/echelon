package coordinator

import (
	bs "github.com/SimonRichardson/echelon/internal/selectors"
	"github.com/SimonRichardson/echelon/farm/store"
	s "github.com/SimonRichardson/echelon/selectors"
)

type selector struct {
	s.LifeCycleManager

	co    *Coordinator
	store *store.Farm
}

func newSelector(co *Coordinator, store *store.Farm) *selector {
	return &selector{
		LifeCycleManager: newLifeCycleService(),

		co:    co,
		store: store,
	}
}

func (s *selector) Select(key, field bs.Key) (s.KeyFieldScoreTxnValue, error) {
	return s.store.Select(key, field)
}

func (s *selector) SelectRange(key bs.Key, limit int, maxSize s.KeySizeExpiry) ([]s.KeyFieldScoreTxnValue, error) {
	return s.store.SelectRange(key, limit, maxSize)
}
