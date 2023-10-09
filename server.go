package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
	"server.go/rutas"
)

func main() {
	// Inicializa la configuración de Firebase
	ctx := context.Background()
	opt := option.WithCredentialsFile(".env")
	config := &firebase.Config{ProjectID: "test-5eebf"}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	// Inicializa el cliente de Firestore
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	// Cerrar el cliente de Firestore cuando ya no se necesite
	defer client.Close()

	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "UP and running...") // imprime la respuesta del cliente
	})

	router.HandleFunc("/usuarios", rutas.GetUser).Methods("GET")
	router.HandleFunc("/usuarios", rutas.AddUser).Methods("POST")
	router.HandleFunc("/usuarios/{ID}", rutas.UpdateUser).Methods("PUT")
	router.HandleFunc("/usuarios/{ID}", rutas.DeleteUser).Methods("DELETE")

	router.HandleFunc("/registrar", func(w http.ResponseWriter, r *http.Request) {
		// Llama a la función para registrar un usuario con correo y contraseña
		rutas.RegisterUser(w, r, app)
	}).Methods("POST")

	log.Println("Server listening on port", port) // imprime en el servidor
	log.Fatal(http.ListenAndServe(port, router))
}
