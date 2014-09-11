package aqi

import (
	"errors"
	"reflect"
)

const (
	MepPrimaryPollutantClassified       = 50
	MepNonAttainmentPollutantClassified = 100
)

var (
	mepIAQIs          []float64
	mepColors         []MepColor
	mepConcentrations map[string][]float64
	mepComputableMaxs map[string]float64
)

type MepPollutant struct {
	SO2Pollutant24H  float64 `json:"so2_24h"`
	SO2Pollutant1H   float64 `json:"so2_1h"`
	NO2Pollutant24H  float64 `json:"no2_24h"`
	NO2Pollutant1H   float64 `json:"no2_1h"`
	COPollutant24H   float64 `json:"co_24h"` // mg/m³ (mass concentraion)
	COPollutant1H    float64 `json:"co_1h"`  // mg/m³ (mass concentraion)
	O3Pollutant1H    float64 `json:"o3_1h"`
	O3Pollutant8H    float64 `json:"o3_8h"`
	PM10Pollutant24H float64 `json:"pm10_24h"`
	PM25Pollutant24H float64 `json:"pm25_24h"`
}

// initliaze all mep official suggests colors
func init() {
	mepColors = make([]MepColor, 0)
	mepColors = append(mepColors,
		MepColor{
			Name: "GREEN",
			Color: Color{
				R: 0, G: 228, B: 0, C: 40, M: 0, Y: 100, K: 0,
			},
		})
	mepColors = append(mepColors,
		MepColor{
			Name: "YELLOW",
			Color: Color{
				R: 255, G: 255, B: 0, C: 0, M: 0, Y: 100, K: 0,
			},
		})
	mepColors = append(mepColors,
		MepColor{
			Name: "ORANGE",
			Color: Color{
				R: 255, G: 126, B: 0, C: 0, M: 52, Y: 100, K: 0,
			},
		})
	mepColors = append(mepColors,
		MepColor{
			Name: "RED",
			Color: Color{
				R: 255, G: 0, B: 0, C: 0, M: 100, Y: 100, K: 0,
			},
		})
	mepColors = append(mepColors,
		MepColor{
			Name: "PURPLE",
			Color: Color{
				R: 153, G: 0, B: 76, C: 10, M: 100, Y: 40, K: 30,
			},
		})
	mepColors = append(mepColors,
		MepColor{
			Name: "MAROON",
			Color: Color{
				R: 126, G: 0, B: 35, C: 30, M: 100, Y: 100, K: 30,
			},
		})

	mepIAQIs = []float64{0, 50, 100, 150, 200, 300, 400, 500}

	mepConcentrations = make(map[string][]float64)
	mepComputableMaxs = make(map[string]float64)

	mepConcentrations["so2_24h"] = []float64{0, 50, 150, 475, 800, 1600, 2100, 2620}
	mepConcentrations["so2_1h"] = []float64{0, 150, 500, 650, 800}
	mepConcentrations["no2_24h"] = []float64{0, 40, 80, 180, 280, 565, 750, 940}
	mepConcentrations["no2_1h"] = []float64{0, 100, 200, 700, 1200, 2340, 3090, 3840}
	mepConcentrations["co_24h"] = []float64{0, 2, 4, 14, 24, 36, 48, 60}
	mepConcentrations["co_1h"] = []float64{0, 5, 10, 35, 60, 90, 120, 150}
	mepConcentrations["o3_1h"] = []float64{0, 160, 200, 300, 400, 800, 1000, 1200}
	mepConcentrations["o3_8h"] = []float64{0, 100, 160, 215, 265, 800}
	mepConcentrations["pm10_24h"] = []float64{0, 50, 150, 250, 350, 420, 500, 600}
	mepConcentrations["pm25_24h"] = []float64{0, 35, 75, 115, 150, 250, 350, 500}

	for k, v := range mepConcentrations {
		mepComputableMaxs[k] = v[len(v)-1]
	}
}

func (mep *MepPollutant) GetAQI() float64 {
	var result float64
	result = -1
	allIAQI := mep.GetAllIAQI()
	for _, v := range allIAQI {
		if v >= result {
			result = v
		}
	}
	return result
}

func (mep *MepPollutant) GetAllIAQI() map[string]float64 {
	var result map[string]float64
	result = make(map[string]float64)
	val := reflect.ValueOf(mep).Elem()

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		valueField := val.Field(i)
		if tag := typeField.Tag.Get("json"); mepPollutantCalculable(tag) {
			v, ok := valueField.Interface().(float64)
			if !ok {
				v = -1
			}
			iaqi, _ := GetMepIAQI(tag, v)
			result[tag] = iaqi
		}
	}
	return result
}

func (mep *MepPollutant) PrimaryPollutants() []string {
	var result []string
	var max float64
	max = 0
	result = make([]string, 0)
	allIAQI := mep.GetAllIAQI()
	for _, v := range allIAQI {
		if v >= max {
			max = v
		}
	}
	if max > MepPrimaryPollutantClassified {
		for k, v := range allIAQI {
			if v == max {
				result = append(result, k)
			}
		}
	}
	return result
}

func (mep *MepPollutant) NonAttainmentPollutants() []string {
	var result []string
	result = make([]string, 0)
	allIAQI := mep.GetAllIAQI()
	for k, v := range allIAQI {
		if v > MepNonAttainmentPollutantClassified {
			result = append(result, k)
		}
	}
	return result
}

func GetMepPM25IAQI(concentration float64) (float64, error) {
	return GetMepIAQI("pm25_24h", concentration)
}

func GetMepPM10IAQI(concentration float64) (float64, error) {
	return GetMepIAQI("pm10_24h", concentration)
}

func GetMepIAQI(pollutant string, concentration float64) (float64, error) {
	if concentration == 0 {
		return 0, nil
	}
	if !mepPollutantCalculable(pollutant) {
		return -1, errors.New("Invalid pollutant metric")
	} else {
		if concentration >= mepComputableMaxs[pollutant] {
			if concentration == mepComputableMaxs[pollutant] {
				return mepIAQIs[len(mepConcentrations[pollutant])-1], nil
			} else {
				switch pollutant {
				case "so2_1h", "o3_8h":
					return -2, errors.New("Concentration value out of range")
				default:
					return 911, nil
				}
			}
		} else {
			var bpLow, bpHigh, iaqiLow, iaqiHigh float64
			for i, v := range mepConcentrations[pollutant] {
				if concentration <= v {
					bpLow = mepConcentrations[pollutant][i-1]
					bpHigh = v
					iaqiLow = mepIAQIs[i-1]
					iaqiHigh = mepIAQIs[i]
					break
				}
			}
			if (bpHigh - bpLow) == 0 {
				return -3, errors.New("Divided by 0???")
			}
			return ((iaqiHigh-iaqiLow)/(bpHigh-bpLow))*(concentration-bpLow) + iaqiLow, nil
		}
	}
	return 0, nil
}

func mepPollutantCalculable(pollutant string) bool {
	return mepComputableMaxs[pollutant] > 0
}
