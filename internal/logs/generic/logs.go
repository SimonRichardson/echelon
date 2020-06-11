package generic

import (
	"os"

	"github.com/SimonRichardson/echelon/internal/logs"
	"github.com/SimonRichardson/echelon/internal/logs/plaintext"
)

func DefaultLog() logs.Log {
	L = plaintext.NewSync(os.Stdout)
	return L
}

var (
	L logs.Log
)
