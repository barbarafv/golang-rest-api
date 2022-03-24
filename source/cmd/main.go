package main

import (
	"aplicacao/source/routes"
)

func main() {
	route := routes.InitRouter()

	route.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
