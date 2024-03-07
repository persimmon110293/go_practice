package entities

import "time"

type User struct {
	Name       string    `json:"name"`
	Age        int       `json:"age"`
	Email      string    `json:"email"`
	created_at time.Time `json:"created_at"`
	updated_at time.Time `json:"updated_at"`
}

func NewUser(name string, age int, email string) *User {
	return &User{
		Name:  set_name(name),
		Age:   set_age(age),
		Email: set_email(email),
	}
}

func set_name(name string) string {
	return name
}

func set_age(age int) int {
	return age
}

func set_email(email string) string {
	return email
}
