package aqi

import (
	"errors"
	"reflect"
	"strconv"
)

type EpaPollutant struct {
	SO2Pollutant1H   float64 `json:"so2_1h" truncate:"0"` // ppb
	NO2Pollutant1H   float64 `json:"no2_1h" truncate:"0"` // ppb
	COPollutant8H    float64 `json:"co_8h" truncate:"1"`  // ppm
	O3Pollutant1H    float64 `json:"o3_1h" truncate:"3"`  // ppm
	O3Pollutant8H    float64 `json:"o3_8h" truncate:"3"`  // ppm
	PM10Pollutant24H float64 `json:"pm10_24h" truncate:"0"`
	PM25Pollutant24H float64 `json:"pm25_24h" truncate:"1"`
}

type EPABreakPoint struct {
	From, To float64
}

var (
	epaIAQIs          []EPABreakPoint
	epaColors         []EpaColor
	epaConcentrations map[string][]EPABreakPoint
	epaComputableMaxs map[string]float64
	epaTruncateRules  map[string]int
)

// initliaze all epa official suggests colors
func init() {
	epaColors = make([]EpaColor, 0)
	epaColors = append(epaColors,
		EpaColor{
			Name: "GREEN",
			Color: Color{
				R: 0, G: 228, B: 0, C: 40, M: 0, Y: 100, K: 0,
			},
		})
	epaColors = append(epaColors,
		EpaColor{
			Name: "YELLOW",
			Color: Color{
				R: 255, G: 255, B: 0, C: 0, M: 0, Y: 100, K: 0,
			},
		})
	epaColors = append(epaColors,
		EpaColor{
			Name: "ORANGE",
			Color: Color{
				R: 255, G: 126, B: 0, C: 0, M: 52, Y: 100, K: 0,
			},
		})
	epaColors = append(epaColors,
		EpaColor{
			Name: "RED",
			Color: Color{
				R: 255, G: 0, B: 0, C: 0, M: 100, Y: 100, K: 0,
			},
		})
	epaColors = append(epaColors,
		EpaColor{
			Name: "PURPLE",
			Color: Color{
				R: 153, G: 0, B: 76, C: 10, M: 100, Y: 40, K: 30,
			},
		})
	epaColors = append(epaColors,
		EpaColor{
			Name: "MAROON",
			Color: Color{
				R: 126, G: 0, B: 35, C: 30, M: 100, Y: 100, K: 30,
			},
		})

	epaIAQIs = []EPABreakPoint{
		EPABreakPoint{0, 50},
		EPABreakPoint{51, 100},
		EPABreakPoint{101, 150},
		EPABreakPoint{151, 200},
		EPABreakPoint{201, 300},
		EPABreakPoint{301, 400},
		EPABreakPoint{401, 500},
	}

	epaConcentrations = make(map[string][]EPABreakPoint)
	epaComputableMaxs = make(map[string]float64)

	epaConcentrations["o3_8h"] = []EPABreakPoint{
		//0, 0.059, 0.075, 0.095, 0.115, 0.374
		EPABreakPoint{0.000, 0.059},
		EPABreakPoint{0.060, 0.075},
		EPABreakPoint{0.076, 0.095},
		EPABreakPoint{0.096, 0.115},
		EPABreakPoint{0.116, 0.374},
	}

	// ? 0.124
	epaConcentrations["o3_1h"] = []EPABreakPoint{
		//0, 0, 0.124, 0.164, 0.204, 0.404, 0.504, 0.604
		EPABreakPoint{0, 0},
		EPABreakPoint{0, 0},
		EPABreakPoint{0.125, 0.164},
		EPABreakPoint{0.165, 0.204},
		EPABreakPoint{0.205, 0.404},
		EPABreakPoint{0.405, 0.504},
		EPABreakPoint{0.505, 0.604},
	}

	epaConcentrations["pm10_24h"] = []EPABreakPoint{
		//0, 54, 154, 254, 354, 424, 504, 604
		EPABreakPoint{0, 54},
		EPABreakPoint{55, 154},
		EPABreakPoint{155, 254},
		EPABreakPoint{255, 354},
		EPABreakPoint{355, 424},
		EPABreakPoint{425, 504},
		EPABreakPoint{505, 604},
	}

	epaConcentrations["pm25_24h"] = []EPABreakPoint{
		//0, 15.4, 40.4, 65.4, 150.4, 250.4, 350.4, 500.4
		EPABreakPoint{0.0, 15.4},
		EPABreakPoint{15.5, 40.4},
		EPABreakPoint{40.5, 65.4},
		EPABreakPoint{65.5, 150.4},
		EPABreakPoint{150.5, 250.4},
		EPABreakPoint{250.5, 350.4},
		EPABreakPoint{350.4, 500.4},
	}

	epaConcentrations["co_8h"] = []EPABreakPoint{
		//0, 4.4, 9.4, 12.4, 15.4, 30.4, 40.4, 50.4
		EPABreakPoint{0.0, 4.4},
		EPABreakPoint{4.5, 9.4},
		EPABreakPoint{9.5, 12.4},
		EPABreakPoint{12.5, 15.4},
		EPABreakPoint{15.5, 30.4},
		EPABreakPoint{30.5, 40.4},
		EPABreakPoint{40.5, 50.4},
	}

	epaConcentrations["so2_1h"] = []EPABreakPoint{
		//0, 35, 75, 185, 304, 604, 804, 1004
		EPABreakPoint{0, 35},
		EPABreakPoint{36, 75},
		EPABreakPoint{76, 185},
		EPABreakPoint{186, 304},
		EPABreakPoint{305, 604},
		EPABreakPoint{605, 804},
		EPABreakPoint{805, 1004},
	}

	epaConcentrations["no2_1h"] = []EPABreakPoint{
		//0, 53, 100, 360, 649, 1249, 1649, 2049
		EPABreakPoint{0, 53},
		EPABreakPoint{54, 100},
		EPABreakPoint{101, 360},
		EPABreakPoint{361, 649},
		EPABreakPoint{650, 1249},
		EPABreakPoint{1250, 1649},
		EPABreakPoint{1650, 2049},
	}

	for k, v := range epaConcentrations {
		epaComputableMaxs[k] = v[len(v)-1].To
	}

	epaTruncateRules = GetEPATruncateRules()
}

