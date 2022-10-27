package main

import "fmt"

func main() {
	//PANIC:nos permite entrar en panico, osea finalizar la ejecución cuando panic es llamada
	//RECOVER: nos puede recuperar después de la generación de un pánico
	division(10, 5)
	division(9, 3)
	division(1, 0)
	division(8, 2)

}
func division(dividendo, divisor int) {
	//aqui se ejecuta el recover y siempre va con un defer para que tenga sentido
	defer func() {
		r := recover() //me devuelve un valor del panic
		if r != nil {  //solo se ejecutara si hay un panic
			fmt.Println("recuperandome del panic", r)
		}
	}()
	validandoDivisor(divisor)
	fmt.Println(dividendo / divisor)
}
func validandoDivisor(divisor int) {
	if divisor == 0 {
		panic("🤦") //Aqui se usa el panic
	}
}
