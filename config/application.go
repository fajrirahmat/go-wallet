package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

//applicationConfig struct to store application configuration
type applicationConfig struct {
	Server      serverConfig
	Certificate certConfig
}

type serverConfig struct {
	Host string
	Port int
}

type certConfig struct {
	ServerCertificatePath string
	ServerPrivateKeyPath  string
	CACertificatePath     string
}

//App variable of applicationConfig to store application-related configuration
var App *applicationConfig

//Init initialite application config
func Init() {
	log.Printf("Load configuration...")
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error() + ", Reading from environtment variable")
	}
	App = &applicationConfig{}
	App.Server.Host = os.Getenv("SERVER_HOST")
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Println("Failed to convert port value from configuration, set to default value")
		port = 8080
	}
	App.Server.Port = port
	App.Certificate.CACertificatePath = os.Getenv("CA_CERT_PATH")
	App.Certificate.ServerCertificatePath = os.Getenv("SERVER_CERT")
	App.Certificate.ServerPrivateKeyPath = os.Getenv("SERVER_KEY")
}
