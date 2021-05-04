package utils

import (
	"fmt"
	"time"
)

type BookingPeriod struct {
	From *BookingDate `json:"from"`
	To   *BookingDate `json:"to"`
}

type BookingDate struct {
	Year int `json:"year"`
	Day  int `json:"day"`
}

func FromDateString(dateAsString string) (*BookingDate, error) {

	date, errParsing := time.Parse(resharmonicsDateLayout, dateAsString)
	if errParsing != nil {
		return nil, errParsing
	}
	return &BookingDate{
		Year: date.Year(),
		Day:  date.YearDay(),
	}, nil
}

const resharmonicsDateLayout = `2006-01-02`

func (period *BookingPeriod) Validate() error {
	if period == nil {
		return fmt.Errorf("No period specified")
	}
	if period.From == nil {
		return fmt.Errorf("No from specified")
	}
	if period.To == nil {
		return fmt.Errorf("No to specified")
	}

	if period.From.Year > period.To.Year {
		return fmt.Errorf("Year of from can't be after year of To")
	}

	if period.From.Year == period.To.Year {
		if period.From.Day >= period.To.Year {
			return fmt.Errorf("From and to on the same year but day of from has to be before day of to")
		}
	}
	return nil
}

func (bookingDate *BookingDate) ToDate() time.Time {
	asDate := time.Date(bookingDate.Year, time.January, 0, 0, 0, 0, 0, time.UTC)
	daysToAdd := time.Duration(time.Hour.Nanoseconds() * (int64)(24*bookingDate.Day))
	asDate = asDate.Add(daysToAdd)
	return asDate
}

func (bookingDate *BookingDate) ToResharmonicsString() string {
	asDate := bookingDate.ToDate()
	return asDate.Format(resharmonicsDateLayout)
}
