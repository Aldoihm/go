package greet //El nombre del paquete es el nombre de la carpeta padre

var emoji = "🧑‍💻" //variable a nivel de paquete y las usa en todas las funciones
/*
	Se puede compartir en paquetes:
	-Variables
	-Constantes
	-Estructuras
	-Funciones

	En los paquetes existen componentes que se pueden exportar y otras que no quiere que sean exportadas
	1.- No exportadas: Desde otro paquete, no se puede acceder al contenido de la variable
	2.- Exportadas: Desde otros paquetes, se puede usar las funciones o contenido de las variables.

	Para que go sepa que importar y que no importar:
	1.- No exportadas: el identificador comienza con minúscula
	2.- Exportadas: el identificador comienza con mayúscula
*/

// English retorna saludo en ingles
func English() string {
	return "Hi " + emoji
}

// Italian retorna saludo en italiano
func Italian() string {
	return "Ciao " + emoji
}
