package main

import (
	"fmt"
	"schedule/database"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const TOKEN = "6437936182:AAF2KPhfMTKcVYklIT6Zr4mEuc7eMZKcB20"

var bot *tgbotapi.BotAPI
var chatid int64
var who string
var group string
var name string

var weektype string
var weekday int
var para int

func contains(arr []string, str string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}

func chetnechet(chatID int64) {
	btn1 := tgbotapi.NewKeyboardButton("Четная неделя")
	btn2 := tgbotapi.NewKeyboardButton("Нечетная неделя")
	row := tgbotapi.NewKeyboardButtonRow(btn1, btn2)
	msg := tgbotapi.NewMessage(chatID, "Выбери вариант:")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(row)
	bot.Send(msg)
}

func sendStudentWeekdays(chatID int64) {
	btn1 := tgbotapi.NewKeyboardButton("Понедельник")
	btn2 := tgbotapi.NewKeyboardButton("Вторник")
	btn3 := tgbotapi.NewKeyboardButton("Среда")
	btn4 := tgbotapi.NewKeyboardButton("Четверг")
	btn5 := tgbotapi.NewKeyboardButton("Пятница")
	btn6 := tgbotapi.NewKeyboardButton("Вернуться назад")

	row1 := tgbotapi.NewKeyboardButtonRow(btn1, btn2, btn3)
	row2 := tgbotapi.NewKeyboardButtonRow(btn4, btn5, btn6)

	msg := tgbotapi.NewMessage(chatID, "Выбери вариант:")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(row1, row2)
	bot.Send(msg)
}

func sendStudentDayS(chatID int64) {
	btn1 := tgbotapi.NewKeyboardButton("Какая следующая пара?")
	btn2 := tgbotapi.NewKeyboardButton("Расписание на определенный день недели")
	btn3 := tgbotapi.NewKeyboardButton("Расписание на сегодня")
	btn4 := tgbotapi.NewKeyboardButton("Расписание на завтра")
	btn5 := tgbotapi.NewKeyboardButton("Изменить тип недели")

	row1 := tgbotapi.NewKeyboardButtonRow(btn1, btn2)
	row2 := tgbotapi.NewKeyboardButtonRow(btn3, btn4)
	row3 := tgbotapi.NewKeyboardButtonRow(btn5)

	msg := tgbotapi.NewMessage(chatID, "Выбери вариант:")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(row1, row2, row3)
	bot.Send(msg)
}

func sendStudentGroup(chatID int64) {
	btn1 := tgbotapi.NewKeyboardButton("231-1")
	btn2 := tgbotapi.NewKeyboardButton("231-2")
	btn3 := tgbotapi.NewKeyboardButton("232-1")
	btn4 := tgbotapi.NewKeyboardButton("232-2")
	btn5 := tgbotapi.NewKeyboardButton("233-1")
	btn6 := tgbotapi.NewKeyboardButton("233-2")

	row1 := tgbotapi.NewKeyboardButtonRow(btn1, btn2, btn3)
	row2 := tgbotapi.NewKeyboardButtonRow(btn4, btn5, btn6)

	msg := tgbotapi.NewMessage(chatID, "Выбери вариант:")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(row1, row2)
	bot.Send(msg)
}

func sendKeyboard(chatID int64) {
	btn1 := tgbotapi.NewKeyboardButton("Студент")
	btn2 := tgbotapi.NewKeyboardButton("Преподaватель")
	row := tgbotapi.NewKeyboardButtonRow(btn1, btn2)
	msg := tgbotapi.NewMessage(chatID, "Выбери одну из кнопок:")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(row)
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
	arr := []string{"Понедельник", "Вторник", "Среда", "Четверг", "Пятница"}
	connectwithtg()
	updateconfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateconfig) {
		if update.Message != nil && update.Message.Text == "/start" {
			chatid = update.Message.Chat.ID
			sendMessage("Привет, я твой бот для расписания")
			chetnechet(chatid)
		} else if update.Message != nil && update.Message.Text == "Четная неделя" {
			sendKeyboard(chatid)
			weektype = "четная"
		} else if update.Message != nil && update.Message.Text == "Нечетная неделя" {
			sendKeyboard(chatid)
			weektype = "нечетная"
		} else if update.Message != nil && update.Message.Text == "Студент" {
			sendStudentGroup(chatid)
			who = "student"
		} else if update.Message != nil && update.Message.Text == "Преподaватель" {
			sendMessage("Введите своё Ф.И.О. в формате: ИвановИ.И.")
			who = "teacher"

		} else if strings.Count(string(update.Message.Text), ".") == 2 {
			name = string(update.Message.Text)
			sendStudentDayS(chatid)

		} else if update.Message != nil && (update.Message.Text == "231-1" || update.Message.Text == "231-2" || update.Message.Text == "232-1" || update.Message.Text == "232-2" || update.Message.Text == "233-1" || update.Message.Text == "233-2") {
			sendStudentDayS(chatid)
			group = string(update.Message.Text)
		} else if update.Message != nil && update.Message.Text == "Изменить тип недели" {
			if weektype == "четная" {
				weektype = "нечетная"
				sendMessage("Тип недели изменён на нечётную")
			} else {
				weektype = "четная"
				sendMessage("Тип недели изменён на чётную")
			}

		} else if update.Message != nil && update.Message.Text == "Расписание на сегодня" && who == "student" {
			weekday := int(time.Now().Weekday()) - 1

			convert := map[int]string{
				0: "понедельник",
				1: "вторник",
				2: "среда",
				3: "четверг",
				4: "пятница",
			}
			result := database.GetStudentsDataBy(group, convert[weekday], weektype)
			lessons := result.Lessons
			for i, info := range lessons {
				var k string
				para := info[0]
				teachername := info[1]
				adress := info[2]
				typ := info[3]
				timepara := info[5]
				if teachername == "" {
					continue
				}
				k = fmt.Sprintf("Номер пары: %v \nИмя преподавателя: %v \nПредмет: %v \nМесто проведения: %v \nТип пары : %v \nВремя начала : %v ", i, teachername, para, adress, typ, timepara)
				sendMessage(k)
			}

		} else if update.Message != nil && update.Message.Text == "Расписание на сегодня" && who == "teacher" {
			weekday := int(time.Now().Weekday()) - 1
			convert := map[int]string{
				0: "понедельник",
				1: "вторник",
				2: "среда",
				3: "четверг",
				4: "пятница",
			}
			result := database.GetTeachersDataBy(name, convert[weekday], weektype)
			lessons := result.Lessons
			if result.Teacher_name == "" {
				sendMessage("Нет пар в этот день")

			} else {

				for i, info := range lessons {
					var k string
					var groupname []string

					for _, inform := range info {
						if inform[1] != "" {
							groupname = append(groupname, inform[1])
						}
					}

					para := info[0][0]
					adress := info[0][2]
					typ := info[0][3]
					timepara := info[0][5]

					k = fmt.Sprintf("Номер пары: %v \nНомер группы: %v \nПредмет: %v \nМесто проведения: %v \nТип пары : %v \nВремя начала : %v ", i, groupname, para, adress, typ, timepara)
					sendMessage(k)

				}
			}

		} else if update.Message != nil && update.Message.Text == "Расписание на завтра" && who == "student" {
			weekday := int(time.Now().Weekday())
			convert := map[int]string{
				0: "понедельник",
				1: "вторник",
				2: "среда",
				3: "четверг",
				4: "пятница",
			}
			result := database.GetStudentsDataBy(group, convert[weekday], weektype)
			lessons := result.Lessons
			for i, info := range lessons {
				var k string
				para := info[0]
				teachername := info[1]
				adress := info[2]
				typ := info[3]
				timepara := info[5]
				if teachername == "" {
					continue
				}
				k = fmt.Sprintf("Номер пары: %v \nИмя преподавателя: %v \nПредмет: %v \nМесто проведения: %v \nТип пары : %v \nВремя начала : %v ", i, teachername, para, adress, typ, timepara)
				sendMessage(k)
			}
		} else if update.Message != nil && update.Message.Text == "Расписание на завтра" && who == "teacher" {
			weekday := int(time.Now().Weekday())
			convert := map[int]string{
				0: "понедельник",
				1: "вторник",
				2: "среда",
				3: "четверг",
				4: "пятница",
			}
			result := database.GetTeachersDataBy(name, convert[weekday], weektype)
			lessons := result.Lessons
			if result.Teacher_name == "" {
				sendMessage("Нет пар в этот день")

			} else {

				for i, info := range lessons {
					var k string
					var groupname []string

					for _, inform := range info {
						if inform[1] != "" {
							groupname = append(groupname, inform[1])
						}
					}

					para := info[0][0]
					adress := info[0][2]
					typ := info[0][3]
					timepara := info[0][5]

					k = fmt.Sprintf("Номер пары: %v \nНомер группы: %v \nПредмет: %v \nМесто проведения: %v \nТип пары : %v \nВремя начала : %v ", i, groupname, para, adress, typ, timepara)
					sendMessage(k)

				}
			}

		} else if update.Message != nil && update.Message.Text == "Расписание на определенный день недели" && who == "student" {
			sendStudentWeekdays(chatid)
		} else if update.Message != nil && update.Message.Text == "Расписание на определенный день недели" && who == "teacher" {
			sendStudentWeekdays(chatid)
		} else if update.Message != nil && contains(arr, update.Message.Text) && who == "student" {
			weekdaystr := strings.ToLower(update.Message.Text)
			fromWeekToNum := map[string]int{
				"понедельник": 0,
				"вторник":     1,
				"среда":       2,
				"четверг":     3,
				"пятница":     4,
			}

			weekday = fromWeekToNum[weekdaystr]
			convert := map[int]string{
				0: "понедельник",
				1: "вторник",
				2: "среда",
				3: "четверг",
				4: "пятница",
			}
			result := database.GetStudentsDataBy(group, convert[weekday], weektype)
			lessons := result.Lessons
			for i, info := range lessons {
				var k string
				para := info[0]
				teachername := info[1]
				adress := info[2]
				typ := info[3]
				timepara := info[5]
				if teachername == "" {
					continue
				}
				k = fmt.Sprintf("Номер пары: %v \nИмя преподавателя: %v \nПредмет: %v \nМесто проведения: %v \nТип пары : %v \nВремя начала : %v ", i, teachername, para, adress, typ, timepara)
				sendMessage(k)
			}
		} else if update.Message != nil && update.Message.Text == "Вернуться назад" {
			sendStudentDayS(chatid)

		} else if update.Message != nil && contains(arr, update.Message.Text) && who == "teacher" {
			weekdaystr := strings.ToLower(update.Message.Text)
			fromWeekToNum := map[string]int{
				"понедельник": 0,
				"вторник":     1,
				"среда":       2,
				"четверг":     3,
				"пятница":     4,
			}

			weekday = fromWeekToNum[weekdaystr]
			convert := map[int]string{
				0: "понедельник",
				1: "вторник",
				2: "среда",
				3: "четверг",
				4: "пятница",
			}
			result := database.GetTeachersDataBy(name, convert[weekday], weektype)
			lessons := result.Lessons

			if result.Teacher_name == "" {
				sendMessage("Нет пар в этот день")
			} else {

				for i, info := range lessons {
					var k string
					var groupname []string

					for _, inform := range info {
						if inform[1] != "" {
							groupname = append(groupname, inform[1])
						}
					}

					para := info[0][0]
					adress := info[0][2]
					typ := info[0][3]
					timepara := info[0][5]

					k = fmt.Sprintf("Номер пары: %v \nНомер группы: %v \nПредмет: %v \nМесто проведения: %v \nТип пары : %v \nВремя начала : %v ", i, groupname, para, adress, typ, timepara)
					sendMessage(k)
				}
			}

		} else if update.Message != nil && update.Message.Text == "Какая следующая пара?" && who == "student" {
			weekday := int(time.Now().Weekday()) - 1
			currentTime := time.Now()

			currentMinutes := currentTime.Hour()*60 + currentTime.Minute()

			switch {
			case currentMinutes >= 0*60 && currentMinutes < 8*60:
				para = 1
			case currentMinutes >= 8*60 && currentMinutes < 9*60+50:
				para = 2
			case currentMinutes >= 9*60+50 && currentMinutes < 11*60+30:
				para = 3
			case currentMinutes >= 11*60+30 && currentMinutes < 13*60+20:
				para = 4
			case currentMinutes >= 13*60+20 && currentMinutes < 15*60:
				para = 5
			case currentMinutes >= 15*60 && currentMinutes < 16*60+40:
				para = 6
			case currentMinutes >= 16*60+40 && currentMinutes < 18*60+20:
				para = 7
			case currentMinutes >= 18*60+20 && currentMinutes < 20*60:
				para = 8
			case currentMinutes >= 20*60:
				sendMessage("Пар сегодня больше нет ")
			default:
				para = 0
			}
			convert := map[int]string{
				0: "понедельник",
				1: "вторник",
				2: "среда",
				3: "четверг",
				4: "пятница",
			}
			result := database.GetStudentsDataBy(group, convert[weekday], weektype)
			lessons := result.Lessons

			nowlesson := lessons[para]
			lesson := nowlesson[0]
			teachername := nowlesson[1]
			adress := nowlesson[2]
			typ := nowlesson[3]
			timepara := nowlesson[5]
			if teachername == "" {
				sendMessage("Следующей пары нет")
			}
			k := fmt.Sprintf("Номер пары: %v \nИмя преподавателя: %v \nПредмет: %v \nМесто проведения: %v \nТип пары : %v \nВремя начала : %v ", para, teachername, lesson, adress, typ, timepara)
			sendMessage(k)

		} else if update.Message != nil && update.Message.Text == "Какая следующая пара?" && who == "teacher" {
			weekday := int(time.Now().Weekday()) - 1
			currentTime := time.Now()
			currentMinutes := currentTime.Hour()*60 + currentTime.Minute()

			switch {
			case currentMinutes >= 0*60 && currentMinutes < 8*60:
				para = 1
			case currentMinutes >= 8*60 && currentMinutes < 9*60+50:
				para = 2
			case currentMinutes >= 9*60+50 && currentMinutes < 11*60+30:
				para = 3
			case currentMinutes >= 11*60+30 && currentMinutes < 13*60+20:
				para = 4
			case currentMinutes >= 13*60+20 && currentMinutes < 15*60:
				para = 5
			case currentMinutes >= 15*60 && currentMinutes < 16*60+40:
				para = 6
			case currentMinutes >= 16*60+40 && currentMinutes < 18*60+20:
				para = 7
			case currentMinutes >= 18*60+20 && currentMinutes < 20*60:
				para = 8
			case currentMinutes >= 20*60:
				sendMessage("Пар сегодня больше нет ")
			default:
				para = 0
			}
			convert := map[int]string{
				0: "понедельник",
				1: "вторник",
				2: "среда",
				3: "четверг",
				4: "пятница",
			}

			result := database.GetTeachersDataBy(name, convert[weekday], weektype)
			lessons := result.Lessons
			if result.Teacher_name == "" {
				sendMessage("Cледующей пары нет")
				break
			}

			for i, info := range lessons {
				var k string
				var groupname []string
				var lesson int
				var address string
				var lessonType string
				var lessonTime string

				for _, inform := range info {
					if inform[1] != "" {

						groupname = append(groupname, inform[1])
					}
				}

				adress := info[0][2]
				typ := info[0][3]
				timepara := info[0][5]

				if i == para {
					lesson = para
					address = adress
					lessonType = typ
					lessonTime = timepara
					if len(groupname) > 0 {
						k = fmt.Sprintf("Номер пары: %v \nНомер группы: %v \nПредмет: %v \nМесто проведения: %v \nТип пары : %v \nВремя начала : %v ", i, groupname, lesson, address, lessonType, lessonTime)
						sendMessage(k)

					}
				}
			}

		}
	}
}
