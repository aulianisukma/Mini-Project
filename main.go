package main

import (
	"Mini-Project/config"
	"Mini-Project/routes"
)

func main() {
	config.InitDB()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}
