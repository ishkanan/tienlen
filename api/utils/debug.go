package utils

import "encoding/json"

// PrettyPrint prints a pretty version of i using the JSON marshaller
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}
