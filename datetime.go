package doygo

import (
	// "fmt"
	"strconv"
)

type Datetime struct {
	YEAR        int64
	MONTH       int64
	DAY         int64
	DOY         int64
	HOUR        int64
	MINUTE      int64
	SECOND      float64
	GPSWEEK     int64
	GPSWEEKDAY  int64
	GPSSECOND   int64
	DECIMALYEAR float64
	JD          float64
	MJD         float64
}

func Doy(left ...interface{}) (datetime Datetime) {

	// if left >= 2, then "year month day"
	// if left == 1, then "year doy"
	// if left == 0, then "today"

	if len(left) == 3 {

		// year parsing
		switch left[0].(type) {
		case int64:
			datetime.YEAR = int64(left[0].(int64))
		case int:
			datetime.YEAR = int64(left[0].(int))
		case string:
			year, _ := strconv.Atoi(left[0].(string))
			datetime.YEAR = int64(year)
		}
		// month parsing
		switch left[1].(type) {
		case int64:
			datetime.MONTH = int64(left[1].(int64))
		case int:
			datetime.MONTH = int64(left[1].(int))
		case string:
			month, _ := strconv.Atoi(left[1].(string))
			datetime.MONTH = int64(month)
		}
		// day parsing
		switch left[2].(type) {
		case int64:
			datetime.DAY = int64(left[2].(int64))
		case int:
			datetime.DAY = int64(left[2].(int))
		case string:
			day, _ := strconv.Atoi(left[2].(string))
			datetime.DAY = int64(day)
		}

		// hour, min, and second
		datetime.HOUR = 0
		datetime.MINUTE = 0
		datetime.SECOND = 0.0

		// get doy
		_, datetime.DOY = Cal2YearDoy(
			datetime.YEAR, datetime.MONTH, datetime.DAY)
		// get decimal year
		if LeapYear(datetime.YEAR) {
			datetime.DECIMALYEAR = float64(datetime.DOY-1) / 366.0
		} else {
			datetime.DECIMALYEAR = float64(datetime.DOY-1) / 365.0
		}
		// get jd and mjd
		datetime.JD = Cal2Jd(
			datetime.YEAR,
			datetime.MONTH,
			datetime.DAY,
			datetime.HOUR,
			datetime.MINUTE,
			datetime.SECOND)
		datetime.MJD = datetime.JD - 2400000.5
		// get gps time
		datetime.GPSWEEK, datetime.GPSSECOND = Cal2GPSweeksec(
			datetime.YEAR,
			datetime.MONTH,
			datetime.DAY,
			datetime.HOUR,
			datetime.MINUTE,
			datetime.SECOND)
		datetime.GPSWEEKDAY = Mjd2Weekday(datetime.MJD)

	} else if len(left) == 6 {

		// year parsing
		switch left[0].(type) {
		case int64:
			datetime.YEAR = int64(left[0].(int64))
		case int:
			datetime.YEAR = int64(left[0].(int))
		case string:
			year, _ := strconv.Atoi(left[0].(string))
			datetime.YEAR = int64(year)
		}
		// month parsing
		switch left[1].(type) {
		case int64:
			datetime.MONTH = int64(left[1].(int64))
		case int:
			datetime.MONTH = int64(left[1].(int))
		case string:
			month, _ := strconv.Atoi(left[1].(string))
			datetime.MONTH = int64(month)
		}
		// day parsing
		switch left[2].(type) {
		case int64:
			datetime.DAY = int64(left[2].(int64))
		case int:
			datetime.DAY = int64(left[2].(int))
		case string:
			day, _ := strconv.Atoi(left[2].(string))
			datetime.DAY = int64(day)
		}
		// hour parsing
		switch left[3].(type) {
		case int64:
			datetime.HOUR = int64(left[3].(int64))
		case int:
			datetime.HOUR = int64(left[3].(int))
		case string:
			hour, _ := strconv.Atoi(left[3].(string))
			datetime.HOUR = int64(hour)
		}
		// min parsing
		switch left[4].(type) {
		case int64:
			datetime.MINUTE = int64(left[4].(int64))
		case int:
			datetime.MINUTE = int64(left[4].(int))
		case string:
			min, _ := strconv.Atoi(left[4].(string))
			datetime.MINUTE = int64(min)
		}
		// second parsing
		switch left[5].(type) {
		case int64:
			datetime.SECOND = float64(left[5].(int64))
		case int:
			datetime.SECOND = float64(left[5].(int))
		case string:
			datetime.SECOND, _ = strconv.ParseFloat(left[5].(string), 64)
		}

		// get doy
		_, datetime.DOY = Cal2YearDoy(
			datetime.YEAR, datetime.MONTH, datetime.DAY)
		// get decimal year
		if LeapYear(datetime.YEAR) {
			datetime.DECIMALYEAR = float64(datetime.DOY-1) / 366.0
		} else {
			datetime.DECIMALYEAR = float64(datetime.DOY-1) / 365.0
		}
		// get jd and mjd
		datetime.JD = Cal2Jd(
			datetime.YEAR,
			datetime.MONTH,
			datetime.DAY,
			datetime.HOUR,
			datetime.MINUTE,
			datetime.SECOND)
		datetime.MJD = datetime.JD - 2400000.5
		// get gps time
		datetime.GPSWEEK, datetime.GPSSECOND = Cal2GPSweeksec(
			datetime.YEAR,
			datetime.MONTH,
			datetime.DAY,
			datetime.HOUR,
			datetime.MINUTE,
			datetime.SECOND)
		datetime.GPSWEEKDAY = Mjd2Weekday(datetime.MJD)

	} else if len(left) == 4 {

		if left[0] == "w" && left[2] == "d" {

			// week parsing
			switch left[1].(type) {
			case int64:
				datetime.GPSWEEK = int64(left[1].(int64))
			case int:
				datetime.GPSWEEK = int64(left[1].(int))
			case string:
				week, _ := strconv.Atoi(left[1].(string))
				datetime.GPSWEEK = int64(week)
			}
			// weekday parsing
			switch left[3].(type) {
			case int64:
				datetime.GPSWEEKDAY = int64(left[3].(int64))
			case int:
				datetime.GPSWEEKDAY = int64(left[3].(int))
			case string:
				wd, _ := strconv.Atoi(left[3].(string))
				datetime.GPSWEEKDAY = int64(wd)
			}
			// weekday to cal
			datetime.YEAR,
				datetime.MONTH,
				datetime.DAY,
				datetime.HOUR,
				datetime.MINUTE,
				datetime.SECOND =
				GPSweekday2Cal(datetime.GPSWEEK, datetime.GPSWEEKDAY)

			// get doy
			_, datetime.DOY = Cal2YearDoy(
				datetime.YEAR, datetime.MONTH, datetime.DAY)
			// get decimal year
			if LeapYear(datetime.YEAR) {
				datetime.DECIMALYEAR = float64(datetime.DOY-1) / 366.0
			} else {
				datetime.DECIMALYEAR = float64(datetime.DOY-1) / 365.0
			}
			// get jd and mjd
			datetime.JD = Cal2Jd(
				datetime.YEAR,
				datetime.MONTH,
				datetime.DAY,
				datetime.HOUR,
				datetime.MINUTE,
				datetime.SECOND)
			datetime.MJD = datetime.JD - 2400000.5
		}

	} else if len(left) == 2 {

		// year parsing
		switch left[0].(type) {
		case int64:
			datetime.YEAR = int64(left[0].(int64))
		case int:
			datetime.YEAR = int64(left[0].(int))
		case string:
			year, _ := strconv.Atoi(left[0].(string))
			datetime.YEAR = int64(year)
		}
		// doy parsing
		switch left[1].(type) {
		case int64:
			datetime.DOY = int64(left[1].(int64))
		case int:
			datetime.DOY = int64(left[1].(int))
		case string:
			doy, _ := strconv.Atoi(left[1].(string))
			datetime.DOY = int64(doy)
		}

		// hour, min, and second
		datetime.HOUR = 0
		datetime.MINUTE = 0
		datetime.SECOND = 0.0

		// get month and day
		_, datetime.MONTH, datetime.DAY = YearDoy2Cal(
			datetime.YEAR, datetime.DOY)
		// get decimal year
		if LeapYear(datetime.YEAR) {
			datetime.DECIMALYEAR = float64(datetime.DOY-1) / 366.0
		} else {
			datetime.DECIMALYEAR = float64(datetime.DOY-1) / 365.0
		}
		// get jd and mjd
		datetime.JD = Cal2Jd(
			datetime.YEAR,
			datetime.MONTH,
			datetime.DAY,
			datetime.HOUR,
			datetime.MINUTE,
			datetime.SECOND)
		datetime.MJD = datetime.JD - 2400000.5
		// get gps time
		datetime.GPSWEEK, datetime.GPSSECOND = Cal2GPSweeksec(
			datetime.YEAR,
			datetime.MONTH,
			datetime.DAY,
			datetime.HOUR,
			datetime.MINUTE,
			datetime.SECOND)
		datetime.GPSWEEKDAY = Mjd2Weekday(datetime.MJD)

	} else if len(left) == 5 {

		// year parsing
		switch left[0].(type) {
		case int64:
			datetime.YEAR = int64(left[0].(int64))
		case int:
			datetime.YEAR = int64(left[0].(int))
		case string:
			year, _ := strconv.Atoi(left[0].(string))
			datetime.YEAR = int64(year)
		}
		// doy parsing
		switch left[1].(type) {
		case int64:
			datetime.DOY = int64(left[1].(int64))
		case int:
			datetime.DOY = int64(left[1].(int))
		case string:
			doy, _ := strconv.Atoi(left[1].(string))
			datetime.DOY = int64(doy)
		}
		// hour parsing
		switch left[2].(type) {
		case int64:
			datetime.HOUR = int64(left[2].(int64))
		case int:
			datetime.HOUR = int64(left[2].(int))
		case string:
			hour, _ := strconv.Atoi(left[2].(string))
			datetime.HOUR = int64(hour)
		}
		// min parsing
		switch left[3].(type) {
		case int64:
			datetime.MINUTE = int64(left[3].(int64))
		case int:
			datetime.MINUTE = int64(left[3].(int))
		case string:
			min, _ := strconv.Atoi(left[3].(string))
			datetime.MINUTE = int64(min)
		}
		// second parsing
		switch left[4].(type) {
		case int64:
			datetime.SECOND = float64(left[4].(int64))
		case int:
			datetime.SECOND = float64(left[4].(int))
		case string:
			datetime.SECOND, _ = strconv.ParseFloat(left[4].(string), 64)
		}

		// get month and day
		_, datetime.MONTH, datetime.DAY = YearDoy2Cal(
			datetime.YEAR, datetime.DOY)
		// get decimal year
		if LeapYear(datetime.YEAR) {
			datetime.DECIMALYEAR = float64(datetime.DOY-1) / 366.0
		} else {
			datetime.DECIMALYEAR = float64(datetime.DOY-1) / 365.0
		}
		// get jd and mjd
		datetime.JD = Cal2Jd(
			datetime.YEAR,
			datetime.MONTH,
			datetime.DAY,
			datetime.HOUR,
			datetime.MINUTE,
			datetime.SECOND)
		datetime.MJD = datetime.JD - 2400000.5
		// get gps time
		datetime.GPSWEEK, datetime.GPSSECOND = Cal2GPSweeksec(
			datetime.YEAR,
			datetime.MONTH,
			datetime.DAY,
			datetime.HOUR,
			datetime.MINUTE,
			datetime.SECOND)
		datetime.GPSWEEKDAY = Mjd2Weekday(datetime.MJD)

	}

	return
}
