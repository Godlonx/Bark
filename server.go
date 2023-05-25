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

	for {
		// On écoute les messages émis par les clients
		buffer := make([]byte, 4096)       // taille maximum du message qui sera envoyé par le client
		length, err := conn.Read(buffer)   // lire le message envoyé par client
		message := string(buffer[:length]) // supprimer les bits qui servent à rien et convertir les bytes en string

		if err != nil {
			fmt.Println("Le client s'est déconnecté")
			break
		}

		if message == "exit" {
			break
		}

		// on affiche le message du client en le convertissant de byte à string

		fmt.Print("Client:", message)

		// On envoie le message au client pour qu'il l'affiche
		conn.Write([]byte(message + "\n"))
	}
}
