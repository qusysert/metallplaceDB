package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"math"
	"metallplace/internal/pkg/utils"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func findInRowRange(book *excelize.File, sheet string, column string, centerRow int, delta int, target string) (string, int, error) {
	for row := centerRow - delta; row < centerRow+delta; row++ {
		name, err := book.GetCellValue(sheet, column+strconv.Itoa(row))
		if err != nil {
			return "", 0, fmt.Errorf("cannot get name: %w", err)
		}
		if target == name {
			return name, row, nil
		}
	}
	return "", 0, fmt.Errorf("cannot find %v", target)
}

func getMonthDateForPredict(month string) (time.Time, error) {
	monthsMap := map[string]string{
		"Янв": "Jan",
		"Фев": "Feb",
		"Мар": "Mar",
		"Апр": "Apr",

		"Май": "May",
		"Июн": "Jun",
		"Июл": "Jul",
		"Авг": "Aug",
		"Сен": "Sep",
		"Окт": "Oct",
		"Ноя": "Nov",
		"Дек": "Dec",
	}
	format := "Jan'06"
	arr := strings.Split(month, "'")
	monthInEnglish := monthsMap[arr[0]]
	if monthInEnglish == "" {
		return time.Time{}, fmt.Errorf("error: Unrecognized month")
	}
	t, err := time.Parse(format, monthInEnglish+"'"+arr[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("Error parsing date: %w", err)
	}
	return t, nil
}

func getLastNotEmptyElement(slice []string) string {
	lastElement := ""
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] != "" {
			lastElement = slice[i]
			break
		}
	}
	return lastElement
}

func areNextCellsEmpty(book *excelize.File, sheet string, col int, row int, n int) (bool, error) {
	for i := row; i < row+n; i++ {
		c, err := utils.IntToAlphabet(int32(col))
		if err != nil {
			return false, err
		}
		value, err := book.GetCellValue(sheet, c+strconv.Itoa(i))
		if err != nil {
			return false, fmt.Errorf("cant get cell value in areNextCellsEmpty: %w", err)
		}
		if value != "" {
			return false, nil
		}
	}
	return true, nil
}

func stringToDate(str string, style string) (time.Time, error) {
	if style == "week" {
		re := regexp.MustCompile(`\((.*?)\)`)
		match := re.FindStringSubmatch(str)
		if len(match) > 1 {
			date, err := time.Parse("02.01.2006", match[1])
			if err != nil {
				return time.Time{}, fmt.Errorf("cant convert string %s to date: %w", match[1], err)
			}
			return date, nil
		} else {
			return time.Time{}, fmt.Errorf("cant parce date from string %s", str)
		}
	}
	if style == "weekNum" {
		arr := strings.Split(str, " ")
		week, err := strconv.Atoi(arr[0])
		if err != nil {
			return time.Time{}, fmt.Errorf("cant parce week (%v): %w", arr[0], err)
		}
		yearStr := arr[1]
		yearStr = yearStr[1 : len(yearStr)-1]
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			return time.Time{}, fmt.Errorf("cant parce year: %w", err)
		}
		mon := firstDayOfISOWeek(year, week)
		return mon, nil
	}
	if style == "month" {
		arr := strings.Split(str, " ")
		monthStr := arr[0]
		month, err := monthStrToNumber(monthStr)
		if err != nil {
			return time.Time{}, fmt.Errorf("cant parce month: %w", err)
		}
		year, err := strconv.Atoi(arr[1])
		if err != nil {
			return time.Time{}, fmt.Errorf("cant parce year: %w", err)
		}
		return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC), nil
	}
	return time.Time{}, fmt.Errorf("wrong style")

}

