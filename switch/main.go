package main

import "fmt"

func main() {
	/* PRIMERA FORMA
	switch expresion {
	case condition :

	case condition :

	default:
	} */ //no es necesario la palabra reservada break

	emoji := "🐺"
	/* switch emoji {
	case "🐈", "🐕", "🐺":
		fmt.Println("es un animal")
	case "🍌", "🍎":
		fmt.Printf("es un fruta")
	default:
		fmt.Println("no es un animal ni una fruta")
	} */

	/* SEGUNDA FORMA:
	switch  {
	case operador_comparacion & operadores_logicos :

	case == || == :

	default:
	}
	*/

	switch {
	case emoji == "🐈" || emoji == "🐕" || emoji == "🐺":
		fmt.Println("es un animal")
	case emoji == "🍌" || emoji == "🍎":
		fmt.Printf("es un fruta")
	default:
		fmt.Println("no es un animal ni una fruta")
	}

}
