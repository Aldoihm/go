package main

import "fmt"

func main() {
	//declaración de un array de 3 posiciones y no se puede modificar el tamaño después de ser declarado
	/* var food [3]string
	food[0] = "🍔"
	food[1] = "🍕"
	food[2] = "🌭"
	fmt.Println(food) */

	/*food := [3]string{"🍔", "🍕", "🌭"}*/

	//para que GO pueda inferir el tamaño del array se hace de esta manera y cambiará el tamañao según el número
	//de elementos;

	food := [...]string{"🍔", "🍕", "🌭", "🥗", "🍑"}
	fmt.Println(food)

	//NOTA: 1 vez que GO ya determinó el tamaño del array, este no puede ser modificado.
}
