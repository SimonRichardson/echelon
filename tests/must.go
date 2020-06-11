package tests

import (
	"encoding/json"

	"github.com/SimonRichardson/echelon/internal/typex"
)

func MustMarshal(data interface{}) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		typex.Fatal(err)
	}
	return bytes
}

func MustUnmarshal(bytes []byte, data interface{}) {
	if err := json.Unmarshal(bytes, data); err != nil {
		typex.Fatal(err)
	}
}
