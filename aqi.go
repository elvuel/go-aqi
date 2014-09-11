// Package aqi calculate AQI
package aqi

import (
	"fmt"
)

const (
	Author  = "elvuel"
	Email   = "elvuel@gmail.com"
	Version = "0.1-beta"
	Licence = "MIT"
)

func init() {
	fmt.Printf(
		"Author: %s\nEmail: %s\nVersion: %s\nLicence: %s\n(Initialize information will be removed)\n\n",
		Author,
		Email,
		Version,
		Licence,
	)
}
