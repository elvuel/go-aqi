package aqi

import (
	"errors"
	"reflect"
)

const (
	MepPrimaryPollutantClassified       = 50
	MepNonAttainmentPollutantClassified = 100
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

type MEPBreakPoint struct {
	From, To float64
}

var (
	mepIAQIs          []MEPBreakPoint
	mepColors         []MepColor
	mepConcentrations map[string][]MEPBreakPoint
	mepComputableMaxs map[string]float64
)

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

	mepIAQIs = []MEPBreakPoint{
		//0, 50, 100, 150, 200, 300, 400, 500
		MEPBreakPoint{0, 50},
		MEPBreakPoint{51, 100},
		MEPBreakPoint{101, 150},
		MEPBreakPoint{151, 200},
		MEPBreakPoint{201, 300},
		MEPBreakPoint{301, 400},
		MEPBreakPoint{401, 500},
	}

	mepConcentrations = make(map[string][]MEPBreakPoint)
	mepComputableMaxs = make(map[string]float64)

	mepConcentrations["so2_24h"] = []MEPBreakPoint{
		//0, 50, 150, 475, 800, 1600, 2100, 2620
		MEPBreakPoint{0, 50},
		MEPBreakPoint{51, 150},
		MEPBreakPoint{151, 475},
		MEPBreakPoint{476, 800},
		MEPBreakPoint{801, 1600},
		MEPBreakPoint{1601, 2100},
		MEPBreakPoint{2101, 2620},
	}

	mepConcentrations["so2_1h"] = []MEPBreakPoint{
		//0, 150, 500, 650, 800
		MEPBreakPoint{0, 50},
		MEPBreakPoint{51, 150},
		MEPBreakPoint{151, 500},
		MEPBreakPoint{501, 650},
		MEPBreakPoint{651, 800},
	}

	mepConcentrations["no2_24h"] = []MEPBreakPoint{
		//0, 40, 80, 180, 280, 565, 750, 940
		MEPBreakPoint{0, 40},
		MEPBreakPoint{41, 80},
		MEPBreakPoint{81, 180},
		MEPBreakPoint{181, 280},
		MEPBreakPoint{281, 565},
		MEPBreakPoint{566, 750},
		MEPBreakPoint{751, 940},
	}
	mepConcentrations["no2_1h"] = []MEPBreakPoint{
		//0, 100, 200, 700, 1200, 2340, 3090, 3840
		MEPBreakPoint{0, 100},
		MEPBreakPoint{101, 200},
		MEPBreakPoint{201, 700},
		MEPBreakPoint{701, 1200},
		MEPBreakPoint{1201, 2340},
		MEPBreakPoint{2341, 3090},
		MEPBreakPoint{3091, 3840},
	}
	mepConcentrations["co_24h"] = []MEPBreakPoint{
		//0, 2, 4, 14, 24, 36, 48, 60
		MEPBreakPoint{0, 2},
		MEPBreakPoint{3, 4},
		MEPBreakPoint{5, 14},
		MEPBreakPoint{15, 24},
		MEPBreakPoint{25, 36},
		MEPBreakPoint{37, 48},
		MEPBreakPoint{49, 60},
	}
	mepConcentrations["co_1h"] = []MEPBreakPoint{
		//0, 5, 10, 35, 60, 90, 120, 150
		MEPBreakPoint{0, 5},
		MEPBreakPoint{6, 10},
		MEPBreakPoint{11, 35},
		MEPBreakPoint{36, 60},
		MEPBreakPoint{61, 90},
		MEPBreakPoint{91, 120},
		MEPBreakPoint{121, 150},
	}
	mepConcentrations["o3_1h"] = []MEPBreakPoint{
		//0, 160, 200, 300, 400, 800, 1000, 1200
		MEPBreakPoint{0, 160},
		MEPBreakPoint{161, 200},
		MEPBreakPoint{201, 300},
		MEPBreakPoint{301, 400},
		MEPBreakPoint{401, 800},
		MEPBreakPoint{801, 1000},
		MEPBreakPoint{1001, 1200},
	}
	mepConcentrations["o3_8h"] = []MEPBreakPoint{
		//0, 100, 160, 215, 265, 800
		MEPBreakPoint{0, 100},
		MEPBreakPoint{101, 160},
		MEPBreakPoint{161, 215},
		MEPBreakPoint{216, 265},
		MEPBreakPoint{266, 800},
	}
	mepConcentrations["pm10_24h"] = []MEPBreakPoint{
		//0, 50, 150, 250, 350, 420, 500, 600
		MEPBreakPoint{0, 50},
		MEPBreakPoint{51, 150},
		MEPBreakPoint{151, 250},
		MEPBreakPoint{251, 350},
		MEPBreakPoint{351, 420},
		MEPBreakPoint{421, 500},
		MEPBreakPoint{501, 600},
	}
	mepConcentrations["pm25_24h"] = []MEPBreakPoint{
		//0, 35, 75, 115, 150, 250, 350, 500
		MEPBreakPoint{0, 35},
		MEPBreakPoint{36, 75},
		MEPBreakPoint{76, 115},
		MEPBreakPoint{116, 150},
		MEPBreakPoint{151, 250},
		MEPBreakPoint{251, 350},
		MEPBreakPoint{351, 500},
	}

	for k, v := range mepConcentrations {
		mepComputableMaxs[k] = v[len(v)-1].To
	}
}

func mepPollutantCalculable(pollutant string) bool {
	return mepComputableMaxs[pollutant] > 0
}

func GetMepIAQI(pollutant string, concentration float64) (int, error) {
	if concentration == 0 {
		return 0, nil
	}
	if !mepPollutantCalculable(pollutant) {
		return -1, errors.New("Invalid pollutant metric")
	} else {
		if concentration > mepComputableMaxs[pollutant] {
			switch pollutant {
			case "so2_1h", "o3_8h":
				return -2, errors.New("Concentration value out of range")
			default:
				return 911, nil
			}
		} else {
			var bpLow, bpHigh, iaqiLow, iaqiHigh float64
			for i, point := range mepConcentrations[pollutant] {
				if concentration >= point.From && concentration <= point.To {
					bpLow = point.From
					bpHigh = point.To
					iaqiLow = mepIAQIs[i].From
					iaqiHigh = mepIAQIs[i].To
					break
				}
			}
			if (bpHigh - bpLow) == 0 {
				return -3, errors.New("Divided by 0???")
			}
			roundValue := Round(((iaqiHigh-iaqiLow)/(bpHigh-bpLow))*(concentration-bpLow)+iaqiLow, 1)
			intValue := int(roundValue)
			if roundValue*10 > float64(intValue*10) {
				intValue += 1
			}
			return intValue, nil
		}
	}
	return 0, nil
}

func GetMepPM25IAQI(concentration float64) (int, error) {
	return GetMepIAQI("pm25_24h", concentration)
}

func GetMepPM10IAQI(concentration float64) (int, error) {
	return GetMepIAQI("pm10_24h", concentration)
}

func (mep *MepPollutant) GetAllIAQI() map[string]int {
	var result map[string]int
	result = make(map[string]int)
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

func (mep *MepPollutant) GetAQI() int {
	var result int
	result = -1
	allIAQI := mep.GetAllIAQI()
	for _, v := range allIAQI {
		if v >= result {
			result = v
		}
	}
	return result
}

func (mep *MepPollutant) ResponsiblePollutants() []string {
	var result []string
	var max int
	max = -911
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
