/*
Copyright © 2023 musaubrian
*/
package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/musaubrian/turgo/cmd"
	"github.com/musaubrian/turgo/internal/util"
)

func main() {
	loadEnv()
	cmd.Execute()
}

func loadEnv() {
	envLoc, err := util.GetEnvLoc()
	if err != nil {
		log.Fatal(err)
	}
	if err := godotenv.Load(envLoc); err != nil {
		// Create the .env
		if err := util.CreateEnv(envLoc); err != nil {
			log.Fatal(err)
		}
		/* // Reload the .env
		if err := godotenv.Load(envLoc); err != nil {
			log.Fatal(err)
		} */
	}
}
