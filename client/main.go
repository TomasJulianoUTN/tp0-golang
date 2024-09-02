package main

import (
	"client/globals"
	"client/utils"
)

func main() {
	//Crea el 'tp0.log'
	utils.ConfigurarLogger()

	//Abre y guarda en la estrutura globals.ClienteConfig el 'config.json'
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")

	//Envio al server el valor de CLAVE
	utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Clave)

	//Con la funcion de 'utils.GenerarYEnviarPaquete' genero y envio el paquete con las lineas de la consola del cliente.
	utils.GenerarYEnviarPaquete()

}
