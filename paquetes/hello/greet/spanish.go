// Paquete es una colección de archivos
package greet //El nombre del paquete es el nombre de la carpeta padre

func Spanish() string {
	return "Hola " + emoji // emoji se encuenta en greet.go y se puede usar porque spanish.go y greet.go, pertenecen al mismo paquete
}
