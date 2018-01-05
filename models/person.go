package models

type Person struct {
	UserID string		`json:"userID"`
	FirstName string	`json:"firstName"`
	LastName string		`json:"lastName"`
	Age int				`json:"age"`
}
