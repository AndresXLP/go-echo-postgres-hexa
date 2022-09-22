package models

type User struct {
	ID       uint   `json:"id,omitempty"`
	Name     string `json:"name,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}
