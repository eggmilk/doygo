package doygo

import (
	"math"

	"leung.com/utilities/mathematics"
)

func LeapYear(year int64) bool {

	var bool_leap bool = false

	leap_year := math.Mod(float64(year), 4.0)

	// Judge whether year is leap year by condition 1 and 2
	if leap_year > 0 {
		bool_leap = false
	} else {
		if math.Mod(float64(year), 100.0) == 0 &&
			math.Mod(float64(year), 400.0) != 0 {
			bool_leap = false
		} else {
			bool_leap = true
		}
	}

	return bool_leap
}

func (dt *Datetime) CheckDate() {
	jd := Cal2Jd(dt.YEAR, dt.MONTH, dt.DAY,
		dt.HOUR, dt.MINUTE, dt.SECOND)
	dt.YEAR, dt.MONTH, dt.DAY,
		dt.HOUR, dt.MINUTE, dt.SECOND = Jd2Cal(jd)
}

func Cal2YearDoy(year, month, day int64) (oyear, odoy int64) {

	bool_leap := LeapYear(year)

	switch month {
	case 1:
		odoy = day
	case 2:
		odoy = day + 31
	case 3:
		odoy = day + 59
	case 4:
		odoy = day + 90
	case 5:
		odoy = day + 120
	case 6:
		odoy = day + 151
	case 7:
		odoy = day + 181
	case 8:
		odoy = day + 212
	case 9:
		odoy = day + 243
	case 10:
		odoy = day + 273
	case 11:
		odoy = day + 304
	case 12:
		odoy = day + 334
	}

	if bool_leap == true {
		// if month <= 2 {
		// 	odoy = odoy
		// }
		if month >= 3 {
			odoy = odoy + 1
		}
	}

	oyear = year
	return
}

func YearDoy2Cal(year, doy int64) (oyear, omonth, oday int64) {

	bool_leap := LeapYear(year)

	if bool_leap == false {

		_, max_doy := Cal2YearDoy(year, 12, 31)

		get_month := func(year, doy int64) (oyear, omonth, oday int64) {
			if doy > 0 && doy <= max_doy {
				oyear = year
				for i, v := range []int64{
					334, 304, 273, 243, 212, 181, 151, 120, 90, 59, 31} {
					if doy > v {
						omonth = 12 - int64(i)
						oday = doy - v
						break
					} else {
						oyear = year
						omonth = 1
						oday = doy
					}
				}
			}
			return
		}

		if doy > 0 && doy <= max_doy {

			oyear, omonth, oday = get_month(year, doy)

		} else if doy <= 0 {

			for {
				year -= 1
				_, max_doy = Cal2YearDoy(year, 12, 31)
				doy += max_doy
				if doy > 0 && doy <= max_doy {
					break
				}
			}
			oyear, omonth, oday = get_month(year, doy)

		} else if doy > max_doy {

			for {
				doy -= max_doy
				year += 1
				_, max_doy = Cal2YearDoy(year, 12, 31)
				if doy > 0 && doy <= max_doy {
					break
				}
			}
			oyear, omonth, oday = get_month(year, doy)

		}

		return

	} else {

		_, max_doy := Cal2YearDoy(year, 12, 31)

		get_month := func(year, doy int64) (oyear, omonth, oday int64) {
			if doy > 0 && doy <= max_doy {
				for i, v := range []int64{
					335, 305, 274, 244, 213, 182, 152, 121, 91, 60, 31} {
					if doy > v {
						oyear = year
						omonth = 12 - int64(i)
						oday = doy - v
					} else {
						oyear = year
						omonth = 1
						oday = doy
					}
				}
			}
			return
		}

		if doy > 0 && doy <= max_doy {

			oyear, omonth, oday = get_month(year, doy)

		} else if doy <= 0 {

			for {
				year -= 1
				_, max_doy = Cal2YearDoy(year, 12, 31)
				doy += max_doy
				if doy > 0 && doy <= max_doy {
					break
				}
			}
			oyear, omonth, oday = get_month(year, doy)

		} else if doy > max_doy {

			for {
				doy -= max_doy
				year += 1
				_, max_doy = Cal2YearDoy(year, 12, 31)
				if doy > 0 && doy <= max_doy {
					break
				}
			}
			oyear, omonth, oday = get_month(year, doy)

		}

		return
	}
	return
}

