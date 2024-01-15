package application

import "weather/domain"

type result struct {
	Date        string  `json:"date"`
	Temperature float64 `json:"avgTemp"`
	WindSpeed   float64 `json:"maxWind"`
	Humidity    float64 `json:"avgHumidity"`
	WillRain    float64 `json:"willRain"`
	WillSnow    float64 `json:"willSnow"`
	Condition   string  `json:"condition"`
	FeelsLike   float64 `json:"feelsLike"`
	Cloudiness  float64 `json:"cloudiness"`
}

func WeatherNow(cityName string) result {
	info := domain.WeatherInfo{City: cityName}
	temp := info.CurrentWeather(cityName)
	res := result{
		Temperature: temp["temperature"].(float64),
		FeelsLike:   temp["feelsLike"].(float64),
		WindSpeed:   temp["windSpeed"].(float64),
		Humidity:    temp["humidity"].(float64),
		Cloudiness:  temp["cloudiness"].(float64),
		Condition:   temp["condition"].(string),
	}
	return res
}

func WeatherForDay(cityName string) result {
	info := domain.WeatherInfo{City: cityName}
	temp := info.WeatherToDay(cityName)
	res := result{
		Temperature: temp["avgTemp"].(float64),
		WindSpeed:   temp["maxWind"].(float64),
		Humidity:    temp["avgHumidity"].(float64),
		WillRain:    temp["willRain"].(float64),
		WillSnow:    temp["willSnow"].(float64),
		Condition:   temp["condition"].(string),
	}
	return res
}

func WeatherFor3Days(cityName string) []result {
	info := domain.WeatherInfo{City: cityName}
	temp := info.WeatherNearThreeDays(cityName)
	var res []result
	for _, value := range temp {
		res = append(res, result{
			Date:        value["date"].(string),
			Temperature: value["avgTemp"].(float64),
			WindSpeed:   value["maxWind"].(float64),
			Humidity:    value["avgHumidity"].(float64),
			WillRain:    value["willRain"].(float64),
			WillSnow:    value["willSnow"].(float64),
			Condition:   value["condition"].(string)})
	}
	return res
}
