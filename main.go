package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {

	flag.Parse()
	//extract configuration
	config, _ := ExtractConfiguration()
	fmt.Println("Connecting to database")
	fmt.Println(config)
	dbhandler, _ := NewPersistenceLayer(config.Databasetype, config.DBConnection)
	fmt.Println(dbhandler)
	//RESTful API start
	log.Fatal(ServeAPI(config.ServerAddress, dbhandler))
}
