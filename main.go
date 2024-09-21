package main

import (
	"assignment-4/database"
	"assignment-4/router"
	"fmt"
	"os"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	database.StartDB()

	r := router.StartApp()

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
