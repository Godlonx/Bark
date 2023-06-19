package bark

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/submit", formHandler)
	http.ListenAndServe(":8080", nil)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		// Récupérer les valeurs du formulaire
		tag := r.Form.Get("tag")

		// Faites quelque chose avec les données du formulaire
		fmt.Printf("Nom : %s\nEmail : %s\n", tag)

		// Répondre à la demande
		fmt.Fprintf(w, "Formulaire soumis avec succès!")
	} else {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}
