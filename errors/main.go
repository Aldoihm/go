package main

import (
	"errors"
	"fmt"
	"strconv"
)

var errNotFound = errors.New("not found")
var errSintaxis = errors.New("sintaxis invalid")

// Mapas
var food = map[int]string{
	1: "🍕",
	2: "🍔",
}

func main() {
	// Manejo de errores como normalment ya lo hacía
	fmt.Println("-----------MANEJO DE ERRORES COMO NORMALMENTE LO HACIA (YO LO HAGO MEJOR 😄---------)")
	num, err := strconv.Atoi("1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)

	//Manejo de errores con el paquete errors, y una variable global
	fmt.Println("-----------MANEJO DE ERRORES CON VARIABLE GLOBAL DE TIPO ¡ERRORS!---------)")
	found, err1 := search("a")
	if errors.Is(err1, errNotFound) {
		fmt.Println("Segunda forma de controlar el error")
		fmt.Println("Error controlado: search() ", err1)
		return
	}

	if err1 != nil {
		fmt.Println("Primera forma de controlar el error")
		fmt.Println("search()", err1)
		return
	}
	fmt.Println(found)

}

// Primera forma de controlar el error y es el que más me gusto
func search(key string) (string, error) {
	num, err := strconv.Atoi(key)
	if err != nil {
		return "", fmt.Errorf("strconv.Atoi(key): %v", errSintaxis) //fmt.Errorf retorna un error
	}
	return findFood(num)
}

// Segunda forma de controlar el error, es la más compleja pero más exacta
func findFood(id int) (string, error) {
	value, ok := food[id] //retorna el valor y un booleano, que dice si existe o no esa llave (así trabajan los mapas)
	if !ok {
		return "", fmt.Errorf("food[id]: %w", errNotFound)
	}
	return value, nil
}

//En la parte final de la lección habla de la función errors.Is() y del verbo %w en lugar del verbo %v en la función fmt.Errorf(). Para probar los errores editar la línea 30, que es dónde marcamos el error al momento de buscar en el mapa
//Me gustó más la primera forma, ya que en un solo if, puede devolverte los diferentes tipos de errores y en dónde se produjeron y en la segunda forma no me gusta, porque tu le tienes que decir que tipo de error en específico buscas y si supiera el error de antemano es porque lo estoy provocando.
