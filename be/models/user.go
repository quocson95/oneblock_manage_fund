package models

import (
	"errors"

	"gorm.io/gorm"
)

type RoleUser int

const (
	RoleUserAdmin   = 1
	RoleUserManager = 2
	RoluserMemner   = 3
	// RolerUserGuest = 2
)

type User struct {
	gorm.Model
	UserName     string   `json:"user_name,omitempty"`
	Email        string   `json:"email,omitempty"`
	Role         RoleUser `json:"role,omitempty"`
	UsdtInWallet int64    `json:"usdt_in_wallet,omitempty"` // usdt
}

func (u *User) Create() error {
	u.ID = 0
	return DefaulDB.DB().Create(u).Error
}

func (u *User) Read() error {
	if u.ID < 0 {
		return errors.New("invalid id")
	}
	tx := DefaulDB.DB().Model(u).First(u)
	return tx.Error
}

func (u *User) FindByEmail(email string) error {
	if len(u.Email) < 0 {
		return errors.New("invalid email")
	}
	tx := DefaulDB.DB().Model(u).Where("email= ? ", email).First(u)
	return tx.Error
}
