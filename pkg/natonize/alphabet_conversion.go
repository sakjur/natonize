package natonize

import (
	"strings"
	"errors"
	"fmt"
)

var alphabet = map[string]string{
	"0": "zero",
	"1": "one",
	"2": "two",
	"3": "three",
	"4": "four",
	"5": "five",
	"6": "six",
	"7": "seven",
	"8": "eight",
	"9": "nine",
	"a": "alpha",
	"b": "bravo",
	"c": "charlie",
	"d": "delta",
	"e": "echo",
	"f": "foxtrot",
	"g": "golf",
	"h": "hotel",
	"i": "india",
	"j": "juliett",
	"k": "kilo",
	"l": "lima",
	"m": "mike",
	"n": "november",
	"o": "oscar",
	"p": "papa",
	"q": "quebec",
	"r": "romeo",
	"s": "sierra",
	"t": "tango",
	"u": "uniform",
	"v": "victor",
	"w": "whisky",
	"x": "x-ray",
	"y": "yankee",
	"z": "zulu",
	" ": "space",
	"-": "dash",
	".": "period",
	",": "comma",
	"!": "bang",
}

func ToNatoAlphabet(str string) string {
	var characters []string

	for _, char := range str {
		natoCharacter, _ := convertCharacter(string(char))

		if natoCharacter == "" {
			natoCharacter = fmt.Sprintf("unknown(%s)", string(char))
		}

		characters = append(characters, natoCharacter)
	}

	return strings.Join(characters, " ")
}

func FromNatoAlphabet(natoString string) string {
	characters := strings.Split(natoString, " ")
	reverseAlphabet := reverseAlphabet()
	regularString := ""

	for _, char := range characters {
		regularString += reverseAlphabet[char]
	}

	return regularString
}

func convertCharacter(c string) (string, error) {
	if len(c) != 1 {
		return "", errors.New("must convert single character")
	}

	outString := ""

	if (strings.ToLower(c) != c) {
		outString += "capital-"
	}
	outString += alphabet[strings.ToLower(c)]

	return outString, nil
}

func reverseAlphabet() map[string]string {
	reverseAlphabet := make(map[string]string, len(alphabet))

	for key, val := range alphabet {
		reverseAlphabet[val] = key

		if (strings.ToUpper(key) != key) {
			reverseAlphabet["capital-"+val] = strings.ToUpper(key)
		}
	}

	return reverseAlphabet
}