package services

import (
	"app/pkg/codes"
	"log"

	"github.com/skip2/go-qrcode"
)

func GenerateQR(url string) ([]byte, error) {
	qr, err := qrcode.Encode(url, qrcode.Medium, 256)

	if err != nil {
		log.Printf("%s | %s | %s\n", codes.E007, url, err)
		return nil, err
	}

	return qr, nil
}
