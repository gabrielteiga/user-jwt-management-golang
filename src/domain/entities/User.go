package entities

type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)

type User struct {
	name     string
	email    string
	password string
	role     UserRole
}

func NewUser(name, email, password string, role UserRole) *User {
	return &User{
		name:     name,
		email:    email,
		password: password,
		role:     role,
	}
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetRole() UserRole {
	return u.role
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) SetRole(role UserRole) {
	u.role = role
}
