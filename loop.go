package loop

import (
	"time"
)

func Loop(fn func(i int)) error {
	for i := 0; ; i++ {
		if err := _KTry(fn, i); err != nil {
			return err
		}
	}
}

func Range(args ...int) func(fn func(i int)) error {
	return func(fn func(i int)) error {
		if len(args) == 0 {
			return Loop(fn)
		}

		if len(args) == 1 {
			for i := 0; i < args[0]; i++ {
				if err := _KTry(fn, i); err != nil {
					return err
				}
			}
		}

		if len(args) == 2 {
			for i := args[0]; i < args[1]; i++ {
				if err := _KTry(fn, i); err != nil {
					return err
				}
			}
		}

		if len(args) == 3 {
			for i := args[0]; i < args[1]; i += args[2] {
				if err := _KTry(fn, i); err != nil {
					return err
				}
			}
		}

		return nil
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

func Ticker(fn func(dur time.Time) uint) error{
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
