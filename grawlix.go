package gograwlix

import (
	"math/rand"
	"strings"
)

var graws []rune = []rune{
	'%', '§', '&', '?',
}

func Grawlix(input string) string {
	var in []rune = []rune(strings.ToLower(input))
	for i, v := range in {
		switch v {
		case ' ':
			in[i] = ' '
		case 'e':
			in[i] = '*'
		case 'o':
			in[i] = '*'
		case 'u':
			in[i] = '*'
		case 'a':
			in[i] = '@'
		case 's':
			in[i] = '$'
		case 'h':
			in[i] = '#'
		case 'i':
			in[i] = '*'
		case 't':
			in[i] = '!'
		default:
			in[i] = graws[rand.Intn(len(graws))]
		}
	}
	return string(in)
}
