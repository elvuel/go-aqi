package main

import (
	"fmt"

	aqi "github.com/elvuel/go-aqi"
)

func main() {
	mepSample()
	epaSample()
}

func mepSample() {
	fmt.Println("\n######## MEP ########")
	v, _ := aqi.GetMepIAQI("no2_1h", 210)
	fmt.Printf("NO2_IAQI:%#v\n", v)

	v, _ = aqi.GetMepPM25IAQI(25)
	fmt.Printf("PM25_IAQI:%#v\n", v)

	v, _ = aqi.GetMepPM10IAQI(40)
	fmt.Printf("PM10_IAQI:%#v\n", v)

	mep := &aqi.MepPollutant{
		PM25Pollutant24H: 80,
		PM10Pollutant24H: 65,
		COPollutant24H:   24,
		NO2Pollutant24H:  32,
		O3Pollutant1H:    93,
		O3Pollutant8H:    104,
		SO2Pollutant24H:  9,
	}
	fmt.Printf("MEP AQI: %#v\n", mep.GetAQI())
	fmt.Printf("MEP ALL IAQIS: %#v\n", mep.GetAllIAQI())
	fmt.Printf("MEP Responsible Pollutants: %#v\n", mep.ResponsiblePollutants())
	fmt.Printf("MEP Non Attainment Pollutants: %#v\n", mep.NonAttainmentPollutants())
}

func epaSample() {
	fmt.Println("\n######## EPA ########")
	v, _ := aqi.GetEpaIAQI("no2_1h", 210)
	fmt.Printf("NO2_IAQI:%#v\n", v)

	v, _ = aqi.GetEpaPM25IAQI(25)
	fmt.Printf("PM25_IAQI:%#v\n", v)

	v, _ = aqi.GetEpaPM10IAQI(40)
	fmt.Printf("PM10_IAQI:%#v\n", v)

	epa := &aqi.EpaPollutant{
		COPollutant8H:    8.4,
		O3Pollutant8H:    0.08742,
		PM25Pollutant24H: 40.9,
	}
	fmt.Printf("EPA AQI: %#v\n", epa.GetAQI())
	fmt.Printf("EPA ALL IAQIS: %#v\n", epa.GetAllIAQI())
	fmt.Printf("EPA Responsible Pollutants: %#v\n", epa.ResponsiblePollutants())
}
