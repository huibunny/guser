// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// User -.
type User struct {
	Username string `json:"username"  example:"alice"`
	Password string `json:"password"  example:"123456"`
}
