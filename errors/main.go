package main

import (
	"errors"
	"fmt"
	"strconv"
)

var errNotFound = errors.New("not found")

func main() {
	// Manejo de errores como normalment ya lo hacía
	num, err := strconv.Atoi("1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)

	//Manejo de errores con el paquete errors, y una variable global
	found, err1 := search("a")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(found)

}

func search(key string) (string, error) {

	return "", errNotFound
}
