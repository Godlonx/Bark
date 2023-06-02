package bark

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Client() {
	// on écoute sur le port 8080
	conn, err := net.Dial("tcp", fmt.Sprintf(IP+":"+PORT))
	if err != nil {
		panic(err)
	}
	for {
		// entrée utilisateur
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("client: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		// On envoie le message au serveur
		conn.Write([]byte(text))

		// On écoute tous les messages émis par le serveur et on rajouter un retour à la ligne
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(err)
		}

		// on affiche le message utilisateur
		fmt.Print("serveur : " + message)
	}
}
