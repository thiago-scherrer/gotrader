package convert

import (
	"fmt"
	"strconv"
)

// BytesToString convert bytes to string
func BytesToString(data []byte) string {
	return string(data[:])
}

// StringToBytes convert string to byte
func StringToBytes(data string) []byte {
	return []byte(data[:])
}

// StringToInt convert string to int64
func StringToInt(data string) int64 {
	result, _ := strconv.ParseInt(data, 10, 64)
	return result
}

// IntToString convert string to int64
func IntToString(data int64) string {
	return strconv.FormatInt(int64(data), 10)
}

// FloatToString convert float64 to string
func FloatToString(data float64) string {
	return fmt.Sprintf("%f", data)
}

// StringToFloat64 convert string to float64
func StringToFloat64(data string) float64 {
	result, _ := strconv.ParseFloat(data, 64)
	return result
}

// FloatToInt convert int to float64
func FloatToInt(data float64) int64 {
	return int64(data)
}