func epaPollutantCalculable(pollutant string) bool {
	return epaComputableMaxs[pollutant] > 0
}

func GetEPATruncateRules() map[string]int {
	epa := &EpaPollutant{}
	result := make(map[string]int)
	val := reflect.ValueOf(epa).Elem()
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		if tag := typeField.Tag.Get("json"); epaPollutantCalculable(tag) {
			t, err := strconv.ParseInt(typeField.Tag.Get("truncate"), 10, 0)
			if err != nil {
				result[tag] = 0
			} else {
				result[tag] = int(t)
			}
		}
	}
	return result
}

func GetEpaIAQI(pollutant string, concentration float64) (int, error) {
	if concentration <= 0 {
		return 0, nil
	}
	if !epaPollutantCalculable(pollutant) {
		return -1, errors.New("Invalid pollutant metric")
	} else {
		concentration = TruncateFloat(concentration, epaTruncateRules[pollutant])
		if concentration > epaComputableMaxs[pollutant] {
			switch pollutant {
			case "o3_8h":
				return -2, errors.New("Concentration value out of range")
			default:
				return 911, nil
			}
		} else {
			var bpLow, bpHigh, iaqiLow, iaqiHigh float64
			for i, point := range epaConcentrations[pollutant] {
				if concentration >= point.From && concentration <= point.To {
					bpLow = point.From
					bpHigh = point.To
					iaqiLow = epaIAQIs[i].From
					iaqiHigh = epaIAQIs[i].To
					break
				}
			}
			if (bpHigh - bpLow) == 0 {
				return -3, errors.New("Divided by 0???")
			}
			return int(Round(((iaqiHigh-iaqiLow)/(bpHigh-bpLow))*(concentration-bpLow)+iaqiLow, 0)), nil
		}
	}
	return 0, nil
}

func GetEpaPM25IAQI(concentration float64) (int, error) {
	return GetEpaIAQI("pm25_24h", concentration)
}

func GetEpaPM10IAQI(concentration float64) (int, error) {
	return GetEpaIAQI("pm10_24h", concentration)
}

func (epa *EpaPollutant) GetAllIAQI() map[string]int {
	//// reflect approach
	var result map[string]int
	result = make(map[string]int)
	val := reflect.ValueOf(epa).Elem()

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		valueField := val.Field(i)
		if tag := typeField.Tag.Get("json"); epaPollutantCalculable(tag) {
			v, ok := valueField.Interface().(float64)
			if !ok {
				v = -1
			}
			iaqi, _ := GetEpaIAQI(tag, v)
			result[tag] = iaqi
		}
	}
	return result
	//// encoding/json approach
	//value := make(map[string]float64)
	//result := make(map[string]int)
	//data, err := json.Marshal(epa)
	//if err != nil {
	//	return result
	//}
	//err = json.Unmarshal(data, &value)
	//if err != nil {
	//	return result
	//}
	//for pollutant, concentration := range value {
	//	if epaPollutantCalculable(pollutant) {
	//		v, err := GetEpaIAQI(pollutant, concentration)
	//		if err == nil {
	//			result[pollutant] = v
	//		}
	//	}
	//}
	//return result
}

func (epa *EpaPollutant) GetAQI() int {
	var result int
	result = -911
	allIAQI := epa.GetAllIAQI()
	for _, v := range allIAQI {
		if v >= result {
			result = v
		}
	}
	return result
}

func (epa *EpaPollutant) ResponsiblePollutants() []string {
	var result []string
	var max int
	result = make([]string, 0)
	allIAQI := epa.GetAllIAQI()
	for _, v := range allIAQI {
		if v >= max {
			max = v
		}
	}
	for k, v := range allIAQI {
		if v == max {
			result = append(result, k)
		}
	}
	return result
}
