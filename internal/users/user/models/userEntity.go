package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

const RoleUser = 1

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

func (p *User) IsPasswordValid(hashPassword string) error {
	byteHash := []byte(hashPassword)
	bytePass := []byte(p.Password)
	result := bcrypt.CompareHashAndPassword(byteHash, bytePass)
	if result != nil {
		return errors.New("Login or password is not correct")
	}

	return nil
}

func (p *User) CreatePasswordHash() error  {
	bytePassword := []byte(p.Password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.MinCost)
	if err != nil {
		return err
	}
	password := string(hash)
	p.Password = password

	return nil
}

func (p *User) AddDefaultRole ()  {
	if p.RoleId == 0 {
		p.RoleId = RoleUser
	}
}
