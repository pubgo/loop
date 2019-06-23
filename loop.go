package loop

import (
	"github.com/pubgo/errors"
	"time"
)

func Wait(fn func(dur time.Duration) bool) {
	defer errors.Handle(func() {})

	var _b = true
	for i := 0; _b; i++ {
		errors.ErrHandle(errors.Try(func() {
			_b = fn(time.Second * time.Duration(i))
		}), func(err *errors.Err) {
			if Cfg.Debug {
				err.P()
			}
		})

		if !_b {
			return
		}

		time.Sleep(time.Second)
	}
}

func Ticker(fn func(dur time.Time) time.Duration) {
	defer errors.Handle(func() {})

	var _dur = time.Duration(0)
	for i := 0; ; i++ {
		errors.ErrHandle(errors.Try(func() {
			_dur = fn(time.Now())
		}), func(err *errors.Err) {
			if Cfg.Debug {
				err.P()
			}
		})

		if _dur < 0 {
			return
		}

		if _dur == 0 {
			_dur = time.Second
		}

		time.Sleep(_dur)
	}
}
