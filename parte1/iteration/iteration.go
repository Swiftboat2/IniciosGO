package iteration

import (
	"strings"
)

func Repeat(character string, contador int) string {
	var repeated string
	for i := 0; i < contador; i++ {
		repeated += character
	}
	return repeated
}

func Repetir(c string, co int) string {

	Repeti := strings.Repeat(c, co)

	return Repeti
}
