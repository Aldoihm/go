package main

import "fmt"

// En go, todas las variables que declaramos, tienen que ser utilizadas para que no marque error
func main() {
	/* Método largo para declarar variables
	var apple string
	var banana string
	var orange string

	apple="🍎"
	banana="🍌"
	orange="🍊"
	*/

	/* Método para declarar una variable en una sola línea
	var apple string = "🍎"
	var banana string = "🍌"
	var orange string ="🍊"
	*/

	/*  Agrupar variables
	var (
		apple  string = "🍎"
		banana string = "🍌"
		orange string = "🍊"
	)
	*/
	/*declarar todas las variables en una sola línea
	var apple, banana, orange string = "🍎", "🍌", "🍊"
	*/

	//Operador de asignación de variable corta
	//Go determina dinamicamente el tipo de datos, dependiendo de lo que asignes a la variable
	apple, banana, orange := "🍎", "🍌", "🍊"
	//una vez que se asigna el tipo de dato a la variable, ya no se puede modificar el tipo de dato.

	/*para cambiar el valor de la variable ya no se usa el operador de asignación de variable corta, sino el igual
	forma correcta
	apple = "manznita"
	forma incorreta porque ya se había asignado anteriormente un valor
	apple := "manzanita"
	*/

	//se puede cambiar el valor de una variable usando el operador (:=) siempre y cuando esté jutno con una nueva variable
	apple, lemon := "manzana", "🍋"
	fmt.Println(apple, banana, orange, lemon)
}
