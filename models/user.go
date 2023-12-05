package models

import (
  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  Email string `gorm:"unique" json:"email"`
  Password string `json:"password"`
  Admin bool `json:"-"`
}

type UserRegisterInput struct {
  gorm.Model
  Email string `gorm:"unique" json:"email"`
  Password string `json:"password"`
  AdminSecret string `json:"adminSecret"`
}
