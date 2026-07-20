package convertnums

import (
	"strings"
)

// enList defines the transition from numbers to english
var enList = map[string]string{
	"1": " one ",
	"2": " two ",
	"3": " three ",
	"4": " four ",
	"5": " five ",
	"6": " six ",
	"7": " seven ",
	"8": " eight ",
	"9": " nine ",
	"0": " zero ",
}

// ConvertNumToEnglish converts numbers to English vocabularies
func ConvertNumToEnglish(text string) string {
	result := text
	for num, vocab := range enList {
		result = strings.ReplaceAll(result, num, vocab)
	}
	return result
}
