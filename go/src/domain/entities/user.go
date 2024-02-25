package entities

type User struct {
	name  string
	age   int
	email string
	// created_at time.Time
	// updated_at time.Time
}

func NewUser(name string, age int, email string) *User {
	return &User{
		name:  set_name(name),
		age:   set_age(age),
		email: set_email(email),
		// created_at: time.Now(),
		// updated_at: time.Now(),
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
