package design_pattern

import "fmt"

type tempV1 struct {
	temp float64
}

func (t *tempV1) ShowCelsius() string {
	return fmt.Sprintf("%.1f°C", t.temp)
}

func (t *tempV1) ShowFahrenheit() string {
	return fmt.Sprintf("%.1f°F", t.temp*9/5+32)
}

type tempV2 struct {
	temp float64
}

func (t *tempV2) GetTemp() float64 {
	return t.temp
}

type TempDisplayed struct {
}

func (d *TempDisplayed) Celsius(temp float64) string {
	return fmt.Sprintf("%.1f°C", temp)
}

func (d *TempDisplayed) Fahrenheit(temp float64) string {
	return fmt.Sprintf("%.1f°F", temp*9/5+32)
}
