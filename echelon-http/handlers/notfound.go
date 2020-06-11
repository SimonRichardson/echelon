package handlers

import (
	"net/http"

	"github.com/SimonRichardson/echelon/echelon-http/responses"
	"github.com/SimonRichardson/echelon/errors"
	"github.com/SimonRichardson/echelon/internal/typex"
)

func NotFound() func(http.ResponseWriter, *http.Request) {
	return handle(func(w http.ResponseWriter, r *http.Request) {
		responses.NotFound(w, r, typex.Errorf(errors.Source, errors.MissingContent,
			"Not Found"))
		return
	})
}
