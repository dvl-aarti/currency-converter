package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Currency  struct{

	ID           primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	CurrencyCode string `bson:"currency_code,omitempty" json:"currency_code,omitempty"`
	Currency     string `bson:"currency,omitempty" json:"currency,omitempty"`
	Exchange     float64            `bson:"exchange, omitempty" json:"exchange,omitempty"`

}

type Converter struct {
	Currency string  `bson:"currency,omitempty" json:"currency,omitempty"`
	Amount   float64 `bson:"amount, omitempty" json:"amount,omitempty"`
}