package main

import (
	"github.com/Nishith-Savla/wakacli/cmd"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	cmd.Execute()
}
