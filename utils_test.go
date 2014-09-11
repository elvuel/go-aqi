package aqi

import (
	"testing"
)

func TestTruncateFloat(t *testing.T) {
	f := 10 / 3.0
	expections := []float64{3, 3.3, 3.33, 3.333, 3.3333}
	for i, expect := range expections {
		if v := TruncateFloat(f, i); v != expect {
			t.Errorf("err %f, expect %f", v, expect)
		}
	}
}
