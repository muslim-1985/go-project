package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
	IsActive bool   `json:"is_active"`
	Role     string `json:"role"`
	RoleId   int     `json:"role_id"`
	CreatedAt string `json:"created_at"`
}

type Roles struct {
	ID int `json:"id"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
}
