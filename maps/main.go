package main

import "fmt"

func main() {
	//declaración de un mapa
	animals := make(map[string]string) //[tipo de dato de la llave]tipo de dato que almacenera el valor
	//asignación de valor
	animals["cat"] = "🐈"
	animals["dog"] = "🐕"

	fmt.Println(animals)
	//construir un mapa literal

	fruits := map[string]string{
		"apple":  "🍎",
		"banana": "🍌",
	}
	fmt.Println(fruits)
	//eliminar elementos

	delete(fruits, "banana")
	fmt.Println(fruits)

	//obtener un elemento del mapa
	fmt.Println(animals["cat"])

	//recuperar un elmento que no hemos creado
	fmt.Println(animals["gorilla"]) //retorna el valor 0, del tipo del dato del valor
	/* Esta sintaxis: animals["gorilla"], retorna 2 valores:
	Valor_1.- Es el contenido
	Valor_2.- Esa llave existe o no, es un valor bool
	*/
	//Queda así: valor_1, valor_2: animals["gorilla"] retornaría "  false"
	emoji, ok := animals["dog"]
	fmt.Println(emoji, ok)

	_, okei := animals["gorilla"]
	if !okei {
		animals["gorilla"] = "🦍"
	}
	fmt.Println(animals)
}
