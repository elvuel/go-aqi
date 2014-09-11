package main

import (
	"fmt"

	aqi "github.com/elvuel/go-aqi"
)

func main() {
	mep := &aqi.MepPollutant{
		PM25Pollutant24H: 44,
		PM10Pollutant24H: 65,
		COPollutant24H:   1.131,
		NO2Pollutant24H:  32,
		O3Pollutant1H:    93,
		O3Pollutant8H:    104,
		SO2Pollutant24H:  9,
	}
	fmt.Println(mep.GetAQI())

	v, _ := aqi.GetEpaIAQI("pm25_24h", 54.9)
	fmt.Println(v)

}
