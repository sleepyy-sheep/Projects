package main

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"weather/application"
)

var bot *tgbotapi.BotAPI
var chatid int64
var location string

const TOKEN = "6686593013:AAG7mbBYUHklvQYQIsX09B-LpuE6F9UV7Wc"

func getWeatherNow(cityName string) {
	weather := application.WeatherNow(cityName)

	if weather.Condition != "" {
		msg := fmt.Sprintf("Текущая погода в городе %s:\n\nТемпература: %.1f°C\nОщущается как: %.1f°C\nСкорость ветра: %.1f км/ч\nВлажность: %.1f%%\nОблачность: %.1f%%\nУсловия: %s", cityName, weather.Temperature, weather.FeelsLike, weather.WindSpeed, weather.Humidity, weather.Cloudiness, weather.Condition)

		sendMessage(msg)
	} else {
		sendMessage("Не удалось получить информацию о погоде. Проверьте название города и попробуйте снова.")
	}
}

func getWeatherToday(cityName string) {
	weather := application.WeatherForDay(cityName)

	if weather.Condition != "" {
		msg := fmt.Sprintf("Погода на сегодня в городе %s:\n\nСредняя температура: %.1f°C\nМаксимальная скорость ветра: %.1f км/ч\nСредняя влажность: %.1f%%\nВероятность дождя: %.1f%%\nВероятность снега: %.1f%%\nУсловия: %s", cityName, weather.Temperature, weather.WindSpeed, weather.Humidity, weather.WillRain, weather.WillSnow, weather.Condition)

		sendMessage(msg)
	} else {
		sendMessage("Не удалось получить информацию о погоде. Проверьте название города и попробуйте снова.")
	}
}
func getWeatherForThreeDays(cityName string) {
	weather := application.WeatherFor3Days(cityName)

	if weather != nil {
		msg := fmt.Sprintf("Погода на ближайшие 3 дня в городе %s:\n\n", cityName)

		for _, w := range weather {
			msg += fmt.Sprintf("Дата: %s\nСредняя температура: %.1f°C\nМаксимальная скорость ветра: %.1f км/ч\nСредняя влажность: %.1f%%\nВероятность дождя: %.1f%%\nВероятность снега: %.1f%%\nУсловия: %s\n\n", w.Date, w.Temperature, w.WindSpeed, w.Humidity, w.WillRain, w.WillSnow, w.Condition)
		}

		sendMessage(msg)
	} else {
		sendMessage("Не удалось получить информацию о погоде. Проверьте название города и попробуйте снова.")
	}
}

func handleWeatherOption(option string, cityName string) {
	switch option {
	case "Погода сейчас":
		getWeatherNow(cityName)
	case "Погода на сегодня":
		getWeatherToday(cityName)
	case "Погода на 3 дня":
		getWeatherForThreeDays(cityName)
	default:
		sendChoosebutton(chatid)
	}
}

func sendStartbutton(chatID int64) {
	btn1 := tgbotapi.NewKeyboardButton("Указать город")
	btn2 := tgbotapi.NewKeyboardButton("Мое местоположение")

	row1 := tgbotapi.NewKeyboardButtonRow(btn1, btn2)

	msg := tgbotapi.NewMessage(chatID, "Выбери вариант:")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(row1)
	bot.Send(msg)
}

func sendChoosebutton(chatID int64) {
	btn1 := tgbotapi.NewKeyboardButton("Погода сейчас")
	btn2 := tgbotapi.NewKeyboardButton("Погода на сегодня")
	btn3 := tgbotapi.NewKeyboardButton("Погода на 3 дня")

	row1 := tgbotapi.NewKeyboardButtonRow(btn1, btn2)
	row2 := tgbotapi.NewKeyboardButtonRow(btn3)

	msg := tgbotapi.NewMessage(chatID, "Выбери вариант:")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(row1, row2)
	bot.Send(msg)
}

func connectwithtg() {
	var err error
	bot, err = tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		panic("Cannot connect to Telegram")
	}
}
func sendMessage(msg string) {
	msgconfig := tgbotapi.NewMessage(chatid, msg)
	bot.Send(msgconfig)
}

func main() {

	connectwithtg()
	updateconfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateconfig) {
		if update.Message != nil && update.Message.Text == "/start" {
			chatid = update.Message.Chat.ID
			sendMessage("Привет, я твой бот для погоды")

			sendStartbutton(chatid)

		} else if update.Message != nil && update.Message.Text == "Мое местоположение" {
			sendMessage(" Для этого необходимо нажать на иконку скрепки и выбрать во всплывающем меню пункты «Геопозиция» или «Место».")

		} else if update.Message != nil && update.Message.Location != nil {

			latitudef := update.Message.Location.Latitude
			longitudef := update.Message.Location.Longitude
			latitude := strconv.FormatFloat(latitudef, 'f', -1, 64)
			longitude := strconv.FormatFloat(longitudef, 'f', -1, 64)
			location = latitude + "," + longitude

			sendMessage("Спасибо за предоставленное местоположение!")
			sendChoosebutton(chatid)

		} else if update.Message != nil && update.Message.Text == "Указать город" {
			sendMessage("Пожалуйста, укажите город:")

			continue
		} else if update.Message != nil && location == "" {
			location = update.Message.Text

			sendMessage("Спасибо за предоставленное местоположение!")
			sendChoosebutton(chatid)
		} else if update.Message != nil {
			handleWeatherOption(update.Message.Text, location)
		}
	}
}