func firstDayOfISOWeek(year int, week int) time.Time {
	date := time.Date(year, 0, 0, 0, 0, 0, 0, time.UTC)
	isoYear, isoWeek := date.ISOWeek()
	for date.Weekday() != time.Monday { // iterate back to Monday
		date = date.AddDate(0, 0, -1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoYear < year { // iterate forward to the first day of the first week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoWeek < week { // iterate forward to the first day of the given week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	return date
}

func GetDateFiveWeeksAgo(currentDate time.Time) time.Time {
	// Вычитаем 5 недель из текущей даты
	fiveWeeksAgo := currentDate.AddDate(0, 0, -35)

	// Получаем день недели исходной даты
	weekday := currentDate.Weekday()

	// Ищем ближайший день недели, соответствующий исходной дате
	for i := 0; i < 7; i++ {
		if fiveWeeksAgo.Weekday() == weekday {
			break
		}
		fiveWeeksAgo = fiveWeeksAgo.AddDate(0, 0, -1)
	}

	return fiveWeeksAgo
}

func formatMonth(input string) string {
	year := ""

	arr := strings.Split(input, "-")
	arr[0] = strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(arr[0], ".", ""), " ", ""))

	if len(arr) == 2 {
		year = "-" + arr[1]
	}

	switch {
	case contains([]string{"i", "янв", "jan", "январь"}, arr[0]):
		return "Янв" + year
	case contains([]string{"ii", "фев", "февраль", "feb"}, arr[0]):
		return "Фев" + year
	case contains([]string{"iii", "мар", "март", "mar"}, arr[0]):
		return "Мар" + year
	case contains([]string{"iv", "апр", "апрель", "apr"}, arr[0]):
		return "Апр" + year
	case contains([]string{"v", "май", "may"}, arr[0]):
		return "Май" + year
	case contains([]string{"vi", "июн", "июнь", "jun"}, arr[0]):
		return "Июн" + year
	case contains([]string{"vii", "июл", "июль", "jul"}, arr[0]):
		return "Июл" + year
	case contains([]string{"viii", "iix", "авг", "август", "aug"}, arr[0]):
		return "Авг" + year
	case contains([]string{"ix", "сен", "сентябрь", "сент", "sep"}, arr[0]):
		return "Сен" + year
	case contains([]string{"x", "х", "окт", "октябрь", "oct"}, arr[0]):
		return "Окт" + year
	case contains([]string{"xi", "ноя", "ноябрь", "нояб", "nov"}, arr[0]):
		return "Ноя" + year
	case contains([]string{"xii", "дек", "декабрь", "dec"}, arr[0]):
		return "Дек" + year
	default:
		return "undefined"
	}
}

func monthStrToNumber(month string) (int, error) {
	switch strings.ToLower(month) {
	case "январь":
		return 1, nil
	case "февраль":
		return 2, nil
	case "март":
		return 3, nil
	case "апрель":
		return 4, nil
	case "май":
		return 5, nil
	case "июнь":
		return 6, nil
	case "июль":
		return 7, nil
	case "август":
		return 8, nil
	case "сентябрь":
		return 9, nil
	case "октябрь":
		return 10, nil
	case "ноябрь":
		return 11, nil
	case "декабрь":
		return 12, nil
	}
	return 0, fmt.Errorf("wrong month string")
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func round(num float64, n int) float64 {
	digits := math.Pow(10, float64(n))
	return math.Round((num)*digits) / digits
}

func addSpacesToNumber(number float64) string {
	numberStr := strconv.FormatFloat(number, 'f', -1, 64)
	parts := strings.Split(numberStr, ".")
	integerPart := parts[0]
	decimalPart := ""
	if len(parts) > 1 {
		decimalPart = "." + parts[1]
	}

	sign := ""
	if integerPart[0] == '-' {
		sign = "-"
		integerPart = integerPart[1:]
	}

	length := len(integerPart)
	if length <= 3 {
		return sign + numberStr
	}

	var result strings.Builder
	for i := 0; i < length; i++ {
		if i > 0 && (length-i)%3 == 0 {
			result.WriteString(" ")
		}
		result.WriteByte(integerPart[i])
	}

	return sign + result.String() + decimalPart
}

func removeDuplicateLabels(input []string) []string {
	output := make([]string, len(input))

	prev := ""
	for i, v := range input {
		if i > 0 && v == prev {
			output[i] = ""
		} else {
			output[i] = v
		}
		prev = v
	}
	return output
}
