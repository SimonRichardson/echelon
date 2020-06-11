package coordinator

import (
	"github.com/SimonRichardson/echelon/coordinator/strategies"
	"github.com/SimonRichardson/echelon/farm/store"
	s "github.com/SimonRichardson/echelon/selectors"
)

type repairer struct {
	s.LifeCycleManager

	co       *Coordinator
	store    *store.Farm
	strategy strategies.RepairStrategy
}

func newRepairer(co *Coordinator, store *store.Farm, strategy strategies.RepairStrategy) *repairer {
	return &repairer{
		LifeCycleManager: newLifeCycleService(),

		co:       co,
		store:    store,
		strategy: strategy,
	}
}

func (s *repairer) Repair(members []s.KeyFieldTxnValue, maxSize s.KeySizeExpiry) error {
	return s.strategy(s.store, members, maxSize)
}
