package loop_test

import (
	"fmt"
	"github.com/pubgo/assert"
	"github.com/pubgo/loop"
	"testing"
)

func TestLoop(t *testing.T) {
	fmt.Println(loop.Loop(func() interface{} {
		panic("sss")

		return nil
	}, func(err error) {
		assert.ErrWrap(err,"ss999")
	}))
}
