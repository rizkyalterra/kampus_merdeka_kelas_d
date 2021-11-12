package main

import (
	"kelasd/configs"
	"kelasd/routes"
)

func main() {
	configs.CnnectDB()
	e := routes.New()
	e.Start(":8000")
}
