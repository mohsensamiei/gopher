package mathext

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func RoundN(v float64, n int) float64 {
	return math.Round(v*math.Pow10(n)) / math.Pow10(n)
}

func RoundDecimal(v float64) float64 {
	return RoundN(v, 2)
}

func DecimalToInt64(v float64) int64 {
	i, err := strconv.ParseInt(strings.Replace(fmt.Sprintf("%.2f", v), ".", "", -1), 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func Int64ToDecimal(v int64) float64 {
	return float64(v) / 100
}
