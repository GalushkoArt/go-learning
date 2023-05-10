package basics

import (
	"errors"
	"fmt"
	"log"
)

func Enums(month *Month) {
	season, err := getSeasonFromMonth(month)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println((*season).String())
}

func getSeasonFromMonth(month *Month) (*Season, error) {
	if month == nil {
		return nil, errors.New("nil was passed instead of month pointer")
	}
	var season Season
	switch *month {
	case December, January, February:
		season = Winter
	case March, April, May:
		season = Spring
	case June, July, August:
		season = Summer
	case September, October, November:
		season = Autumn
	default:
		return nil, errors.New(fmt.Sprintf("%v is a wrong month", month))
	}
	return &season, nil
}

type Season int64

const (
	Spring Season = iota
	Summer
	Autumn
	Winter
)

func (s Season) String() string {
	return [...]string{
		"Spring",
		"Summer",
		"Autumn",
		"Winter",
	}[s]
}

type Month int64

const (
	January Month = iota + 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func (m Month) String() string {
	return [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}[m-1]
}
