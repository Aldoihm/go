package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("-----FUNCIONES POR VALOR----------------------")
	hello("Aldo Ivan", "Hernandez Marin")

	fmt.Println("\n-----FUNCIONES POR REFERENCIA---------------")
	emoji := "🐕"
	change(&emoji)
	fmt.Printf(emoji)

	fmt.Println("\n-----FUNCIONES CON RETURN-------------------")
	fmt.Println(sum(5, 19))

	fmt.Println("\n--FUNCIONES DE RETORNO DE MULTIPLES VALORES-")
	min, may := convert("Aldo Iván Hernández Marín")
	fmt.Println("minusculas: ", min, "\nMAYUSCULAS: ", may)

	//En Go se controlan los errores, en otros lenguajes hay excepciones, es decir, no se controlan
	fmt.Println("\n-----MANEJO DE ERRORES----------------------")
	content, err := ioutil.ReadFile("./hello.txt") //probar aqui direccionando a un archivo que no exista
	//^^^^Retorna 2 valores
	//1.- El contenido el archivo pero en slice de byte, por lo tanto se tiene que castear
	//2.- Un error
	if err != nil {
		fmt.Println(err)
		return //se coloque para que finalice el programa
	}
	fmt.Println(string(content))

	result, err := div(10, 5)
	if err != nil {
		fmt.Printf("Ocurrió un error: %v \n", err)
		return
	}
	fmt.Println(result)
	fmt.Println("\n-----FUNCIONES QUE RECIBEN FUNCIONES--------")

	nums := []int{1, 4, 70, 5, 67, 90, 2}
	resultado := filter(nums, func(num int) bool {
		if num >= 10 {
			return true
		}
		return false
	})
	fmt.Println(resultado)

	fmt.Println("\n-----FUNCIONES QUE REGRESAN FUNCIONES--------")
	//primera forma
	/* x := hola("Aldo Ivan")() //Con los último parentesis, ejecuta la segunda función
	fmt.Println(x) */
	//segunda forma
	x := hola("Aldo Iván y ")
	fmt.Println(x("María Guadalue")) //con los parentesis de x, manda a ejecutar la funcion de x
	fmt.Println(x("Hans Alarcon"))
	fmt.Println(x("Mario Castillo"))

	fmt.Println("\n-----FUNCIONES VARIATICAS-------------------")
	fmt.Println(suma(10, 5, 15, 20, 2, 3))

	fmt.Println("\n-----FUNCIONES ANONIMAS---------------------")
	y := func() {
		fmt.Println("Hola Sotec Movil")
	}
	y()
	fmt.Println("\n-----FUNCIONES ANONIMAS AUTOEJECUTADAS-----")
	func(texto string) {
		fmt.Println("Hola", texto)
	}("Gophers") //se ejecuta en esta línea

}

//-------------------------AREA DE FUNCIONES------------------------------------

// funciones por valor
func hello(firstName string, lastName string) {
	fmt.Printf("Hello %s %s\n", firstName, lastName)
}

// Función por referencia: recibe un puntero para poder cambiar el valor de la varible origen
func change(value *string) {
	*value = "😋"
}

// Funciones con return, cuando los parametros que recibe son iguales, solo se debe indicar a la última variable
func sum(num1, num2 int) int {
	return num1 + num2
}

// Funciones con retorno de multiples valores
func convert(text string) (string, string) { //(recibe)(retorno_1, retorno_2)
	min := strings.ToLower(text) //texto en minusculas
	may := strings.ToUpper(text) //texto en mayusculas
	return min, may
}

// Manejando errores en las funciones
func div(dividendo, divisor int) (result int, err error) { //colocando nombre a nuestros retornos, cuando se colocan nombres a los retornos, los valores los inicializa en cero según su tipo de dato, en int es cero y en string es vacío- No se recomienda este método
	if divisor == 0 {
		err = errors.New("no puede dividir por 0")
		return
	}
	result = dividendo / divisor
	return
}

// Las funciones pueden recibir funciones y retornar funciones
// Esta función recibe una función
func filter(nums []int, callback func(int) bool) []int {
	result := make([]int, 0, len(nums))
	for _, v := range nums {
		if callback(v) {
			result = append(result, v)
		}
	}
	return result
}

// Funcion que retorna una funcion
func hola(name string) func(string) string {
	return func(ap string) string {
		return "Hello " + name + " " + ap
	}
}

// Funciones variáticas osea que son funciones que reciben un # de argumentos variados
func suma(nums ...int) int { //los 3 puntitos (...) significa que varía el número de argumentos que recibe la funcion y va a generar un slice del tipo entero que se almacenará en nums
	suma := 0
	for _, v := range nums {
		suma += v
	}
	return suma
}
