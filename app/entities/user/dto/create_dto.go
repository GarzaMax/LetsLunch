package dto

import (
	"cmd/app/entities/user"
	"time"
)

type CreateUserDto struct {
	Username    string      `json:"username"`
	DisplayName string      `json:"displayName"`
	Birthday    time.Time   `json:"birthday"`
	PhoneNumber string      `json:"phoneNumber"`
	Gender      user.Gender `json:"gender"`
	Password    string      `json:"password"`
}
