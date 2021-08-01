package main

import (
	"fmt"

	"github.com/Ben-harder/estate/server"
)

func main() {
	fmt.Println("Welcome to the Estate")
	svr := server.NewServer(8080)
	svr.ListenAndServe()
}
