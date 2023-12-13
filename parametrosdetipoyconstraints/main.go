package main

import "fmt"

type MyInt int

func main() {
	fmt.Println("Restricciones de unión de elementos")
	fmt.Println(suma(2, 5.5, 7))
	fmt.Println(suma[int](2, 8, 7))
	fmt.Println(suma[float64](2, 5.5, 7))

	var num1 MyInt = 7
	var num2 MyInt = 8
	var num3 MyInt = 5
	fmt.Println("Restricciones de aproximación de elementos")
	fmt.Println(sum(num1, num2))
	fmt.Println(sum[MyInt](num3, 20))
	fmt.Println(sum(2, 4.6))

}

//Constrains, en español significa restricciones y son para las funciones de any
//las constrains van dentro de los corchetes []. Llamados como parametros de tipo
//Los constraints se pueden usar tanto en la firma de la función como en el cuerpo de la función
//Los constrains reciben un nombre que empieza con mayúscula por convención
/* 3 tipos de constraints
1.- Arbitrario
2.- De unión de elementos
3.- Aproximación de elementos
*/

// Las constraints arbitrarias son una pendejada inservible, por eso no las puse.
// Constraints de unión de elementos con |
func suma[T int | float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// Constraints de aproximación de elementos
// se añade el simbolo ~ para indicar para cualquier tipo derivado de entero
func sum[T ~int | float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}
