package bark

import (
	"fmt"
	"net"
)

const (
	IP   = "127.0.0.1"
	PORT = "8080"
)

func Server() {
	println("Lancement du serveur...")

	// on écoute sur le port 8080
	ln, err := net.Listen("tcp", fmt.Sprintf(IP+":"+PORT))
	if err != nil {
		panic(err)
	}

	// On accepte les connexions entrantes sur le port 8080
	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}

	// Information sur les clients qui se connectent
	fmt.Println("Un client est connecté depuis le port", conn.RemoteAddr())
}
