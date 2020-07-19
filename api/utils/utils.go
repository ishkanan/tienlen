package utils

import (
	"encoding/json"
	"fmt"
	"math"
)

// Ordinal provides an English Ordinal representation of a number
func Ordinal(num int) string {
	ords := map[int]string{0: "th", 1: "st", 2: "nd", 3: "rd", 4: "th", 5: "th", 6: "th", 7: "th", 8: "th", 9: "th"}
	positiveNum := int(math.Abs(float64(num)))
	if ((positiveNum % 100) >= 11) && ((positiveNum % 100) <= 13) {
		return "th"
	}
	return fmt.Sprintf("%d%s", num, ords[positiveNum])
}

// PrettyPrint prints a pretty version of i using the JSON marshaller
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}
