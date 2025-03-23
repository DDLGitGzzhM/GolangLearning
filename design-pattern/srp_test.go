package design_pattern

import "testing"

func BenchmarkTempV1_Allocs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := tempV1{temp: 25.0}
		_ = t.ShowCelsius()
		_ = t.ShowFahrenheit()
	}
}

func BenchmarkTempV2_Allocs(b *testing.B) {
	display := TempDisplayed{}
	for i := 0; i < b.N; i++ {
		t := tempV2{temp: 25.0}
		_ = display.Celsius(t.GetTemp())
		_ = display.Fahrenheit(t.GetTemp())
	}
}
