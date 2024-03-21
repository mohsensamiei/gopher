package mathext

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func RoundDecimal(v float64) float64 {
	return math.Round(v*100) / 100
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
