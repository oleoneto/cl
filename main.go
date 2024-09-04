package main

import (
	"embed"

	"cl/cmd/cli"

	_ "github.com/joho/godotenv/autoload"
)

//go:embed data
var data embed.FS

var BuildHash string = ""

func main() {
	cli.Execute(data, BuildHash)
}
