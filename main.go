package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// connect to DB
	env := new(Env)
	var err error
	env.DB, err = ConnectDB()
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	router := gin.Default()

	router.Run("8080")

}
