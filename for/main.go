package main

import "fmt"

func main() {
	//En go solo existe una funcion para trabajar con ciclos, y es for, no existe do while. do , while
	// 1.- for clasico
	/* for i := 0; i <= 10; i++ {
		fmt.Println(i)
	} */
	// 2.- for continuo == while en otros lenguajes de programación
	/* j := 0
	for j <= 10 {
		fmt.Println(j)
		j++
	} */
	// 3.- forever == es un for que dura para siempre ctrl+c se cancela la ejecución, se usa para procesos que necesitan ser escuchados permanentemente como sockets u otras conexiones
	/* i := 0
	for {
		fmt.Println(i)
		i++
	} */
	// 4.-  for range == itera sobre slice, mapas y strings
	//4.1.- FOR RANGE SLICE
	fmt.Println("------------FOR RANGE SLICE-----------")
	nums := []uint8{2, 4, 6, 8} // es un slice porque no tiene nada entre los corchetes
	for range nums {
		fmt.Println("EDteam")
	}
	//Para imprimir los valores
	for i, v := range nums { //i= indice, v= valor
		fmt.Printf("Índice : %v, Valor : %v\n", i, v)
	}

	for i := range nums { //como v no la uso, la puedo eliminar
		nums[i] *= 2
	}
	fmt.Println(nums)
	//4.2.- FOR RANGE MAPS
	fmt.Println("\n------------FOR RANGE MAPS-----------")
	sport := map[string]string{"basketball": "🏀", "soccer": "⚽"}
	for key, v := range sport {
		fmt.Println("key:[", key, "] Valor: ", v)
	}

	//4.3.- FOR RANGE STRING
	fmt.Println("\n------------FOR RANGE STRINGS-----------")
	hello := "hello"
	for _, v := range hello {
		fmt.Println(string(v)) //se castea con string, porque si no nos da el valor pero en byte
	}
}
