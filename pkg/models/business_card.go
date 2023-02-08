package models

type BusinessCard struct {
	QRCode string `json:"qrCode"`
	User   User   `json:"userInfo"`
}
