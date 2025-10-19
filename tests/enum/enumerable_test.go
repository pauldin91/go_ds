package enum

import (
	"fmt"
	"testing"

	"github.com/pauldin91/go_ds/src/enum"
)

func TestMap(t *testing.T) {
	var init = []float32{0.01, 0.05, 0.1, 0.9, 0.95, 0.99}
	var transformed = enum.Map(init, func(elem float32) string { return fmt.Sprintf("%.2f", elem) })
	for i, s := range transformed {
		if fmt.Sprintf("%.2f", init[i]) != s {
			t.Errorf("Actual at %d : %s, Expected : %.2f\n", i, s, init[i])
		}
	}

}
