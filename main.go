/*
Copyright © 2023 musaubrian
*/
package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/musaubrian/turgo/cmd"
	"github.com/musaubrian/turgo/internal/data"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	// ensure that the table is created
	t := data.InitTurgo()
	if err := t.InitTables(); err != nil {
		log.Fatal(err)

	}
	cmd.Execute()
}
