package ukr2lattransliteration

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	ukrAlphabet = "АБВГҐДЕЄЖЗИІЇЙКЛМНОПРСТУФХЦЧШЩЮЯабвгґдеєжзиіїйклмнопрстуфхцчшщьюя'’ʼ"
)

type diffRune struct {
	First string
	Other string
}

var (
	simpleRunes = map[rune]string{
		'А': "A",
		'Б': "B",
		'В': "V",
		'Г': "H",
		'Ґ': "G",
		'Д': "D",
		'Е': "E",
		'Ж': "Zh",
		'З': "Z",
		'И': "Y",
		'І': "I",
		'К': "K",
		'Л': "L",
		'М': "M",
		'Н': "N",
		'О': "O",
		'П': "P",
		'Р': "R",
		'С': "S",
		'Т': "T",
		'У': "U",
		'Ф': "F",
		'Х': "Kh",
		'Ц': "Ts",
		'Ч': "Ch",
		'Ш': "Sh",
		'Щ': "Shch",
		'а': "a",
		'б': "b",
		'в': "v",
		'г': "h",
		'ґ': "g",
		'д': "d",
		'е': "e",
		'ж': "zh",
		'з': "z",
		'и': "y",
		'і': "i",
		'к': "k",
		'л': "l",
		'м': "m",
		'н': "n",
		'о': "o",
		'п': "p",
		'р': "r",
		'с': "s",
		'т': "t",
		'у': "u",
		'ф': "f",
		'х': "kh",
		'ц': "ts",
		'ч': "ch",
		'ш': "sh",
		'щ': "shch",
	}

	difficultRunes = map[rune]diffRune{
		'є': diffRune{First: "Ye", Other: "ie"},
		'ї': diffRune{First: "Yi", Other: "i"},
		'й': diffRune{First: "Y", Other: "i"},
		'ю': diffRune{First: "Yu", Other: "iu"},
		'я': diffRune{First: "Ya", Other: "ia"},
	}
)

func Transliteration(input string) string {
	var sb strings.Builder
	runes := []rune(input)
	skipNext := false

	for i, r := range runes {

		if skipNext {
			skipNext = false
			continue
		}

		if !strings.ContainsRune(ukrAlphabet, r) {
			sb.WriteRune(r)
			continue
		}

		if strings.ContainsRune("ь'’ʼ", r) {
			continue
		}

		if unicode.ToLower(r) == 'з' {
			if (i+1 < len(runes)) && (unicode.ToLower(runes[i+1]) == 'г') {
				if unicode.IsUpper(r) {
					sb.WriteString("Zgh")
				} else {
					sb.WriteString("zgh")
				}
				skipNext = true
				continue
			}
		}

		str, ok := simpleRunes[r]
		if ok {
			sb.WriteString(str)
			continue
		}

		dr, ok := difficultRunes[unicode.ToLower(r)]
		if ok {
			if (i == 0) || (!strings.ContainsRune(ukrAlphabet, runes[i-1])) { // first letter in word
				if unicode.IsUpper(r) {
					sb.WriteString(dr.First)
				} else {
					sb.WriteString(strings.ToLower(dr.First))
				}
			} else {
				if unicode.IsUpper(r) {
					sb.WriteString(strings.ToTitle(dr.Other))
				} else {
					sb.WriteString(dr.Other)
				}
			}
			continue
		}

		panic(fmt.Sprintf("Undefined ukrainian letter: %c", r))

	}

	return sb.String()
}
