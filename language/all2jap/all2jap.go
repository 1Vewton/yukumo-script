package all2jap

import (
	"regexp"

	"github.com/1Vewton/yukumo-script/language/convertnums"
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

// EngToKana converts English to Kana
func EngToKana(text string) string {
	numResult := convertnums.ConvertNumToEnglish(text)
	return AllToKana(numResult)
}
