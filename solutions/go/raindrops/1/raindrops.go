package raindrops

import (
	"strconv"
	"strings"
)

func dividableBy3(number int) bool {
	return number%3 == 0
}

func dividableBy5(number int) bool {
	return number%5 == 0
}

func dividableBy7(number int) bool {
	return number%7 == 0
}

func Convert(number int) string {
	var result string
	if dividableBy3(number) {
		result = strings.Join([]string{result, "Pling"}, "")
	}
	if dividableBy5(number) {
		result = strings.Join([]string{result, "Plang"}, "")
	}
	if dividableBy7(number) {
		result = strings.Join([]string{result, "Plong"}, "")
	}

	if result == "" {
		return strconv.Itoa(number)
	}
	return result
}
