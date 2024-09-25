package main

import (
	"fmt"
	"log"
	"padaria/src/app/api/router"
	"padaria/src/app/config"
	"padaria/src/infra/postgres"
)

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
