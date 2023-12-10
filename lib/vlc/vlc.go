package vlc

import (
	"strings"
	"unicode"
)

func Encode(str string) string {
	str = prepareText(str)

	// some text -> 00101001

	// split binary by chunks

	// bytes to hex

	// return string
	return ""
}

// prepateText prepares text to be fir for encode:
// changes upper case letters to : !+ lower case letter
// i.g: My bame is Ted => !my name is !ted
func prepareText(str string) string {
	var buf strings.Builder

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(ch))
		} else {
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}
