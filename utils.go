package aqi

import (
	"fmt"
	"math"
	"strconv"
)

func TruncateFloat(v float64, digit int) float64 {
	i := fmt.Sprintf("%"+fmt.Sprintf(".%df", digit), v)
	f, _ := strconv.ParseFloat(i, digit)
	return f
}

// see: https://github.com/DeyV/gotools
func Round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow

	if intermed < 0.0 {
		intermed -= 0.5
	} else {
		intermed += 0.5
	}
	rounder = float64(int64(intermed))

	return rounder / float64(pow)
}
