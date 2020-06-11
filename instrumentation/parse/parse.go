package parse

import (
	"io"
	"strings"
	"time"

	"github.com/SimonRichardson/echelon/common"
	"github.com/SimonRichardson/echelon/errors"
	"github.com/SimonRichardson/echelon/instrumentation"
	"github.com/SimonRichardson/echelon/instrumentation/multi"
	"github.com/SimonRichardson/echelon/instrumentation/noop"
	"github.com/SimonRichardson/echelon/instrumentation/plaintext"
	"github.com/SimonRichardson/echelon/instrumentation/prometheus"
	r "github.com/SimonRichardson/echelon/instrumentation/redis"
	"github.com/SimonRichardson/echelon/instrumentation/statsd"
	p "github.com/SimonRichardson/echelon/internal/redis"
	"github.com/SimonRichardson/echelon/internal/typex"
	"github.com/peterbourgon/g2s"
)

type InstrumentationOptions struct {
	StatsdAddress       string
	StatsdSampleRate    float32
	PlaintextWriter     io.Writer
	RedisAddress        string
	RedisBufferDuration time.Duration
	RedisTimeout        string
}

func ParseString(value string,
	options InstrumentationOptions,
) (instrumentation.Instrumentation, error) {
	parts := strings.Split(value, ";")
	switch common.StripWhitespace(strings.ToLower(parts[0])) {
	case "noop":
		return noop.New(), nil
	case "plaintext":
		return plaintext.New(options.PlaintextWriter), nil
	case "statsd":
		statter := g2s.Noop()
		if options.StatsdAddress != "" {
			var err error
			if statter, err = g2s.Dial("udp", options.StatsdAddress); err != nil {
				typex.Fatal(err)
			}
		}
		return statsd.New(statter, options.StatsdSampleRate), nil
	case "prometheus":
		return prometheus.New("echelon", time.Second*10), nil
	case "redis":
		host := options.RedisAddress
		if err := p.ValidRedisHost(host); err != nil {
			return nil, err
		}

		timeout := options.RedisTimeout
		connTimeout, routing, err := p.Parse(timeout, timeout, timeout, "hash", nil)
		if err != nil {
			return nil, err
		}

		return r.New(p.New(
			[]string{host},
			routing,
			connTimeout,
			100,
			nil,
		), options.RedisBufferDuration), nil
	case "multi":
		instruments := []instrumentation.Instrumentation{}
		for _, v := range parts[1:] {
			if instr, err := ParseString(v, options); err != nil {
				return noop.New(), err
			} else {
				instruments = append(instruments, instr)
			}
		}
		return multi.New(instruments...), nil
	}
	return noop.New(), typex.Errorf(errors.Source, errors.NoCaseFound, "Invalid instrumentation %q", value)
}
