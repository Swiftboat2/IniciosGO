package main

import (
	"fmt"
)

const (
	Nombre = "Lauty "

	Español = "español"
	Frances = "frances"

	HelloIngles = "Hello "
	Holaespañol = "Hola "
	Bounjour    = "Bounjour "
)

func Yo(apellido string, language string) string {
	if apellido == "" {
		apellido = "Pelozo"
	}
	/*
		if language == Español {
			return Holaespañol + Nombre + apellido
		}

		if language == Bounjour {
			return Bounjour + Nombre + apellido
		}
	*/

	return saludandoPrefix(language) + Nombre + apellido
}

func saludandoPrefix(language string) (prefix string) {

	switch language {
	case Español:
		prefix = Holaespañol
	case Frances:
		prefix = Bounjour
	default:
		prefix = HelloIngles
	}

	return
}

func main() {

	fmt.Println(Yo("", ""))

}
