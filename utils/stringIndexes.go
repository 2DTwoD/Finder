package utils

import "strings"

func GetAllRuneIndexesInString(text string, substring string) []int {
	result := make([]int, 0)
	textRunes := []rune(strings.ToLower(text))
	substringRunes := []rune(strings.ToLower(substring))
	substringLen := len(substringRunes)
	if substringLen == 0 {
		return result
	}
	textFinIndex := len(textRunes) - substringLen
	for i := 0; i < textFinIndex; i++ {
		if textRunes[i] != substringRunes[0] {
			continue
		}
		for n := 1; n < substringLen; n++ {
			if textRunes[i+n] != substringRunes[n] {
				break
			}
			if n == substringLen-1 {
				result = append(result, i)
			}
		}
	}
	return result
}
