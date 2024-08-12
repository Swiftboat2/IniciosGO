package iteration

const (
	contador = 5
)

func Repeat(character string) string {
	var repeated string
	for i := 0; i < contador; i++ {
		repeated += character
	}
	return repeated
}
