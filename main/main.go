package main

import (
	"bark"

	_ "github.com/mattn/go-sqlite3"
)

var port = ":8080"

func main() {
	bark.Server()
}
