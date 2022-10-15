package main

import (
	"final-assignment/configs"
	"final-assignment/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs.LoadEnv()

	routes.LoadRoute()
}
