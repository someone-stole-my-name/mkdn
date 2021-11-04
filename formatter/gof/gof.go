package gof

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
)

type GoFormatter struct {}

func (GoFormatter) Format(cert string) (string, error) {
	decoded, _ := pem.Decode([]byte(cert))
	if decoded == nil {
		return "", errors.New("No PEM data")
	}

	crt, err := x509.ParseCertificate(decoded.Bytes)
	if err != nil {
		return "", err
	}

	return crt.Subject.ToRDNSequence().String(), nil
}
