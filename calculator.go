package twilight

import (
	"math"
	"time"
)

// RadianFactor ...
const (
	RadianFactor = math.Pi / 180
	ToRadian     = RadianFactor
	ToDegree     = 1 / RadianFactor

	ZenithOfficial     = 90 + 50/60
	ZenithCivil        = 96
	ZenithNautical     = 102
	ZenithAstronomical = 108
)

// Zenith ...
var Zenith = map[DuskType]float64{
	DuskTypeBorders:      90.3,
	DuskTypeSimple:       90.83,
	DuskTypeMid:          93,
	DuskTypeCivil:        96,
	DuskTypeNautical:     102,
	DuskTypeAstronomical: 108,
}

// CalcRise вспомогательная функция для вычисления только времени восхода.
func CalcRise(latitute, longitude float64, dusk DuskType, year, month, day int) (time.Time, error) {
	from, _, err := Calc(latitute, longitude, dusk, year, month, day)
	if err != nil {
		return time.Time{}, err
	}
	return from, nil
}

// CalcSet вспомогательная функция для вычисления только времени заката.
func CalcSet(latitute, longitude float64, dusk DuskType, year, month, day int) (time.Time, error) {
	_, to, err := Calc(latitute, longitude, dusk, year, month, day)
	if err != nil {
		return time.Time{}, err
	}
	return to, nil
}

