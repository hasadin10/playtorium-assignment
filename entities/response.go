package entities

import ()

type Response struct {
	Status          string        `json:"status"`
	Message         string        `json:"message,omitempty"`
	ErrorCode       string        `json:"errorCode,omitempty"`
	ErrorMessage    []string       `json:"errorMessage,omitempty"`
	TransactionCode string        `json:"transactionCode,omitempty"`
	TotalPrice     float64         `json:"totalPrice,omitempty"`
}

