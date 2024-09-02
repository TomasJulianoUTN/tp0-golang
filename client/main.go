package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()

	log.Println("Hola soy un log")
	// loggear "Hola soy un log" usando la biblioteca log
	
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")
	// validar que la config este cargada correctamente

	log.Println(globals.ClientConfig.Mensaje)
	// loggeamos el valor de la config


	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config

	// leer de la consola el mensaje
	//utils.LeerConsola() <---- la default

	utils.LeerConsolaHastaVacio() // <---- luca's version
	

	// generamos un paquete y lo enviamos al servidor
	// utils.GenerarYEnviarPaquete()
}
