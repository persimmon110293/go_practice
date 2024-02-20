package entities

import "time"

type User struct {
	name       string
	age        int
	email      string
	created_at time.Time
	updated_at time.Time
}

func NewUser(name string, age int, email string) *User {
	return &User{
		name:       name,
		age:        age,
		email:      email,
		created_at: time.Now(),
		updated_at: time.Now(),
	}
}
