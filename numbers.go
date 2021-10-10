package goutils

import (
	"fmt"
	"strconv"
	"strings"
)

// FloatToString : converts float to string
func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', 16, 64)
}

// IntegerToString : converts int to string
func IntegerToString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

// StringToInteger : converts string to int
func StringToInteger(s string) int {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Printf("Fehler: %v \n", err)
		// v is 0 when invalid syntax
		// v is int.min or int.max when out of range
		return 0
	}
	return int(v)
}

// IsInteger : check number
func IsInteger(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return false
	}
	return true
}

// IFormat : formats an int to a string
func IFormat(i, length int, delimiter string) string {
	sign := ""
	if i < 0 {
		sign = "-"
		i = -i
	} else {
		sign = " "
	}
	iSlice := strings.Split(IntegerToString(i), "")
	formatted := ""
	d := "" // delimiter
	for {
		if len(iSlice) <= 3 { // 0...999
			formatted = strings.Join(iSlice, "") + d + formatted
			break
		} else { // 9123456789
			rest := iSlice[len(iSlice)-3:] // 789, 456, 124, 9
			iSlice = iSlice[:len(iSlice)-3]
			formatted = strings.Join(rest, "") + d + formatted
			d = delimiter
		}
	}
	if sign == "-" {
		formatted = sign + formatted
	}
	if len(formatted) < length {
		pref := strings.Repeat(" ", length-len(formatted))
		formatted = pref + formatted
	}
	return formatted
}
