package booking

import (
    "time"
    "fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
    parsedDate, _ := time.Parse(layout, date)
    return parsedDate
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
	scheduledTime, _ := time.Parse(layout, date)
    today := time.Now()
    return scheduledTime.Before(today)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
    schedule, _ := time.Parse(layout, date)
    if (schedule.Hour() >= 12 && schedule.Hour() < 18){
        return true
    }
    return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    layout := "1/2/2006 15:04:05"
    schedule, _ := time.Parse(layout, date)
    return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", schedule.Weekday(), schedule.Month(), schedule.Day(), schedule.Year(), schedule.Hour(), schedule.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	layout := "2006-01-02"
    thisYear := time.Now().Year()
    date := fmt.Sprintf("%d-09-15", thisYear)
    ann,_ := time.Parse(layout, date)
    return ann
}
