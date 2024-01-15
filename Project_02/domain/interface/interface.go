package _interface

type Weather interface {
	CurrentWeather(cityName string) map[string]interface{}
	WeatherToDay(cityName string) map[string]interface{}
	WeatherNearThreeDays(cityName string) map[string]interface{}
}
