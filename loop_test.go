package loop

import (
	"fmt"
	"testing"
)

func TestLoop(t *testing.T) {
	fmt.Println(Range(1, 222, 3)(func(i int) {
		fmt.Println(i)
	}))
}
