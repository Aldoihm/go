package main

import "fmt"

func main() {
	fmt.Println(Includes([]string{"a", "b", "c"}, "1"))
	fmt.Println(Includes([]string{"a", "b", "c"}, "a"))
	fmt.Println(Includes([]int{1, 2, 3}, 1))

	fmt.Println(filter([]int{2, 12, 23, 98, 1, 79}, lessThanTwenty))

}

// comparable es una palabra reservada de go, que usa una interfaz para aceptar cualquier tipo de dato y es
// usado para saber si un dato es igual o diferente
func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

// yo defino una interfaz como la que importa de goland
// no lo importo porque si no desmadro todos los demas programas, por eso hago un simulacro
type Ordered interface {
	int | float32 | float64 | ~string
}

// Esta función recibe una función
func filter[T Ordered](nums []T, callback func(T) bool) []T {
	result := make([]T, 0, len(nums))
	for _, v := range nums {
		if callback(v) {
			result = append(result, v)
		}
	}
	return result
}

func lessThanTwenty(num int) bool {
	return num < 20
}