func Jd2Cal(julia float64) (
	oyear, omonth, oday, ohour, omin int64, sec float64) {
	// 公元 1582 年 10 月 4 日 24:00 点之前使用儒略历
	// 公元 1582 年 10 月 15 日 00:00 点之后使用公历

	var bc int64
	var j0, dd float64
	var n1, n2, n3 float64
	var year, month, day, hour, min float64
	var year0 float64

	if julia < 1721423.5 {
		bc = 1
	} else {
		bc = 0
	}

	// start from Julian March 1, 4801 B.C.
	if julia < 2299160.5 {
		// before 1582.10.4. 24:00 is Julian calender
		j0 = math.Floor(julia + 0.5)
		dd = julia + 0.5 - j0
	} else {
		// after 1582.10.15. 00:00 is Gregorian calender
		// number of certury years that are not leap year
		n1 = math.Floor((julia-2342031.5)/36524.25/4) + 1 // 1700.3.1.0
		n2 = math.Floor((julia-2378555.5)/36524.25/4) + 1 // 1800.3.1.0
		n3 = math.Floor((julia-2415079.5)/36524.25/4) + 1 // 1900.3.1.0
		j0 = n1 + n2 + n3 + 10 + julia
		dd = j0 + 0.5 - math.Floor(j0+0.5)
		j0 = math.Floor(j0 + 0.5)
	}

	j0 += 32083
	year0 = math.Ceil(j0/365.25) - 1.0
	year = year0 - 4800.0
	day = j0 - math.Floor(year0*365.25)
	month = math.Floor((day-0.6)/30.6) + 3.0
	day = day - mathematics.Round((month-3.0)*30.6, 0)

	if month > 12 {
		month -= 12
		year += 1
	}

	year -= float64(bc)

	sec = mathematics.Round(dd*86400.0, 0)

	hour = math.Floor(sec / 3600.0)
	sec -= hour * 3600.0
	min = math.Floor(sec / 60.0)
	sec = sec - min*60.0

	oyear = int64(year)
	omonth = int64(month)
	oday = int64(day)
	ohour = int64(hour)
	omin = int64(min)

	return
}

func Cal2Jd(year, month, day, hour, min int64, sec float64) (jd float64) {
	fyear := float64(year)
	fmonth := float64(month)
	fhour := float64(hour)
	fmin := float64(min)

	fday := float64(day) + (fhour*3600.0+fmin*3600.0+sec)/86400.0
	y := fyear + 4800.0 // 4801 B.C. is a century year and also a leap year

	if year < 0 {
		y += 1.0
	}

	if month <= 2 {
		fmonth += 12.0
		y -= 1.0
	}

	e := math.Floor(30.6 * (fmonth + 1.0))

	a := math.Floor(y / 100.0) // number of centuries

	// 教皇格雷戈里十三世于 1582 年 2 月 24 日以教皇训令颁布，将 1582 年
	// 10 月 5 日至 14 抹掉。 1582 年 10 月 4 日过完后第二天是 10 月 15 日
	var b, c float64
	if year < 1582 || (year == 1582 && month < 10) ||
		(year == 1582 && month == 10 && day < 15) {
		b = -38.0
	} else {
		//number of century years that are not leap years
		b = math.Floor((a / 4.0) - a)
	}

	// Julian calendar years and leap years
	c = math.Floor(365.25 * y)

	jd = b + c + e + fday - 32167.5

	return
}

func Mjd2Cal(mjd float64) (year, month, day, hour, min int64, sec float64) {
	year, month, day, hour, min, sec = Jd2Cal(mjd + 2400000.5)
	return
}

func Cal2Mjd(year, month, day, hour, min int64, sec float64) (mjd float64) {
	jd := Cal2Jd(year, month, day, hour, min, sec)
	mjd = jd - 2400000.5
	return
}

func Jd2Mjd(jd float64) (mjd float64) {
	mjd = jd - 2400000.5
	return
}

func Mjd2Jd(mjd float64) (jd float64) {
	jd = mjd + 2400000.5
	return
}

func Mjd2Weekday(mjd float64) (wd int64) {
	//  借助 MJD，由公历年月日推算星期几，按照格里高利十三世的历法改革
	//  去掉 1582 年 10 月 5 日至 14 日
	//
	// 这个函数根据儒略日推算星期。
	// 根据教皇格雷戈里十三世的历法改革，儒略日中没有 1582 年 10 月 5 日至 14 日。
	// 1582 年 10 月 4 日过完后第二天是 10 月 15 日。
	// 这个函数在公元 1582 年 10 月 4 日 24:00 点之前使用儒略历， 公元 1582 年
	// 10 月 15 日 00:00 点之后使用公历
	//
	// 1 = 星期一
	// ...
	// 7 = 星期天

	//mjd := math.Floor(Cal2Mjd(year, month, day, 0, 0, 0.0))

	// 2009 年 3 月 9 日（MJD 54899）是星期一
	wd = int64(math.Mod(mjd-54899, 7.0)) + 1

	if wd <= 0 {
		wd += 7
	}

	return
}

