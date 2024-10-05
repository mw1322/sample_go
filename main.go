package main

import (
	"fmt"
	"log"
	"runners-mysql/config"
	"runners-mysql/server"

	// "runners-mysql/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("Starting Runers App")

	log.Println("Initializig configuration")
	config := config.InitConfig("runners")

	fmt.Println(config.GetString("database.connection_string"))


	log.Println("Initializig database")

	dbHandler := server.InitDatabase(config)
	fmt.Println(dbHandler)

	// log.Println("Initializig HTTP sever")
	// httpServer := server.InitHttpServer(config, dbHandler)

	// httpServer.Start()
}
