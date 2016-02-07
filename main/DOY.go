package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"leung.com/doygo"
)

func main() {

	number_of_args := len(os.Args[1:])

	dt := new(doygo.Datetime)

	if number_of_args == 3 {
		// year nonth day
		dt.YEAR, _ = strconv.ParseInt(os.Args[1], 10, 64)
		dt.MONTH, _ = strconv.ParseInt(os.Args[2], 10, 64)
		dt.DAY, _ = strconv.ParseInt(os.Args[3], 10, 64)
		dt.CheckDate()
		_, dt.DOY = doygo.Cal2YearDoy(dt.YEAR, dt.MONTH, dt.DAY)
		dt.JD = doygo.Cal2Jd(dt.YEAR, dt.MONTH, dt.DAY, 0, 0, 0.0)
		dt.MJD = doygo.Jd2Mjd(dt.JD)
		dt.GPSWEEK, dt.GPSWEEKDAY, dt.GPSSECOND =
			doygo.Mjd2GPSweekdaysec(dt.MJD)
		if doygo.LeapYear(dt.YEAR) {
			dt.DECIMALYEAR = float64(dt.YEAR) + float64(dt.DOY-1)/366.0
		} else {
			dt.DECIMALYEAR = float64(dt.YEAR) + float64(dt.DOY-1)/365.0
		}

	} else if number_of_args == 2 {

		if strings.HasSuffix(os.Args[1], "w") {
			//gpsweek gpsweekday
			dt.GPSWEEK, _ = strconv.ParseInt(
				os.Args[1][0:len(os.Args[1])], 10, 64)
			dt.GPSWEEKDAY, _ = strconv.ParseInt(os.Args[2], 10, 64)
			dt.YEAR, dt.MONTH, dt.DAY, dt.HOUR, dt.MINUTE, dt.SECOND =
				doygo.GPSweekday2Cal(dt.GPSWEEK, dt.GPSWEEKDAY)
			dt.JD = doygo.Cal2Jd(dt.YEAR, dt.MONTH, dt.DAY, 0, 0, 0.0)
			dt.MJD = doygo.Jd2Mjd(dt.JD)
			_, dt.DOY = doygo.Cal2YearDoy(dt.YEAR, dt.MONTH, dt.DAY)
			if doygo.LeapYear(dt.YEAR) {
				dt.DECIMALYEAR = float64(dt.YEAR) + float64(dt.DOY-1)/366.0
			} else {
				dt.DECIMALYEAR = float64(dt.YEAR) + float64(dt.DOY-1)/365.0
			}

		} else {
			// year doy
			dt.YEAR, _ = strconv.ParseInt(os.Args[1], 10, 64)
			dt.DOY, _ = strconv.ParseInt(os.Args[2], 10, 64)
			dt.YEAR, dt.MONTH, dt.DAY = doygo.YearDoy2Cal(dt.YEAR, dt.DOY)
			dt.JD = doygo.Cal2Jd(dt.YEAR, dt.MONTH, dt.DAY, 0, 0, 0.0)
			dt.MJD = doygo.Jd2Mjd(dt.JD)
			dt.GPSWEEK, dt.GPSWEEKDAY, dt.GPSSECOND =
				doygo.Mjd2GPSweekdaysec(dt.MJD)
			if doygo.LeapYear(dt.YEAR) {
				dt.DECIMALYEAR = float64(dt.YEAR) + float64(dt.DOY-1)/366.0
			} else {
				dt.DECIMALYEAR = float64(dt.YEAR) + float64(dt.DOY-1)/365.0
			}

		}

	} else {

	}

	format := "Date %d/%d/%d %d:%d hrs, DOY %03d JD %.4f MJD %.4f\n" +
		"GPS week %d Day of week %d, GPS Seconds %d Day of week %s\n" +
		"Decimal Year %.9f\n"
	fmt.Printf(format,
		dt.YEAR,
		dt.MONTH,
		dt.DAY,
		dt.HOUR,
		dt.MINUTE,
		dt.DOY,
		dt.JD,
		dt.MJD,
		dt.GPSWEEK,
		dt.GPSWEEKDAY,
		dt.GPSSECOND,
		doygo.DayofWeek(dt.GPSWEEKDAY),
		dt.DECIMALYEAR)
}
