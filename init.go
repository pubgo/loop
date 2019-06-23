package loop

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLog() {
	log.Logger = log.Output(zerolog.NewConsoleWriter()).With().Str("module", "loop").Caller().Logger()
}

var Cfg = struct {
	Debug bool
}{
	Debug: true,
}