func Cal2Weekday(year, month, day int64) (wd int64) {
	// 由公历年月日推算星期几，按照英国人的做法去
	// 掉 1752 年 9 月 3 日至 13 日 cal2wd2 不借助 MJD，由公历
	// 年月日推算星期几，去掉 1582 年 10 月 5 日至 14 日

	// 英国人在 1752 年才采用格利戈里历法。英国人去掉了 1752 年 9 月
	// 3 日至 13 日。
	// 这个函数在公元 1752 年 9 月 2 日 24:00 点之前使用儒略历， 公
	// 元 1752 年 9 月 14 日 00:00 点之后使用公历

	var A float64

	if (month == 1) || month == 2 {
		month += 12
		year -= 1
	}

	// 判断是否在 1752 年 9 月 3 日前
	if (year < 1752) || ((year == 1752) && (month < 9)) ||
		((year == 1752) && (month == 9) && (day < 3)) {
		// 1752 年 9 月 3 日前的公式
		A = math.Mod(
			(float64(day+2*month)+math.Floor(float64(3*(month+1)))/5.0)+
				float64(year)+math.Floor(float64(year)/4.0)+5,
			7)
	} else {
		A = math.Mod(
			(float64(day+2*month)+math.Floor(float64(3*(month+1)))/5.0)+
				float64(year)+math.Floor(float64(year)/4.0)-
				math.Floor(float64(year)/100.0)+
				math.Floor(float64(year)/400.0),
			7)
	}

	wd = int64(A) + 1

	return
}

func Cal2GPSweekday(year, month, day, hour, min int64, sec float64) (
	gpsweek, gpsweekday int64) {
	// 将公历 GPS 时间转换到 GPS 周和周内的秒

	mjd := Cal2Mjd(year, month, day, hour, min, sec)

	// GPS day count from MJD 44244
	elapse := mjd - 44244.0
	week := math.Floor(elapse / 7.0)
	elapse -= week * 7.0

	gpsweek = int64(week)
	gpsweekday = int64(mathematics.Round(elapse, 0))

	return
}

func Mjd2GPSweekdaysec(mjd float64) (
	gpsweek, gpsweekday, gpsweeksec int64) {

	// GPS day count from MJD 44244
	elapse := mjd - 44244.0
	week := math.Floor(elapse / 7.0)
	elapse -= week * 7.0

	gpsweek = int64(week)
	gpsweekday = int64(mathematics.Round(elapse, 0))
	gpsweeksec = int64(mathematics.Round(elapse*86400.0, 0))

	return

}

func GPSweekday2Cal(gpsweek, gpsweekday int64) (
	year, month, day, hour, min int64, sec float64) {
	// 将 GPS 周和周内的秒转换到公历 GPS 时间

	// GPS start from MJD 44244.0
	mjd := 44244.0 + float64(gpsweek*7+gpsweekday)

	return Mjd2Cal(mjd)
}

func Cal2GPSweeksec(year, month, day, hour, min int64, sec float64) (
	gpsweek, gpsweeksec int64) {
	// 将公历 GPS 时间转换到 GPS 周和周内的秒

	mjd := Cal2Mjd(year, month, day, hour, min, sec)

	// GPS day count from MJD 44244
	elapse := mjd - 44244.0
	week := math.Floor(elapse / 7.0)
	elapse -= week * 7.0

	gpsweek = int64(week)
	gpsweeksec = int64(mathematics.Round(elapse*86400.0, 0))

	return
}

func GPSweeksec2Cal(gpsweek, gpsweeksec int64) (
	year, month, day, hour, min int64, sec float64) {
	// 将 GPS 周和周内的秒转换到公历 GPS 时间

	// GPS start from MJD 44244.0
	mjd := 44244.0 + (float64(gpsweek*7*86400+gpsweeksec) / 86400.0)

	return Mjd2Cal(mjd)
}

func DateGPSweeksec2Cal(iyear, imonth, iday, igpsweeksec int64) (
	year, month, day, hour, min int64, sec float64) {
	// 由公历日期和 gps 周内秒计算公历 GPS 时间
	mjd := Cal2Mjd(iyear, imonth, iday, 0, 0, 0.0)

	week := int64(math.Floor((mjd - 44244.0) / 7.0))
	return GPSweeksec2Cal(week, igpsweeksec)
}

func DayofWeek(dow int64) string {
	if dow >= 0 && dow <= 6 {
		dows := map[int]string{
			0: "Sun",
			1: "Mon",
			2: "Tue",
			3: "Wed",
			4: "Thu",
			5: "Fri",
			6: "Sat",
		}
		return dows[int(dow)]
	}
	return ""
}