// Calc основная функция для вычисления диапазона сумерек.
// longitude should be positive for East and negative for West
func Calc(latitute, longitude float64, dusk DuskType, year, month, day int) (from, to time.Time, err error) {

	// проверка входных параметров
	// возврат error
	// -- переход через ночь -> 2 отдельные функции
	// -- timezone / timeoffset -> возвращаем в UTC
	// проверка условий
	// 		if (cosH >  1) the sun never rises on this location (on the specified date)
	// 		if (cosH < -1) the sun never sets on this location (on the specified date)

	// http://www.edwilliams.org/sunrise_sunset_algorithm.htm
	// https: //www.edwilliams.org/sunrise_sunset_example.htm

	// 1. the day of the year
	dayOfYear := calcDayOfYear(year, month, day)
	// fmt.Println(dayOfYear)

	// 2. sunrise/sunset approximate time
	longHour := longitude / 15
	timeRise := dayOfYear + ((6 - longHour) / 24)
	timeSet := dayOfYear + ((18 - longHour) / 24)
	// fmt.Println("timeRise:", timeRise)

	// 3. the Sun's mean anomaly
	meanRise := (0.9856 * timeRise) - 3.289
	meanSet := (0.9856 * timeSet) - 3.289
	// fmt.Println("meanRise:", meanRise)

	// 4. Sun's true longitude
	longRise := meanRise + 1.916*math.Sin(meanRise*ToRadian) + 0.020*math.Sin(2*meanRise*ToRadian) + 282.634
	longSet := meanSet + 1.916*math.Sin(meanSet*ToRadian) + 0.020*math.Sin(2*meanSet*ToRadian) + 282.634

	// potentially needs to be adjusted into the range [0,360) by adding/subtracting 360
	if longRise >= 360 {
		longRise = longRise - 360
	}
	if longRise < 0 {
		longRise = longRise + 360
	}
	if longSet >= 360 {
		longSet = longSet - 360
	}
	if longSet < 0 {
		longSet = longSet + 360
	}
	// fmt.Println("longRise:", longRise)

	// 5. the Sun's right ascension
	ascensionRise := ToDegree * (math.Atan(0.91764 * math.Tan(longRise*ToRadian)))
	ascensionSet := ToDegree * (math.Atan(0.91764 * math.Tan(longSet*ToRadian)))
	// potentially needs to be adjusted into the range [0,360) by adding/subtracting 360
	if ascensionRise >= 360 {
		ascensionRise = ascensionRise - 360
	}
	if ascensionRise < 0 {
		ascensionRise = ascensionRise + 360
	}
	if ascensionSet >= 360 {
		ascensionSet = ascensionSet - 360
	}
	if ascensionSet < 0 {
		ascensionSet = ascensionSet + 360
	}
	// right ascension value needs to be in the same quadrant as true longitude
	longRiseQuadrant := (math.Floor(longRise / 90)) * 90
	ascensionRiseQuadrant := (math.Floor(ascensionRise / 90)) * 90
	ascensionRise = ascensionRise + (longRiseQuadrant - ascensionRiseQuadrant)

	longSetQuadrant := (math.Floor(longSet / 90)) * 90
	ascensionSetQuadrant := (math.Floor(ascensionSet / 90)) * 90
	ascensionSet = ascensionSet + (longSetQuadrant - ascensionSetQuadrant)

	// right ascension value needs to be converted into hours
	ascensionRise = ascensionRise / 15
	ascensionSet = ascensionSet / 15
	// fmt.Println("ascensionRise:", ascensionRise)

	// 6. the Sun's declination
	sinDeclinationRise := 0.39782 * math.Sin(longRise*ToRadian)
	cosDeclinationRise := math.Cos(math.Asin(sinDeclinationRise))
	sinDeclinationSet := 0.39782 * math.Sin(longSet*ToRadian)
	cosDeclinationSet := math.Cos(math.Asin(sinDeclinationSet))
	// fmt.Println("sinDeclinationRise:", sinDeclinationRise)
	// fmt.Println("cosDeclinationRise:", cosDeclinationRise)

	// 7. the Sun's local hour angle
	zenith := Zenith[dusk]
	hourAngleRise := (math.Cos(zenith*ToRadian) - sinDeclinationRise*math.Sin(latitute*ToRadian)) / (cosDeclinationRise * math.Cos(latitute*ToRadian))
	hourAngleSet := (math.Cos(zenith*ToRadian) - sinDeclinationSet*math.Sin(latitute*ToRadian)) / (cosDeclinationSet * math.Cos(latitute*ToRadian))
	// fmt.Println("hourAngleRise:", hourAngleRise)
	// fmt.Println("hourAngleSet:", hourAngleSet)
	// if (cosH >  1) the sun never rises on this location (on the specified date)
	// if (cosH < -1) the sun never sets on this location (on the specified date)

	// calculating H and convert into hours
	HRise := 360 - ToDegree*(math.Acos(hourAngleRise))
	HSet := ToDegree * (math.Acos(hourAngleSet))
	HRise = HRise / 15
	HSet = HSet / 15
	// fmt.Println("HRise:", HRise)

	// 8. local mean time of rising/setting
	TRise := HRise + ascensionRise - (0.06571 * timeRise) - 6.622
	TSet := HSet + ascensionSet - (0.06571 * timeSet) - 6.622
	// fmt.Println("TRise:", TRise)

	// 9. adjust back to UTC
	UTRise := TRise - longHour
	UTSet := TSet - longHour
	// potentially needs to be adjusted into the range [0,24) by adding/subtracting 24
	if UTRise >= 24 {
		UTRise = UTRise - 24
	}
	if UTRise < 0 {
		UTRise = UTRise + 24
	}
	if UTSet >= 24 {
		UTSet = UTSet - 24
	}
	if UTSet < 0 {
		UTSet = UTSet + 24
	}
	// fmt.Println("UTRise:", UTRise, dusk)

	h, s := math.Modf(UTRise)
	from = time.Date(year, time.Month(month), day, int(h), int(math.Round(s*60)), 0, 0, time.UTC)
	h, s = math.Modf(UTSet)
	to = time.Date(year, time.Month(month), day, int(h), int(math.Round(s*60)), 0, 0, time.UTC)

	// 10. convert to local time zone of latitude/longitude
	// localT = UT + localOffset

	return from, to, nil
}

func calcDayOfYear(year, month, day int) float64 {
	v1 := math.Floor(275 * float64(month) / 9)
	v2 := math.Floor((float64(month) + 9) / 12)
	v3 := (1 + math.Floor((float64(year)-4*math.Floor(float64(year)/4)+2)/3))
	res := v1 - (v2 * v3) + float64(day) - 30
	return res
}

func toRadian(degree float64) float64 {
	return degree * RadianFactor
}

func toDegree(radian float64) float64 {
	return radian / RadianFactor
}
