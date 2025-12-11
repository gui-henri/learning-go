package util

import (
	"strconv"
	"strings"
)

func ParseValor(s string) float64 {
	if s == "" {
		return 0.0
	}
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, ",", ".")
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0
	}
	return val
}
