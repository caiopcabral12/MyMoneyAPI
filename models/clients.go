package models

import (
	v2 "gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name     string  `json:"NAME" validate:"nonzero"`
	Email    string  `json:"email" validate:"nonzero"`
	Password string  `json:"password" validate:"min=4"`
	Balance  float64 `json:"balance"`
}

type CreditCard struct {
	gorm.Model
	ClientID        uint    `json:"clientId"`
	Client          Client  `gorm:"foreignKey:ClientID"`
	CardName        string  `json:"cardName"`
	Value           float64 `json:"value"`
	ExpirationDay   float64 `json:"ExpirationDay"`
	ExpirationMonth float64 `json:"ExpirationMonth"`
}

type Payable struct {
	gorm.Model
	ClientID uint    `json:"clientId"`
	Client   Client  `gorm:"foreignKey:ClientID"`
	Title    string  `json:"tile"`
	Value    float64 `json:"value"`
	PayDay   float64 `json:"payday"`
}

type Receivable struct {
	gorm.Model
	ClientID uint    `json:"clientId"`
	Client   Client  `gorm:"foreignKey:ClientID"`
	Tile     string  `json:"Title"`
	Value    float64 `json:"value"`
	Day      float64 `json:"day"`
}

type Savings struct {
	gorm.Model
	ClientID uint    `json:"clientId"`
	Client   Client  `gorm:"foreignKey:ClientID"`
	Title    string  `json:"title"`
	Value    float64 `json:"value"`
	UpTo     float64 `json:"upto"`
}

func Validate(client *Client) error {
	if err := v2.Validate(client); err != nil {
		return err
	}

	return nil
}
