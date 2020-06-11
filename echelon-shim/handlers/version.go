package handlers

import (
	"net/http"
	"time"

	"github.com/SimonRichardson/echelon/echelon-shim/responses"
	"github.com/SimonRichardson/echelon/common"
)

func Version(version string) http.HandlerFunc {
	parseErr := common.ParseSemver(version)

	return handle(func(w http.ResponseWriter, r *http.Request) {
		began := time.Now()

		if parseErr != nil {
			responses.InternalServerError(w, r, parseErr)
			return
		}

		responses.OK(w, version, time.Since(began))
		return
	})
}
