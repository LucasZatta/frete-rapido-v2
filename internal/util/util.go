package util

import "regexp"

const (
	cnpjLen    = 14
	zipcodeLen = 8
)

var numericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func ClearString(str string) string {
	return numericRegex.ReplaceAllString(str, "")
}

func ValidateZipcode(fieldValue string) bool {
	return len(fieldValue) == zipcodeLen
}

func ValidateCNPJ(fieldValue string) bool {
	return len(fieldValue) == cnpjLen
}
