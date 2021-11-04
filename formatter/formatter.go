package formatter

import (
	"errors"
	"mkdn/formatter/gof"
)

type Formatter interface {
	Format(cert string) (string, error)
}

func New(format string) (Formatter, error) {
	switch(format) {
	case "go":
		return gof.GoFormatter{}, nil
	default:
		return nil, errors.New("unsuported formatter")
	}
}


