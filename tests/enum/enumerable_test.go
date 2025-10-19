package enum

import (
	"fmt"
	"math"
	"testing"

	"github.com/pauldin91/go_ds/src/enum"
)

var data = []float32{0.01, 0.05, 0.1, 0.9, 0.95, 0.99}

func TestMap(t *testing.T) {
	var transformed = enum.Map(data, func(elem float32) string { return fmt.Sprintf("%.2f", elem) })
	for i, s := range transformed {
		if fmt.Sprintf("%.2f", data[i]) != s {
			t.Errorf("Actual at %d : %s, Expected : %.2f\n", i, s, data[i])
		}
	}

}

func TestFilter(t *testing.T) {
	var filtered = enum.Filter(data, func(elem float32) bool { return elem > 0.8 })
	for i, s := range filtered {
		if s <= 0.8 {
			t.Errorf("Actual at %d : %0.2f, Expected : > 0.8", i, s)
		}
	}
}

func TestReduce(t *testing.T) {

	var pairs = make(map[float32]float32)
	enum.Reduce(data, func(elem float32) {
		_, ok := pairs[elem]
		rounded := math.Round(float64(1.0-elem)*100) / 100
		_, rev := pairs[float32(rounded)]
		if !ok && !rev {
			pairs[elem] = 1 - elem
		}
	})

	expected := map[float32]float32{0.1: 0.9, 0.01: 0.99, 0.05: 0.95}
	for i, s := range pairs {
		if _, ok := expected[i]; !ok {
			t.Errorf("Found at %.2f : %.2f", i, s)
		}
	}

}
