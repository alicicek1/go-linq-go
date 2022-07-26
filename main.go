package main

import (
	"fmt"
	"github.com/ahmetb/go-linq/v3"
	"time"
)

type Student struct {
	Name          string
	Surname       string
	Age           int32
	DateOfBirth   time.Time
	IsRightHanded bool
}

func main() {
	fmt.Println("------------")
	studentList := CreateStudentList()
	fmt.Println("Whole List")
	PrintList(studentList)
	fmt.Println("------------")

	var rightHandedList []Student
	linq.From(studentList).Where(func(i interface{}) bool {
		return i.(Student).IsRightHanded
	}).Select(func(c interface{}) interface{} {
		return c.(Student)
	}).ToSlice(&rightHandedList)
	fmt.Println("Right handed list")
	PrintList(rightHandedList)
	fmt.Println("------------")

	var leftHandedSurname []string
	linq.From(studentList).Where(func(i interface{}) bool {
		return !i.(Student).IsRightHanded
	}).Select(func(i interface{}) interface{} {
		return i.(Student).Surname
	}).ToSlice(&leftHandedSurname)
	fmt.Println("Left handed surname list")
	for _, val := range leftHandedSurname {
		fmt.Println(val)
	}
	fmt.Println("------------")

	var aboveAndEqual45Age []Student
	linq.From(studentList).Where(func(i interface{}) bool {
		return i.(Student).Age >= 45
	}).Select(func(i interface{}) interface{} {
		return i.(Student)
	}).ToSlice(&aboveAndEqual45Age)
	fmt.Println("aboveAndEqual45Age list")
	PrintList(aboveAndEqual45Age)
	fmt.Println("------------")

	var orderByAge []Student
	linq.From(studentList).OrderBy(func(i interface{}) interface{} {
		return i.(Student).Age
	}).ToSlice(&orderByAge)
	fmt.Println("orderByAge")
	PrintList(orderByAge)
	fmt.Println("------------")

	var orderByDescAgeTake2 []Student
	linq.From(studentList).OrderByDescending(func(i interface{}) interface{} {
		return i.(Student).Age
	}).Take(2).ToSlice(&orderByDescAgeTake2)
	fmt.Println("orderByDescAgeTake2")
	PrintList(orderByDescAgeTake2)
	fmt.Println("------------")

	fmt.Println(linq.From(studentList).First())

	fmt.Println("------------")
	fmt.Println(linq.From(studentList).Last())

	fmt.Println("------------")
	fmt.Println(linq.From(studentList).LastWith(func(i interface{}) bool {
		return i.(Student).Age > 35
	}))

	fmt.Println("------------")
	fmt.Println(linq.From(studentList).Where(func(i interface{}) bool {
		return i.(Student).Name == "Albert"
	}).FirstWith(func(i interface{}) bool {
		return i.(Student).Age > 35
	}))

	fmt.Println("------------")
	fmt.Println(linq.From(studentList).Where(func(i interface{}) bool {
		return i.(Student).Name == "Albert"
	}).Any())

}

func PrintList(list []Student) {
	for _, student := range list {
		fmt.Println(student)
	}
}

func CreateStudentList() []Student {
	loc, err := time.LoadLocation("")
	if err != nil {
		fmt.Println(err)
	}
	return []Student{{
		Name:          "Isaac",
		Surname:       "Newton",
		Age:           25,
		DateOfBirth:   time.Date(1964, 1, 4, 0, 0, 0, 0, loc),
		IsRightHanded: true,
	}, {
		Name:          "Nicola",
		Surname:       "Tesla",
		Age:           35,
		DateOfBirth:   time.Date(1986, 7, 10, 0, 0, 0, 0, loc),
		IsRightHanded: true,
	}, {
		Name:          "Albert",
		Surname:       "Einstein",
		Age:           45,
		DateOfBirth:   time.Date(1879, 4, 14, 0, 0, 0, 0, loc),
		IsRightHanded: true,
	}, {
		Name:          "Maira",
		Surname:       "Curie",
		Age:           55,
		DateOfBirth:   time.Date(1867, 11, 7, 0, 0, 0, 0, loc),
		IsRightHanded: false,
	}, {
		Name:          "Leonardo",
		Surname:       "Da Vinci",
		Age:           55,
		DateOfBirth:   time.Date(1877, 1, 7, 0, 0, 0, 0, loc),
		IsRightHanded: false,
	}}
}
