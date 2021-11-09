package main

import (
	"github.com/potproject/goenvgen/gen"
)

func main() {
	err := gen.GenerateFile("test/.env.test", "envgen")
	if err != nil {
		panic(err)
	}
}
