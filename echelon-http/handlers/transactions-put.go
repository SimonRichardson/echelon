package handlers

import (
	"bytes"
	"io"
	"net/http"
	"time"

	bs "github.com/SimonRichardson/echelon/internal/selectors"
	"github.com/SimonRichardson/echelon/echelon-http/responses"
	"github.com/SimonRichardson/echelon/coordinator"
	"github.com/SimonRichardson/echelon/errors"
	"github.com/SimonRichardson/echelon/schemas/pool"
	"github.com/SimonRichardson/echelon/schemas/records"
	"github.com/SimonRichardson/echelon/schemas/schema"
	"github.com/SimonRichardson/echelon/selectors"
	"github.com/SimonRichardson/echelon/internal/typex"
	"gopkg.in/mgo.v2/bson"
)

// TransactionsPut modifies items into the collection
func TransactionsPut(co *coordinator.Coordinator) http.HandlerFunc {
	return guard(func(w http.ResponseWriter, r *http.Request) {
		began := time.Now()

		queryKey := r.URL.Query().Get(":key")
		if !bson.IsObjectIdHex(queryKey) {
			responses.BadRequest(w, r, typex.Errorf(errors.Source, errors.InvalidArgument,
				"Invalid Key: %s", queryKey))
			return
		}

		fieldValues, score, maxSize, expiry, err := readPutRecords(r.Body)
		if err != nil {
			responses.BadRequest(w, r, err)
			return
		}

		var (
			key           = bs.Key(queryKey)
			maxSizeExpiry = selectors.MakeKeySizeSingleton(key, maxSize, expiry)

			elements           = fieldValues.KeyFieldScoreTxnValues(key, score)
			results, modifyErr = co.Modify(elements, maxSizeExpiry)
		)
		if modifyErr != nil {
			responses.Error(w, r, modifyErr)
			return
		}

		responses.OKInt(w, results, time.Since(began))
		return
	})
}

func readPutRecords(read io.ReadCloser) (selectors.FieldTxnValues, float64, int64, time.Duration, error) {
	var (
		buffer bytes.Buffer
		fail   = func(err error) (selectors.FieldTxnValues, float64, int64, time.Duration, error) {
			return nil, 0, 0, time.Duration(0), err
		}
	)
	if _, err := buffer.ReadFrom(read); err != nil {
		return fail(typex.Errorf(errors.Source, errors.InvalidArgument,
			"Invalid Body"))
	}

	body := buffer.Bytes()
	if len(body) < 1 {
		return fail(typex.Errorf(errors.Source, errors.InvalidArgument,
			"Invalid Body Length"))
	}

	var (
		request = schema.GetRootAsPutRequest(body, 0)
		score   = request.Score()
		maxSize = request.MaxSize()
		expiry  = request.Expiry()
	)
	if maxSize < 1 {
		return fail(typex.Errorf(errors.Source, errors.InvalidArgument,
			"Invalid MaxSize: %d", maxSize))
	}
	if expiry < 1 {
		return fail(typex.Errorf(errors.Source, errors.InvalidArgument,
			"Invalid expiry: %d", expiry))
	}

	var (
		num    = request.RecordsLength()
		result = make([]selectors.FieldTxnValue, num)
		fb     = pool.Get()
	)
	defer pool.Put(fb)

	for i := 0; i < num; i++ {
		record := &schema.PutRecord{}
		if !request.Records(record, i) {
			return fail(typex.Errorf(errors.Source, errors.InvalidArgument,
				"Invalid Record: %d", i))
		}

		id, err := readRecordId(record)
		if err != nil {
			return fail(err)
		}

		transaction, err := readRecordTransactionId(record)
		if err != nil {
			return fail(err)
		}

		fb.Reset()

		value, err := records.PutRecordFromSchemaToByte(fb, record)
		if err != nil {
			return fail(err)
		}

		result[i] = selectors.FieldTxnValue{
			Field: id,
			Txn:   transaction,
			Value: records.PackagePutRecord(value),
		}
	}

	return result, score, int64(maxSize), time.Duration(expiry), nil
}
