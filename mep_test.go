package aqi

import (
	"testing"
)

var (
	validMEPPollutants = []string{"so2_24h", "so2_1h", "no2_24h", "no2_1h", "co_24h", "co_1h", "o3_1h", "o3_8h", "pm10_24h", "pm25_24h"}
)

func BenchmarkMepGetAQI(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mep := &MepPollutant{
			PM25Pollutant24H: 44,
			PM10Pollutant24H: 65,
			COPollutant24H:   1.131,
			NO2Pollutant24H:  32,
			O3Pollutant1H:    93,
			O3Pollutant8H:    104,
			SO2Pollutant24H:  9,
		}
		mep.GetAQI()
	}
}

func BenchmarkMepGetMepPM25IAQI(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetMepPM25IAQI(130)
	}
}

func TestMepPollutantCalculable(t *testing.T) {
	for _, v := range validMEPPollutants {
		if !mepPollutantCalculable(v) {
			t.Errorf("%s should be calculable", v)
		}
	}
	for _, v := range []string{"foo", "bar"} {
		if mepPollutantCalculable(v) {
			t.Errorf("%s should not be calculable", v)
		}
	}
}

func TestGetMepIAQI(t *testing.T) {

	err, iaqi := GetMepIAQI("foo", 0)
	if err != nil {
		t.Error("fake foo pollutant should not raise exception")
	}
	if iaqi != 0 {
		t.Error("fake foo pollutant with 0 concentration should return 0")
	}

	err, iaqi = GetMepIAQI("pm25_24h", 0)
	if err != nil {
		t.Error("pm25_24h pollutant should not raise exception")
	}

	if iaqi != 0 {
		t.Error("pm25_24h pollutant with 0 concentration should return 0")
	}
	err, iaqi = GetMepIAQI("foo", 42)
	if err == nil {
		t.Error("fake foo pollutant should raise exception")
	}
	if iaqi != -1 {
		t.Error("fake foo pollutant with gt 0 value should return -1")
	}

	for _, v := range validMEPPollutants {
		points := mepConcentrations[v]
		// with maxi check
		max := points[len(points)-1]
		err, iaqi := GetMepIAQI(v, max)
		if err != nil {
			t.Errorf("%s with max %f should not raise exception", v, max)
		}
		if v == "o3_8h" || v == "so2_1h" {
			if iaqi != mepIAQIs[len(mepConcentrations[v])-1] {
				t.Errorf("%s with max %f, the iaqi should be %f", v, max, mepIAQIs[len(mepConcentrations[v])-1])
			}
		} else {
			if iaqi != mepIAQIs[len(mepIAQIs)-1] {
				t.Errorf("%s with max %f, the iaqi should be %f", v, max, mepIAQIs[len(mepIAQIs)-1])
			}
		}
		overFlow := max + 1
		err, iaqi = GetMepIAQI(v, max+1)
		if v == "o3_8h" || v == "so2_1h" {
			if err == nil {
				t.Errorf("%s with max %f + 1(%f) should raise exception", v, max, overFlow)
			}
			if iaqi != -2 {
				t.Errorf("%s with max %f + 1(%f) should return -2", v, max, overFlow)
			}
		} else {
			if err != nil {
				t.Errorf("%s with max %f + 1(%f) should not raise exception", v, max, overFlow)
			}
			if iaqi != 911 {
				t.Errorf("%s with max %f + 1(%f) should return 911", v, max, overFlow)
			}
		}
	} // end for

	err, iaqi = GetMepIAQI("pm25_24h", 42)
	if err != nil {
		t.Error("pm25_24h with 42 concentration should not raise exception")
	}
	if iaqi != 58.75 {
		t.Errorf("pm25_24h with 42 concentration should return %f, but get %f", 58.75, iaqi)
	}
}

func TestMepPollutantGetAllIAQI(t *testing.T) {
	mep := &MepPollutant{
		PM25Pollutant24H: 44,
		PM10Pollutant24H: 65,
		COPollutant24H:   1.131,
		NO2Pollutant24H:  32,
		O3Pollutant1H:    93,
		O3Pollutant8H:    104,
		SO2Pollutant24H:  9,
	}
	result := mep.GetAllIAQI()
	nonZeroPollutants := []string{"so2_24h", "no2_24h", "o3_8h", "pm10_24h", "pm25_24h", "co_24h", "o3_1h"}
	for _, v := range nonZeroPollutants {
		if result[v] <= 0 {
			t.Errorf("want > 0 actually %f", result[v])
		}
	}
}

func TestMepPrimaryPollutants(t *testing.T) {
	mep := &MepPollutant{
		PM25Pollutant24H: 44,
		PM10Pollutant24H: 65,
		COPollutant24H:   1.131,
		NO2Pollutant24H:  32,
		O3Pollutant1H:    93,
		O3Pollutant8H:    104,
		SO2Pollutant24H:  9,
	}
	result := mep.PrimaryPollutants()
	if len(result) != 1 {
		t.Error("should be pm25_24h, length 1")
	}
	if result[0] != "pm25_24h" {
		t.Error("should be pm25_24h")
	}
	mep1 := &MepPollutant{
		PM25Pollutant24H: 37,
		PM10Pollutant24H: 55,
		COPollutant24H:   0.436,
		NO2Pollutant24H:  22,
		O3Pollutant1H:    107,
		O3Pollutant8H:    115,
		SO2Pollutant24H:  14,
	}
	result = mep1.PrimaryPollutants()
	if len(result) != 1 {
		t.Error("should be o3_8h, length 1")
	}
	if result[0] != "o3_8h" {
		t.Error("should be o3_8h")
	}
	mep2 := &MepPollutant{
		PM25Pollutant24H: 75,
		PM10Pollutant24H: 150,
	}
	result = mep2.PrimaryPollutants()
	if len(result) != 2 {
		t.Errorf("length of result should be 2")
	}
}

func TestMepNonAttainmentPollutants(t *testing.T) {
	mep := &MepPollutant{
		PM25Pollutant24H: 115,
	}
	result := mep.NonAttainmentPollutants()
	if len(result) != 1 {
		t.Error("pm25_24h with 115 should be non attainment pollutant")
	}
	if result[0] != "pm25_24h" {
		t.Error("pm25_24h with 115 should be non attainment pollutant")
	}
	mep1 := &MepPollutant{
		PM25Pollutant24H: 115,
		PM10Pollutant24H: 350,
	}
	result = mep1.NonAttainmentPollutants()
	if len(result) != 2 {
		t.Errorf("length of result should be 2")
	}
}

func TestMepGetAQI(t *testing.T) {
	mep := &MepPollutant{
		PM25Pollutant24H: 44,
		PM10Pollutant24H: 65,
		COPollutant24H:   1.131,
		NO2Pollutant24H:  32,
		O3Pollutant1H:    93,
		O3Pollutant8H:    104,
		SO2Pollutant24H:  9,
	}

	mep1 := &MepPollutant{
		PM25Pollutant24H: 82,
		PM10Pollutant24H: 113,
		COPollutant24H:   0.948,
		NO2Pollutant24H:  28,
		O3Pollutant1H:    85,
		O3Pollutant8H:    101,
		SO2Pollutant24H:  10,
	}
	if mep.GetAQI() != 61.25 {
		t.Error("should pass")
	}
	if mep1.GetAQI() != 108.75 {
		t.Error("should pass")
	}
}
