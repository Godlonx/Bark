package main

import "bark"

var port = ":8080"

type User struct {
	id       int
	pseudo   string
	password string
}

func main() {
	bark.Server()
}
