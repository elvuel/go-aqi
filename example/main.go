package main

import (
	"fmt"

	aqi "github.com/elvuel/go-aqi"
)

func main() {
	mep := &aqi.MepPullutant{
		PM25Pollutant24H: 44,
		PM10Pollutant24H: 65,
		COPollutant24H:   1.131,
		NO2Pollutant24H:  32,
		O3Pollutant1H:    93,
		O3Pollutant8H:    104,
		SO2Pollutant24H:  9,
	}

	fmt.Println(mep.GetAQI())

}
