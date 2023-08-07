package main

import "fmt"

// slice permite trabajar con arrays de forma dinámica
// slice son apuntadores de arrays
// los apuntadores no almacenan ningun dato
func main() {
	fmt.Println("PRIMERA PARTE")
	set := [7]string{"🦁", "🐎", "🐄", "🦋", "🦜", "🛩️", "🛫"}
	//declaración de un slice
	animals := set[0:5] // "🦁", "🐎", "🐄", "🦋", "🦜",
	//animals := set[:5] Go entiende que empiezo desde 0, es lo mismo que arriba
	//el elemento de la última posición de los slice es excluyente
	fly := set[3:7] //  "🦋", "🦜", "🛩️", "🛫"
	//fly := set[3:] Go entiende que termino en 7 es lo mismo que él de arriba
	todo := set[:] //Go entiende que empiezo desde 0 y termino en 7
	fmt.Println("Array: ", set)
	fmt.Println("Animales: ", animals)
	fmt.Println("Voladores: ", fly)
	fmt.Println("Todo: ", todo)
	//acceder a un valor del slice
	fmt.Println("fly[0]:", fly[0])
	//Editamos fly en la posición 0
	fly[0] = "🦅"
	fmt.Println("Valor de fly[0] después de ser editado:", fly[0])

	// NOTA IMPORTANTE: Recordar que los slice, en el último valor en la declaración es excluyente

	//Los dos puntos es propio del array, no de los slice
	//Se vuelve a imprimir, para observar que desde el slice se modifica el array y afecta a los demás punteros.
	fmt.Println("Array: ", set)
	fmt.Println("Animales: ", animals)
	fmt.Println("Voladores: ", fly)
	fmt.Println("Todo: ", todo)

	//PARTE 2

	// len(): # de elementos en el slice
	// cap(): # de elementos del array donde apunta el slice, a partir del índice de donde se creo el slice
	fmt.Println("\n\nSEGUNDA PARTE")
	food := [5]string{"🌭", "🍓", "🍋", "🍔", "🍕"}
	fruits := food[1:3] //"🍓","🍋"
	//agregar un elemento al slice
	fruits = append(fruits, "🍍", "🍈", "🍎")
	//Si agregamos al slice <= elementos del array original entonces
	//Array[4]{"🍓", "🍋", "🍔", "🍕"}
	//Sino
	//NewArray[8]{"🍓","🍋","🍍", "🍈", "🍎"} Se multiplico el array x 2
	fmt.Println("Food: ", food)
	fmt.Println("Fruits: ", fruits)
	fmt.Println("# de elementos en el slice Fruit: ", len(fruits))
	fmt.Println("# de elementos del array a partir del índice (CAPACIDAD): ", cap(fruits))
	//para declarar un slice haciendo referencia a un array nuevo
	//corchetes vacíos = slice, corchetes con número = array
	fmt.Println("\n\nCreando Sclice vacío")
	//fruta := []string{"🍓", "🍋"} Se crea un slice de tamaño 2 y capacidad 2
	fruta := make([]string, 0, 3) //se crea un slice con tamaño 0 y capacidad de 3
	fruta = append(fruta, "🍍", "🍓", "🍋", "🍎")
	fmt.Println("Fruta: ", fruta)
	fmt.Println("# de elementos en el slice Fruta: ", len(fruta))
	fmt.Println("# de elementos del array a partir del índice (CAPACIDAD): ", cap(fruta))

	//Creando Slice Vacios
	var slice []string
	//slice:=[]string{} No es lo mismo que el de arriba, arriba es nulo, este ya esta inicializado
	fmt.Println("slice: ", slice)
	fmt.Println("# de elementos en el slice Slice: ", len(slice))
	fmt.Println("# de elementos del array a partir del índice (CAPACIDAD): ", cap(slice))
	//nil == a nulo
	fmt.Println(slice == nil)
}
