package rutas

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"server.go/entity"
	"server.go/repository"

	//"github.com/go-chi/chi"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

var (
	repo repository.UserRepository = repository.NewUserRepository()
)

func RegisterUser(resp http.ResponseWriter, req *http.Request, app *firebase.App) error {
	ctx := context.Background()
	client, err := app.Auth(ctx)
	if err != nil {
		return err
	}

	// Parsea los datos del usuario desde la solicitud
	email := req.FormValue("email")
	password := req.FormValue("password")

	// Registra al usuario con correo y contraseña
	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password)

	user, err := client.CreateUser(ctx, params)
	if err != nil {
		return err
	}

	fmt.Fprintf(resp, "Usuario registrado exitosamente. UID: %s\n", user.UID)
	return nil
}

func GetUser(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	users, err := repo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error al obtener los usuarios"}`))
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(users)
}

func AddUser(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var user entity.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}
	user.ID = rand.Int63()
	repo.Save(&user)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(user)
}

// revisar
// func UpdateUser(resp http.ResponseWriter, req *http.Request) {
// 	resp.Header().Set("Content-type", "application/json")

// 	// Obtener el ID del usuario a actualizar desde la ruta o el cuerpo de la solicitud
// 	// Puedes pasar el ID en la ruta o en el cuerpo de la solicitud, dependiendo de tu diseño de API
// 	// Aquí asumimos que se pasa en la ruta
// 	userID := chi.URLParam(req, "ID")

// 	var updatedUser entity.User
// 	err := json.NewDecoder(req.Body).Decode(&updatedUser)
// 	if err != nil {
// 		resp.WriteHeader(http.StatusInternalServerError)
// 		resp.Write([]byte(`{"error": "Error al decodificar la solicitud"}`))
// 		return
// 	}

// 	// Actualizar el usuario en el repositorio
// 	err = repo.Update(userID, &updatedUser)
// 	if err != nil {
// 		resp.WriteHeader(http.StatusInternalServerError)
// 		resp.Write([]byte(`{"error": "Error al actualizar el usuario"}`))
// 		return
// 	}

// 	resp.WriteHeader(http.StatusOK)
// 	json.NewEncoder(resp).Encode(updatedUser)
// }
