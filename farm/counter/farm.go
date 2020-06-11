package counter

import (
	bs "github.com/SimonRichardson/echelon/internal/selectors"
	c "github.com/SimonRichardson/echelon/cluster/counter"
	"github.com/SimonRichardson/echelon/farm"
	"github.com/SimonRichardson/echelon/instrumentation"
	s "github.com/SimonRichardson/echelon/selectors"
)

// Tactic defines an alias for the structure of a tactic. A tactic in this
// regard consumes a slice of clusters and runs the function depending on if
// the tactic allows it.
type Tactic func([]c.Cluster, func(int, c.Cluster)) error

// insertStrategy, deleteStrategy and scanStrategy defines a
// way to create the selectors from a farm and a tactic. They have a close
// analogy to building (creating) something from a some of parts.
type (
	selectStrategy func(*Farm, Tactic) s.Selector
	insertStrategy func(*Farm, Tactic) s.Inserter
	modifyStrategy func(*Farm, Tactic) s.Modifier
	deleteStrategy func(*Farm, Tactic) s.Deleter
	scanStrategy   func(*Farm, Tactic) s.Scanner
	repairStrategy func(*Farm, Tactic) s.Repairer
)

// InsertCreator, DeleteCreator and ScanCreator defines a series
// of types that allow the building (creating) of a selector using only the
// farm as an argument. It's expected that each is partially built before hand
// and the farm is the last argument for the partial application.
type (
	InsertCreator interface {
		Apply(*Farm) s.Inserter
	}
	DeleteCreator interface {
		Apply(*Farm) s.Deleter
	}
	ScanCreator interface {
		Apply(*Farm) s.Scanner
	}
	RepairCreator interface {
		Apply(*Farm) s.Repairer
	}
)

// Farm defines a container for all the selectors to be able to query.
type Farm struct {
	clusters        []c.Cluster
	inserter        s.Inserter
	deleter         s.Deleter
	scanner         s.Scanner
	repairer        s.Repairer
	instrumentation instrumentation.Instrumentation
}

// New defines a function for the creation of a farm.
func New(clusters []c.Cluster,
	ins InsertCreator,
	del DeleteCreator,
	sca ScanCreator,
	rep RepairCreator,
	instr instrumentation.Instrumentation,
) *Farm {
	farm := &Farm{
		clusters:        clusters,
		instrumentation: instr,
	}
	farm.inserter = ins.Apply(farm)
	farm.deleter = del.Apply(farm)
	farm.scanner = sca.Apply(farm)
	farm.repairer = rep.Apply(farm)
	return farm
}

// Insert defines a way to insert some members into the store that's associated
// with the key
func (f *Farm) Insert(members []s.KeyFieldScoreTxnValue, maxSize s.KeySizeExpiry) (int, error) {
	// TODO work out when to change strategies
	res, err := f.inserter.Insert(members, maxSize)
	return res, farm.PartialRepairError(err, func() {
		f.Repair(s.KeyFieldScoreTxnValues(members).KeyFieldTxnValues(), maxSize)
	})
}

// Delete removes a set of members associated with a key with in the store
func (f *Farm) Delete(members []s.KeyFieldScoreTxnValue, maxSize s.KeySizeExpiry) (int, error) {
	// TODO work out when to change strategies
	res, err := f.deleter.Delete(members, maxSize)
	return res, farm.PartialRepairError(err, func() {
		f.Repair(s.KeyFieldScoreTxnValues(members).KeyFieldTxnValues(), maxSize)
	})
}

// Keys returns all the keys with in the store
func (f *Farm) Keys() ([]bs.Key, error) {
	return f.scanner.Keys()
}

// Size defines a way to find the size associated with the key
func (f *Farm) Size(key bs.Key) (int, error) {
	return f.scanner.Size(key)
}

// Members defines a way to return all member keys associated with the key
func (f *Farm) Members(key bs.Key) ([]bs.Key, error) {
	return f.scanner.Members(key)
}

// Repair attempts to repair the store depending on the elements
func (f *Farm) Repair(elements []s.KeyFieldTxnValue, maxSize s.KeySizeExpiry) error {
	return f.repairer.Repair(elements, maxSize)
}

func (f *Farm) Topology(clusters []c.Cluster) error {
	for _, v := range f.clusters {
		if err := v.Close(); err != nil {
			return err
		}
	}

	f.clusters = clusters
	return nil
}
