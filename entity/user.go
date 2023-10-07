package entity

type User struct {
	ID              int64  `json:"ID"`
	Nombre          string `json:"Nombre"`
	Apellido        string `json:"Apellido"`
	SegundoApellido string `json:"SegundoApellido"`
	Email           string `json:"Email"`
	Rut             string `json:"Rut"`
	Fono            string `json:"Fono"`
}
