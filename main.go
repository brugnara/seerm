package main

import (
	"log"
	"os"
)

func init() {
	initDB("./database.db")
}

func main() {
	log.Println("Starting seerm")
	listen(os.Getenv("PORT"))
}
