package aqi

import (
	"testing"
)

func TestRGBToHex(t *testing.T) {
	color := Color{R: 63, G: 255, B: 118}
	if hex := color.RGBToHex(); hex != "#3FFF76" {
		t.Errorf("err = %s, want %s", hex, "#3FFF76")
	}
}
