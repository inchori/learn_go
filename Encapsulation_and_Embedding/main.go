package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/calendar"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func main() {
	date := calendar.Date{}
	err := date.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(8)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(21)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date.Year())

	event := calendar.Event{}
	err = event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(4)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(20)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
}
