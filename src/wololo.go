package main

import "strconv"

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
