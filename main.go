/*
Copyright © 2023 musaubrian
*/
package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/musaubrian/turgo/cmd"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	cmd.Execute()
}
