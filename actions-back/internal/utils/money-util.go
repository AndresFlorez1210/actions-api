package utils

import (
	"strconv"
	"strings"
)

func StringMoneyToFloat(money string) float64 {
	money = strings.TrimSpace(money)
	money = strings.TrimPrefix(money, "$")

	money = strings.ReplaceAll(money, ",", "")

	value, err := strconv.ParseFloat(money, 64)
	if err != nil {
		return 0
	}

	return value
}
