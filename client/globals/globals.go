package globals

type Config struct {
	Ip      string `json:"ip"`
	Puerto  int    `json:"puerto"`
	Mensaje string `json:"mensaje"`
	Clave   string `json:"clave"`
}

var ClientConfig *Config
