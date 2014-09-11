package aqi

import (
	"fmt"
	"strconv"
)

func TruncateFloat(v float64, digit int) float64 {
	i := fmt.Sprintf("%"+fmt.Sprintf(".%df", digit), v)
	f, _ := strconv.ParseFloat(i, digit)
	return f
}
