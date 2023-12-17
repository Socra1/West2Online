package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PasswordDigest string
}

// 加密
func (user *User) Encrypt_password(password string) error {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //前者为加密前的密码，后者为hash加密次数，如果给定的成本小于MinCost，则将成本设置为DefaultCost（10）
	if err != nil {
		return err
	}
	user.PasswordDigest = string(b) //传回加密后的密码
	return nil

}

// 比对原始与加密密码
func (user *User) Compare_password(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
