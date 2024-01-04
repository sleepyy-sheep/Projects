package database

import (
  "encoding/json"
  "io/ioutil"
)

type Student struct {
  Lessons       map[int][]string
  Group_number  string
  EvenOrOddWeek string
  DayOfTheWeek  string
}

func readStudentInfo() []Student {

  var Allinfo []Student //база данных,json, хранит формат списка структур Student

  file, err := ioutil.ReadFile("database/students.json") //считываем массив байтов из файла
  if err != nil {
    panic(err)
  }
  err = json.Unmarshal(file, &Allinfo) //превращаем массив байтов в срез структур Student
  if err != nil {
    panic(err)
  }
  return Allinfo
}

func GetStudentsDataBy(group_number string, weekday string, EvenOrOddWeek string) Student {
  schedule := readStudentInfo()
  for _, info := range schedule { //проходим по срезу структуры
    //есл  номер групп, день недели и тип недели совпадают, то возвращаем
    if info.Group_number == group_number && info.DayOfTheWeek == weekday && info.EvenOrOddWeek == EvenOrOddWeek {
      return info
    }
  } //если не совпадает ничего, то пустую структуру возвращаем
  return Student{}
}

func GetStudentNext(group_number, weekday, EvenOrOddWeek string, pair int) []string {
  schedule := readStudentInfo()
  for _, info := range schedule {
    if info.Group_number == group_number && info.DayOfTheWeek == weekday && info.EvenOrOddWeek == EvenOrOddWeek {
      return info.Lessons[pair]
    }
  }
  return nil
}

type Teacher struct {
  Lessons       map[int][][]string
  Teacher_name  string
  EvenOrOddWeek string
  DayOfTheWeek  string
}

func readTeacherInfo() []Teacher {

  var Allinfo []Teacher

  file, err := ioutil.ReadFile("database/teachers.json")
  if err != nil {
    panic(err)
  }
  err = json.Unmarshal(file, &Allinfo)
  if err != nil {
    panic(err)
  }
  return Allinfo
}

func GetTeachersDataBy(teacher_name string, weekday string, EvenOrOddWeek string) Teacher {
  schedule := readTeacherInfo()
  for _, info := range schedule {
    if info.Teacher_name == teacher_name && info.DayOfTheWeek == weekday && info.EvenOrOddWeek == EvenOrOddWeek {
      return info
    }
  }
  return Teacher{}
}

func GetTeachersNext(name, weekday, EvenOrOddWeek string, pair int) [][]string {
  schedule := readTeacherInfo()
  for _, info := range schedule {
    if info.Teacher_name == name && info.DayOfTheWeek == weekday && info.EvenOrOddWeek == EvenOrOddWeek {
      return info.Lessons[pair]
    }
  }
  return nil
}
