package loop

import (
	"log"
	"time"
)

func Loop(fn func() interface{}, efn func(err error)) (v interface{}) {
	for {
		if err := _KTry(func() {
			v = fn()
		}); err != nil {
			if _err := _KTry(efn, err); _err != nil {
				log.Fatalln(_err.(*_KErr).StackTrace())
			}
		}
	}
}

func Wait(fn func(dur time.Duration) bool) error {
	var _b = true
	for i := 0; _b; i++ {
		if err := _KTry(func() {
			_b = fn(time.Second * time.Duration(i))
		}); err != nil {
			return err
		}

		if !_b {
			return nil
		}

		time.Sleep(time.Second)
	}
	return nil
}

func Ticker(fn func(dur time.Time) uint) error {
	var _dur uint = 0
	for i := 0; ; i++ {
		if err := _KTry(func() {
			_dur = fn(time.Now())
		}).(*_KErr); err != nil {
			return err
		}

		if _dur < 0 {
			return nil
		}

		if _dur == 0 {
			_dur = 1
		}

		time.Sleep(time.Duration(_dur))
	}
}
