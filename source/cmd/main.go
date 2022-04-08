//go:build !test

package main

import (
	"app/source/repository"
	"app/source/routes"
)

func main() {
	route := routes.InitRouter()
	repository.OpenConnectionDb()

	route.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
