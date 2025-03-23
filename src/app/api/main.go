package main

import (
	"fmt"
	_ "github.com/swaggo/echo-swagger"
	"log"
	_ "padaria/src/app/api/docs"
	"padaria/src/app/api/router"
	"padaria/src/app/config"
	"padaria/src/infra/postgres"
)

// @title Padaria API
// @version 1.0
// @description This is an example backery server
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /api
func main() {
	setupPostgres()
	serverAddress(config.ServerHost, config.ServerPort)
}

func serverAddress(host string, port int) { //criação de um servidor numa porta específica
	server := router.Start()

	address := fmt.Sprintf("%s:%d", host, port)
	server.Logger.Fatal(server.Start(address))
}

func setupPostgres() {
	err := postgres.SetUpCredentials(config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDBName,
		config.PostgresHost,
		config.PostgresPort,
	) //criar função depois

	if err != nil {
		log.Fatal(err) //para a aplicação
	}
}
