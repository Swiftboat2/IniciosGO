package main

import "fmt"

const HelloIngles = "Hello "
const Nombre = "Lauty "
const Español = "español"
const Holaespañol = "Hola "
const Bounjour = "Bounjour "

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

	var prefix string

	switch language {
	case Español:
		prefix = Holaespañol
	case Bounjour:
		prefix = Bounjour
	default:
		prefix = HelloIngles

	}

	return prefix + Nombre + apellido
}

func main() {

	fmt.Println(Yo("", ""))

}
