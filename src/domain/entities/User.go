package entities

type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)

type User struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password,omitempty"`
	Role     UserRole `json:"role,omitempty"`
}

func NewUser(name, email, password string, role UserRole) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
	}
}
