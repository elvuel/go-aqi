package aqi

import (
	"testing"
)

var (
	validEPAPollutants = []string{"so2_1h", "no2_1h", "co_8h", "o3_1h", "o3_8h", "pm10_24h", "pm25_24h"}
)

func BenchmarkEpaGetAQI(b *testing.B) {
	for n := 0; n < b.N; n++ {
		epa := &EpaPollutant{
			COPollutant8H:    8.4,
			O3Pollutant8H:    0.08742,
			PM25Pollutant24H: 40.9,
		}
		epa.GetAQI()
	}
}

func BenchmarkEpaGetEpaPM25IAQI(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetEpaPM25IAQI(130)
	}
}

func TestEpaPollutantCalculable(t *testing.T) {
	for _, v := range validEPAPollutants {
		if !epaPollutantCalculable(v) {
			t.Errorf("%s should be calculable", v)
		}
	}
	for _, v := range []string{"foo", "bar"} {
		if epaPollutantCalculable(v) {
			t.Errorf("%s should not be calculable", v)
		}
	}
}

func TestGetEPATruncateRules(t *testing.T) {
	pattern := make(map[string]int)
	pattern["pm25_24h"] = 1
	pattern["so2_1h"] = 0
	pattern["no2_1h"] = 0
	pattern["co_8h"] = 1
	pattern["o3_1h"] = 3
	pattern["o3_8h"] = 3
	pattern["pm10_24h"] = 0

	result := GetEPATruncateRules()
	for k, v := range result {
		if pattern[k] != v {
			t.Errorf("Rule %s err %d, want %d", k, v, pattern[k])
		}
	}
}

func TestGetEpaIAQI(t *testing.T) {
	iaqi, err := GetEpaIAQI("foo", 0)
	if err != nil {
		t.Error("fake foo pollutant should not raise exception")
	}
	if iaqi != 0 {
		t.Error("fake foo pollutant with 0 concentration should return 0")
	}

	iaqi, err = GetEpaIAQI("pm25_24h", 0)
	if err != nil {
		t.Error("pm25_24h pollutant should not raise exception")
	}

	if iaqi != 0 {
		t.Error("pm25_24h pollutant with 0 concentration should return 0")
	}

	iaqi, err = GetEpaIAQI("foo", 42)
	if err == nil {
		t.Error("fake foo pollutant should raise exception")
	}
	if iaqi != -1 {
		t.Error("fake foo pollutant with gt 0 value should return -1")
	}

	for _, pollutant := range validEPAPollutants {
		max := epaComputableMaxs[pollutant]
		overFlow := max + 5
		iaqi, err = GetEpaIAQI(pollutant, overFlow)
		if pollutant == "o3_8h" {
			if err == nil {
				t.Errorf("%s with max %f + 1(%f) should raise exception", pollutant, max, overFlow)
			}
			if iaqi != -2 {
				t.Errorf("%s with max %f + 1(%f) should return -2", pollutant, max, overFlow)
			}
		} else {
			if err != nil {
				t.Errorf("%s with max %f + 1(%f) should not raise exception", pollutant, max, overFlow)
			}
			if iaqi != 911 {
				t.Errorf("%s with max %f + 1(%f) should return 911, but return %d", pollutant, max, overFlow, iaqi)
			}
		}
	} // end for

	// spec seeds from EPA-454/B-12-001 Technical Assistance Document
	type Seed struct {
		Pollutant     string
		Concentration float64
		Expection     int
	}

	seeds := make([]Seed, 0)
	seeds = append(seeds,
		Seed{"o3_8h", 0.08742, 129}, Seed{"o3_8h", 0.077, 104},
		Seed{"pm25_24h", 40.9, 102}, Seed{"co_8h", 8.4, 90},
		Seed{"o3_8h", 0.141, 211},
	)

	for _, seed := range seeds {
		v, _ := GetEpaIAQI(seed.Pollutant, seed.Concentration)
		if v != seed.Expection {
			t.Errorf("%s with %f should return %d, but %d", seed.Pollutant, seed.Concentration, seed.Expection, v)
		}
	}
}

func TestEpaPollutantGetAllIAQI(t *testing.T) {
	epa := &EpaPollutant{
		COPollutant8H:    8.4,
		O3Pollutant8H:    0.08742,
		PM25Pollutant24H: 40.9,
	}

	result := epa.GetAllIAQI()
	nonZeroPollutants := []string{"pm25_24h", "o3_8h", "co_8h"}

	for _, v := range nonZeroPollutants {
		if result[v] < 0 {
			t.Errorf("want > 0 actually %f", result[v])
		}
	}
}

func TestEpaGetAQI(t *testing.T) {
	epa := &EpaPollutant{
		COPollutant8H:    8.4,
		O3Pollutant8H:    0.08742,
		PM25Pollutant24H: 40.9,
	}

	if epa.GetAQI() != 129 {
		t.Error("should pass")
	}
}

func TestEpaResponsiblePollutants(t *testing.T) {
	epa := &EpaPollutant{
		COPollutant8H:    8.4,
		O3Pollutant8H:    0.08742,
		PM25Pollutant24H: 40.9,
	}
	result := epa.ResponsiblePollutants()
	if len(result) != 1 {
		t.Error("should be o3_8h, length 1")
	}
	if result[0] != "o3_8h" {
		t.Error("should be o3_8h")
	}
	epa1 := &EpaPollutant{
		COPollutant8H:    8.4,
		O3Pollutant8H:    0.08742,
		PM25Pollutant24H: 54.9,
	}
	result = epa1.ResponsiblePollutants()
	if len(result) != 2 {
		t.Error("should be o3_8h and pm25_24h, length 2")
	}
}
