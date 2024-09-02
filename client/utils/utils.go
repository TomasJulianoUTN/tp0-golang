package utils

import (
	"bufio"
	"bytes"
	"client/globals"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Mensaje struct {
	Mensaje string `json:"mensaje"`
}

type Paquete struct {
	Valores []string `json:"valores"`
}

func IniciarConfiguracion(filePath string) *globals.Config {
	var config *globals.Config
	configFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	if config == nil {
		log.Fatalf("No se pudo cargar la configuración")
	}

	return config
}

func LeerConsola() {
	// Leer de la consola
	reader := bufio.NewReader(os.Stdin)
	log.Println("Ingrese los mensajes")
	text, _ := reader.ReadString('\n')
	log.Print(text)
}

func GenerarYEnviarPaquete() {
	//Generamos el paquete ('lineas' seria el string que incluye todos los mensajes juntos)
	log.Println("Ingrese los mensajes")

	var lineas []string

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		if text == string('\n') {
			break
		}

		text = strings.TrimSpace(text)

		lineas = append(lineas, text)
	}

	// Leemos y cargamos el paquete (cargamos en un paquete los mensajes acumulados)
	paquete := Paquete{Valores: lineas}

	// Enviamos el paquete con las lineas
	log.Printf("Paquete a enviar: %+v", paquete)
	EnviarPaquete(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, paquete)
}

func EnviarMensaje(ip string, puerto int, mensajeTxt string) {
	mensaje := Mensaje{Mensaje: mensajeTxt}
	body, err := json.Marshal(mensaje)
	if err != nil {
		log.Printf("Error codificando mensaje: %s", err.Error())
	}

	url := fmt.Sprintf("http://%s:%d/mensaje", ip, puerto)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error enviando mensaje a ip:%s puerto:%d", ip, puerto)
	}

	log.Printf("Respuesta del servidor: %s", resp.Status)
}

func EnviarPaquete(ip string, puerto int, paquete Paquete) {
	body, err := json.Marshal(paquete)
	if err != nil {
		log.Printf("error codificando mensajes: %s", err.Error())
	}

	url := fmt.Sprintf("http://%s:%d/paquetes", ip, puerto)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error enviando mensajes a ip:%s puerto:%d", ip, puerto)
	}

	log.Printf("Respuesta del servidor: %s", resp.Status)
}

func ConfigurarLogger() {
	logFile, err := os.OpenFile("tp0.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}
