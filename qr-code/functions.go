package qrcode

import (
	"fmt"

	qr "github.com/skip2/go-qrcode"
)

func (qrCode *qrCode) GenerateQr() (qrContent []byte, err error) {
	qrContent, err = qr.Encode(qrCode.Content, qr.Medium, qrCode.Size)
	if err != nil {
		return nil, fmt.Errorf("could not generate a QR Code")
	}
	return qrContent, nil
}
