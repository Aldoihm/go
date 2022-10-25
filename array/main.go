package main

import "fmt"

func main() {
	//declaración de un array de 3 posiciones y no se puede modificar el tamaño después de ser declarado
	/* var food [3]string
	food[0] = "🍔"
	food[1] = "🍕"
	food[2] = "🌭"
	fmt.Println(food) */

	food := [3]string{"🍔", "🍕", "🌭"}
	fmt.Println(food)

}
