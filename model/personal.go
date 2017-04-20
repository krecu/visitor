package model

type Personal struct {
	Ua string		`json:"ua"`
	FirstName string	`json:"first_name"`
	LastName string		`json:"last_name"`
	Patronymic string	`json:"patronymic"`
	Age string		`json:"age"`
	Gender string		`json:"gender"`
}
