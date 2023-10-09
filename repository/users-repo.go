package repository

import (
	"context"
	"log"
	"strconv"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"server.go/entity"
)

type UserRepository interface {
	Save(user *entity.User) (*entity.User, error)
	FindAll() ([]entity.User, error)
	Update(userID int64, user *entity.User) error // Nueva función para actualizar usuario
	Delete(userID int64) error
}

type repo struct{}

// NewUserRepository
func NewUserRepository() UserRepository {
	return &repo{}
}

const (
	projectId      string = "test-5eebf"
	collectionName string = "usuarios"
)

func (*repo) Save(user *entity.User) (*entity.User, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatal("Error al crear un cliente de firestore: ", err)
		return nil, err
	}

	defer client.Close() // cierra el cliente de firestore una vez que la funcion termina

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":              user.ID,
		"Nombre":          user.Nombre,
		"Apellido":        user.Apellido,
		"SegundoApellido": user.SegundoApellido,
		"Email":           user.Email,
		"Rut":             user.Rut,
		"Fono":            user.Fono,
	})

	if err != nil {
		log.Fatal("Error al agregar un nuevo usuario: ", err)
		return nil, err
	}
	return user, nil
}

func (*repo) FindAll() ([]entity.User, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatal("Error al crear un cliente de firestore: ", err)
		return nil, err
	}

	defer client.Close() // cierra el cliente de firestore una vez que la funcion termina
	var users []entity.User

	it := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := it.Next()
		// si termino de listar rompe el for
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal("No se pudo iterar la lista de usuarios: ", err)
			return nil, err
		}
		user := entity.User{
			ID:              doc.Data()["ID"].(int64),
			Nombre:          doc.Data()["Nombre"].(string),
			Apellido:        doc.Data()["Apellido"].(string),
			SegundoApellido: doc.Data()["SegundoApellido"].(string),
			Email:           doc.Data()["Email"].(string),
			Rut:             doc.Data()["Rut"].(string),
			Fono:            doc.Data()["Fono"].(string),
		}
		users = append(users, user)
	}
	return users, nil
}

func (*repo) Update(userID int64, updatedUser *entity.User) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatal("Error al crear un cliente de firestore: ", err)
		return err
	}
	ID := strconv.FormatInt(userID, 10)
	defer client.Close()

	_, err = client.Collection(collectionName).Doc(ID).Set(ctx, updatedUser)
	if err != nil {
		log.Fatal("Error al actualizar el usuario: ", err)
		return err
	}

	return nil
}

func (*repo) Delete(userID int64) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatal("Error al crear un cliente de firestore: ", err)
		return err
	}
	ID := strconv.FormatInt(userID, 10)
	defer client.Close()

	_, err = client.Collection(collectionName).Doc(ID).Delete(ctx)
	if err != nil {
		log.Fatal("Error al actualizar el usuario: ", err)
		return err
	}

	return nil
}
