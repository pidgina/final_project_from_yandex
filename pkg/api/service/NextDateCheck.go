package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const LayoutDate string = "20060102"

var ErrYear = errors.New("Неверный формат лет")
var ErrMounth = errors.New("Неверный формат месяцев")
var ErrWeek = errors.New("Неверный формат недель")
var ErrDay = errors.New("Неверный формат дней")
var ErrRepeat = errors.New("Неверный формат повторения")

func CheckMounth(daysAll [34]bool, mounthAll [13]bool) bool {
	if daysAll[32] || daysAll[33] {
		return true
	}

	for month := 1; month <= 12; month++ {
		if !mounthAll[month] {
			continue
		}

		maxDay := time.Date(
			2000,
			time.Month(month)+1, 0, 0, 0, 0, 0, time.UTC).Day()

		for day := 1; day <= maxDay; day++ {
			if daysAll[day] {
				return true
			}
		}
	}

	return false
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse(LayoutDate, dstart)
	if err != nil {
		return "", err
	}

	if repeat == "" || len(repeat) < 1 {
		return "", fmt.Errorf("%w: repeat не может быть пустым", ErrRepeat)
	}

	rep := strings.Split(repeat, " ")
	if rep[0] != "d" && rep[0] != "y" && rep[0] != "w" && rep[0] != "m" {
		return "", fmt.Errorf("%w: Ошибка при проверке repeat, первый символ не является допустимым.", ErrRepeat)
	}

	switch rep[0] {

	case "m":
		if len(rep) != 2 && len(rep) != 3 {
			return "", fmt.Errorf("%w: repeat не может содержать меньше 1 и больше 3 параметров", ErrRepeat)
		}
		var daysAll [34]bool
		days := strings.Split(rep[1], ",")
		for _, dayMou := range days {
			day, err := strconv.Atoi(dayMou)
			if err != nil {
				return "", fmt.Errorf("%w: день месяца должен быть числом", ErrDay)
			}
			switch {
			case day >= 1 && day <= 31:
				daysAll[day] = true
			case day == -1:
				daysAll[32] = true
			case day == -2:
				daysAll[33] = true
			default:
				return "", fmt.Errorf("%w: некорректный день месяца", ErrDay)
			}
		}
		var mounthAll [13]bool
		if len(rep) == 2 {
			for mounth := 1; mounth <= 12; mounth++ {
				mounthAll[mounth] = true
			}
		} else {
			for _, v := range strings.Split(rep[2], ",") {
				mounth, err := strconv.Atoi(v)
				if err != nil {
					return "", fmt.Errorf("%w: месяц должен быть числом", ErrMounth)
				}
				if mounth < 1 || mounth > 12 {
					return "", fmt.Errorf("%w: месяц должен быть от 1 до 12", ErrMounth)
				}
				mounthAll[mounth] = true

			}
		}

		if !CheckMounth(daysAll, mounthAll) {
			return "", fmt.Errorf("%w: в выбранных месяцах нет указанных дней", ErrRepeat)
		}

		for {
			date = date.AddDate(0, 0, 1)
			if !date.After(now) {
				continue
			}
			day := date.Day()
			month := int(date.Month())
			if !mounthAll[month] {
				continue
			}
			isLastDay := date.AddDate(0, 0, 1).Day() == 1
			isPenultimateDay := date.AddDate(0, 0, 2).Day() == 1
			if daysAll[day] || (daysAll[32] && isLastDay) || (daysAll[33] && isPenultimateDay) {
				break
			}
		}
	case "w":
		if len(rep) != 2 {
			return "", fmt.Errorf("%w: repeat w не может содержать не 2 параметра", ErrRepeat)
		}
		var allowed [8]bool
		days := strings.Split(rep[1], ",")
		for _, dayStr := range days {
			day, err := strconv.Atoi(dayStr)
			if err != nil {
				return "", fmt.Errorf("%w: день недели должен быть числом", ErrWeek)
			}
			if day < 1 || day > 7 {
				return "", fmt.Errorf("%w: Передан невалидный день недели, нужно от 1 до 7", ErrWeek)
			}
			allowed[day] = true
		}

		for {
			date = date.AddDate(0, 0, 1)
			weekday := int(date.Weekday())
			if weekday == 0 {
				weekday = 7
			}
			if date.After(now) && allowed[weekday] {
				break
			}
		}

	case "d":

		if len(rep) != 2 {
			return "", fmt.Errorf("%w: Для повтора d нужен шаг: например, d 7", ErrDay)
		}

		step, err := strconv.Atoi(rep[1])
		if err != nil {
			return "", fmt.Errorf("%w: Некорректный шаг дней", ErrDay)
		}
		if step < 1 || step > 400 {
			return "", fmt.Errorf("%w: Шаг дней должен быть от 1 до 400", ErrDay)
		}

		for {
			date = date.AddDate(0, 0, step)
			if date.After(now) {
				break
			}
		}

	case "y":
		if len(rep) != 1 {
			return "", fmt.Errorf("%w: Для ежегодного повтора необходимо предоставить [y]", ErrDay)
		}

		for {
			date = date.AddDate(1, 0, 0)

			if date.After(now) {
				break
			}
		}
	}

	res := date.Format(LayoutDate)
	return res, nil

}
