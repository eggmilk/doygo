package doygo

import (
	"fmt"
	"testing"
)

func Test_Cal2YearDoy(t *testing.T) {
	fmt.Println("\nCal2YearDoy(year, month, day)")
	fmt.Println(Cal2YearDoy(2001, 10, 23))
	fmt.Println(Cal2YearDoy(2009, 3, 9))
	fmt.Println(Cal2YearDoy(2015, 12, 21))
}

func Test_YearDoy2Cal(t *testing.T) {
	fmt.Println("\nYearDoy2Cal(year, doy)")
	fmt.Println(YearDoy2Cal(2016, 336))
	fmt.Println(YearDoy2Cal(2015, 330))
	fmt.Println(YearDoy2Cal(2001, 297))
	fmt.Println(YearDoy2Cal(2017, 38))
}

func Test_Jd2Cal(t *testing.T) {
	fmt.Println("\nJd2Cal(julia)")
	fmt.Println(Jd2Cal(1721423.5))
	fmt.Println(Jd2Cal(2.4573885e+06))
}

func Test_Mjd2Cal(t *testing.T) {
	fmt.Println("\nMjd2Cal(mjd)")
	fmt.Println(Mjd2Cal(0))
	fmt.Println(Mjd2Cal(44244.0)) // GPS start time
	fmt.Println(Mjd2Cal(44244.933))
}

func Test_Cal2Mjd(t *testing.T) {
	fmt.Println("\nCal2Mjd(year, month, day, hour, min, sec)")
	fmt.Println(Cal2Mjd(2009, 3, 9, 0, 0, 0.0))
	fmt.Println(Cal2Mjd(2008, 7, 27, 0, 0, 0.0))
	fmt.Println(Cal2Mjd(2015, 12, 17, 0, 0, 0.0))
	fmt.Println(Cal2Mjd(1982, 8, 26, 0, 0, 0.0))
}

func Test_Cal2Jd(t *testing.T) {
	fmt.Println("\nCal2Jd(year, month, day, hour, min, sec)")
	fmt.Println(Cal2Jd(1, 1, 1, 0, 0, 0.0))
	fmt.Println(Cal2Jd(2009, 3, 9, 0, 0, 0.0))
	fmt.Println(Cal2Jd(2015, 12, 17, 0, 0, 0.0))
	fmt.Println(Cal2Jd(2015, 12, 32, 0, 0, 0.0))
}

func Test_Mjd2Weekday(t *testing.T) {
	fmt.Println("\nMjd2Weekday(mjd)")
	fmt.Println(Mjd2Weekday(54899.0)) // 2009/3/9
	fmt.Println(Mjd2Weekday(54674.0)) // 2008/7/27
	fmt.Println(Mjd2Weekday(57373.0)) // 2015/12/17
	fmt.Println(Mjd2Weekday(45207.0))
	fmt.Println(Mjd2Weekday(44244.0)) // GPS start
}

func Test_Cal2Weekday(t *testing.T) {
	fmt.Println("\nCal2Weekday(year,month,day)")
	fmt.Println(Cal2Weekday(2009, 3, 9))
	fmt.Println(Cal2Weekday(2008, 7, 27))
	fmt.Println(Cal2Weekday(2015, 12, 17))
}

func Test_Cal2GPSweekday(t *testing.T) {
	fmt.Println("\nCal2GPSweekday(year, month, day, hour, min, sec)")
	fmt.Println(Cal2GPSweekday(1980, 1, 6, 0, 0, 0.0))
	fmt.Println(Cal2GPSweekday(2015, 12, 17, 0, 0, 0.0))
}

func Test_GPSweekday2Cal(t *testing.T) {
	fmt.Println("\nGPSweekday2Cal(gpsweek, gpsweekday)")
	fmt.Println(GPSweekday2Cal(1875, 4))
}

func Test_Cal2GPSweeksec(t *testing.T) {
	fmt.Println("\nCal2GPSweeksec(year, month, day, hour, min, sec)")
	fmt.Println(Cal2GPSweeksec(1980, 1, 6, 0, 0, 0.0))
	fmt.Println(Cal2GPSweeksec(2015, 12, 17, 0, 0, 0.0))
}

func Test_GPSweeksec2Cal(t *testing.T) {
	fmt.Println("\nGPSweeksec2Cal(gpsweek, gpsweeksec)")
	fmt.Println(GPSweeksec2Cal(1875, 4*86400))
}

func Test_DateGPSweeksec2Cal(t *testing.T) {
	fmt.Println("\nDateGPSweeksec2Cal(iyear, imonth, iday, igpsweeksec)")
	fmt.Println(DateGPSweeksec2Cal(2015, 12, 17, 3*86400+12*3600+55*60+12))
}

func Test_CheckMonth(t *testing.T) {
	fmt.Println("\n(dt *Datetime) CheckMonth()")
	dt := new(Datetime)
	dt.YEAR = 2017
	dt.MONTH = -3
	dt.DAY = 1
	dt.HOUR = 1000
	dt.CheckDate()
	fmt.Println(dt)
}
