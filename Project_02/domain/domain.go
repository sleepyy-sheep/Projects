package domain

import "weather/repo"

type WeatherInfo struct {
	City string // Название города, где смотрим погоду

}

func (w *WeatherInfo) CurrentWeather(cityName string) map[string]interface{} {
	a := repo.Db{}
	return a.CWeather(cityName)
}
func (w *WeatherInfo) WeatherToDay(cityName string) map[string]interface{} {
	a := repo.Db{}
	return a.WtD(cityName)
}

func (w *WeatherInfo) WeatherNearThreeDays(cityName string) map[string]map[string]interface{} {
	a := repo.Db{}
	return a.WNTD(cityName)
}
