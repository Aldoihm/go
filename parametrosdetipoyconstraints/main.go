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

	fmt.Println("Restricciones a partir de interfaces")
	fmt.Println(sumas(2.5, 4.6, -5.1))
	fmt.Println(sumas(2.5, 4.6, float64(num3)))
}

//Constrains, en español significa restricciones y son para las funciones de any
//las constrains van dentro de los corchetes []. Llamados como parametros de tipo
//Los constraints se pueden usar tanto en la firma de la función como en el cuerpo de la función
//Los constrains reciben un nombre que empieza con mayúscula por convención
/* 3 tipos de constraints
1.- Arbitrario
2.- De unión de elementos
3.- Aproximación de elementos
4.- A partir de interfaces
5.- Importando paquetes externos
	Usan interfaces como la que hicimos en el 4, pero ya declarados sin que nosotros los escribamos.
	1.- go mod init XXXXX // Las XXXX => Es cualquier nombre que desees asignarle por lo general ponen "main"
	2.- Luego: go mod tidy
	3.-Y por último: go get golang.org/x/exp/constraints
	Con eso se instalara el paquete

	NOTA IMPORTANTE: con go mod init main tiene que ser por cada git raiz.
	Si en un git tenemos varios programas de go, al ejecutar go mod init, generaremos problemas con los demás
	programas por esta razon no hice el del paso 5.
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

// Constraints a partir de interfaces
type Numeric interface {
	~int | ~uint | ~float32 | ~float64
}

func sumas[T Numeric](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}
