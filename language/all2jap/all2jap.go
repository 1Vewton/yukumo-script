package all2jap

import (
	"regexp"

	kanatrans "github.com/Luigi-Pizzolito/English2KanaTransliteration"
)

// AllToKana converts English and Japanese characters to Kana
func AllToKana(text string) string {
	allToKana := kanatrans.NewAllToKana(true)
	convertResult := allToKana.Convert(text)
	re := regexp.MustCompile(`[^\p{Katakana}]+`)
	result := re.ReplaceAllString(convertResult, "")
	return result
}
