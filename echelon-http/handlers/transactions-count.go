package handlers

import (
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	bs "github.com/SimonRichardson/echelon/internal/selectors"
	"github.com/SimonRichardson/echelon/echelon-http/responses"
	"github.com/SimonRichardson/echelon/coordinator"
	"github.com/SimonRichardson/echelon/errors"
	"github.com/SimonRichardson/echelon/internal/typex"
)

// TransactionsCount returns the specific size of a collection with in the store.
func TransactionsCount(co *coordinator.Coordinator) http.HandlerFunc {
	return accepts(func(w http.ResponseWriter, r *http.Request) {
		began := time.Now()

		key := r.URL.Query().Get(":key")
		if !bson.IsObjectIdHex(key) {
			responses.BadRequest(w, r, typex.Errorf(errors.Source, errors.InvalidArgument,
				"Invalid Key: %s", key))
			return
		}

		counts, err := co.Size(bs.Key(key))
		if err != nil {
			responses.InternalServerError(w, r, err)
			return
		}

		responses.OKInt(w, counts, time.Since(began))
		return
	})
}
