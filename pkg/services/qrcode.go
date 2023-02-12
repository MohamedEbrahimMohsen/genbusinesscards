package services

import (
	"log"

	"github.com/skip2/go-qrcode"
)

func GenerateQR(url string) ([]byte, error) {
	qr, err := qrcode.Encode(url, qrcode.Medium, 256)

	if err != nil {
		log.Printf("Got error while generating qrcode for url %s with error: %s\n", url, err)
		return nil, err
	}

	return qr, nil
}
