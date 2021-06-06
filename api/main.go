package main

import (
	"api-server/cmd"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	command := os.Args[1]
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(".env ファイルが見つかりませんでした")
	}
	switch command {
	case "server":
		cmd.Server()
		return
	case "migrate":
		cmd.Migrate()
		return
	default:
		fmt.Println("Cannot find '" + command + "' on apply-api")
	}
}
