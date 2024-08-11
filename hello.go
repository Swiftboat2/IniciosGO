package main

import "fmt"

func Hello(name string) string {
	return "Hola " + name

}

func main() {
	fmt.Println(Hello("mundo"))
}
