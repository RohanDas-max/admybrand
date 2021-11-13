package model

import "time"

type Data struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	DOB         string     `json:"dob"`
	Address     string     `json:"address"`
	Description string     `json:"desciption"`
	CreatedAt   *time.Time `json:"createdat"`
	UpdatedAt   *time.Time `json:"updatedat"`
}

type UserID struct {
	ID int `json:"id"`
}
