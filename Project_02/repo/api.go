package repo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	domain "weather/domain/interface"
)

type Db struct {
	wi domain.Weather
}

func (d *Db) CWeather(cityName string) map[string]interface{} {
	//прогноз погоды на данный момент(current)
	result, err := d.makeRequest(cityName)
	if err != nil {
		return nil
	}

	current := result["current"].(map[string]interface{})
	temperature := current["temp_c"].(float64)
	feelsLike := current["feelslike_c"].(float64)
	windSpeed := current["wind_kph"].(float64)
	humidity := current["humidity"].(float64)
	cloudiness := current["cloud"].(float64)
	condition := current["condition"].(map[string]interface{})["text"].(string)

	weatherNow := map[string]interface{}{
		"temperature": temperature,
		"feelsLike":   feelsLike,
		"windSpeed":   windSpeed,
		"humidity":    humidity,
		"cloudiness":  cloudiness,
		"condition":   condition,
	}

	return weatherNow

}
func (d *Db) WtD(cityName string) map[string]interface{} {
	//на сегодня

	result, err := d.makeRequest(cityName)
	if err != nil {
		return nil
	}

	day := result["forecast"].(map[string]interface{})["forecastday"].([]interface{})[0].(map[string]interface{})
	avgTemp := day["day"].(map[string]interface{})["avgtemp_c"].(float64)
	maxWind := day["day"].(map[string]interface{})["maxwind_kph"].(float64)
	avgHumidity := day["day"].(map[string]interface{})["avghumidity"].(float64)
	willRain := day["day"].(map[string]interface{})["daily_will_it_rain"].(float64)
	willSnow := day["day"].(map[string]interface{})["daily_will_it_snow"].(float64)
	condition := day["day"].(map[string]interface{})["condition"].(map[string]interface{})["text"].(string)

	weatherToday := map[string]interface{}{
		"avgTemp":     avgTemp,
		"maxWind":     maxWind,
		"avgHumidity": avgHumidity,
		"willRain":    willRain,
		"willSnow":    willSnow,
		"condition":   condition,
	}
	return weatherToday
}
func (d *Db) WNTD(cityName string) map[string]map[string]interface{} {
	//На ближайшие 3 дня
	result, err := d.makeRequest(cityName)
	if err != nil {
		return nil
	}

	forecastDays := result["forecast"].(map[string]interface{})["forecastday"].([]interface{})
	weatherForThreeDays := make(map[string]map[string]interface{})

	for i := 0; i < 3; i++ {
		day := forecastDays[i].(map[string]interface{})
		date := day["date"].(string)
		avgTemp := day["day"].(map[string]interface{})["avgtemp_c"].(float64)
		maxWind := day["day"].(map[string]interface{})["maxwind_kph"].(float64)
		avgHumidity := day["day"].(map[string]interface{})["avghumidity"].(float64)
		willRain := day["day"].(map[string]interface{})["daily_will_it_rain"].(float64)
		willSnow := day["day"].(map[string]interface{})["daily_will_it_snow"].(float64)
		condition := day["day"].(map[string]interface{})["condition"].(map[string]interface{})["text"].(string)

		weather := map[string]interface{}{
			"date":        date,
			"avgTemp":     avgTemp,
			"maxWind":     maxWind,
			"avgHumidity": avgHumidity,
			"willRain":    willRain,
			"willSnow":    willSnow,
			"condition":   condition,
		}

		weatherForThreeDays[date] = weather
	}

	return weatherForThreeDays
}

func (d *Db) makeRequest(city string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=f65922aa666c470e88b220206240401&q=%s&lang=ru&days=3", city)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var document interface{}
	err = json.Unmarshal(body, &document)
	if err != nil {
		return nil, err
	}
	return document.(map[string]interface{}), nil
}
