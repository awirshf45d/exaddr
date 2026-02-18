package utils

import "regexp"

// Split text by special characters
func splitBySpecialCharacters(text string) []string {
	re := regexp.MustCompile(`[~!$^*(){}\[\]:,/\s]+`)
	return re.Split(text, -1)
}
