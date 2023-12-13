package main

import "fmt"

// Funciones de forma generica, nos deja trabajar con cualquier tipo de dato
func main() {
	PrintList("SotecMovil", "Gophers", "Go desde cero")
	PrintList(1, 2, 3)
}

func PrintList(list ...any) { //palabra clave para poder trabajar con multiples tipos de datos
	for _, item := range list {
		fmt.Println(item)
	}
}
