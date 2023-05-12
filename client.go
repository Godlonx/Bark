package bark

import (
	"fmt"
	"net"
)

func Client() {
	// on écoute sur le port 8080
	_, err := net.Dial("tcp", fmt.Sprintf(IP+":"+PORT))
	if err != nil {
		panic(err)
	}
}
