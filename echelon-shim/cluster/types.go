package cluster

import (
	bs "github.com/SimonRichardson/echelon/internal/selectors"
	c "github.com/SimonRichardson/echelon/cluster"
)

// Incrementer defines a way to increment a value with in the store
type Incrementer interface {
	Increment(bs.Key) <-chan c.Element
}
