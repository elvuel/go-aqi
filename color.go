package aqi

import (
	"fmt"
)

type Color struct {
	C, M, Y, K, R, G, B uint8
}

func (color Color) RGBToHex() string {
	return fmt.Sprintf("#%02X%02X%02X", color.R, color.G, color.B)
}

type MepColor struct {
	Name string
	Color
}
