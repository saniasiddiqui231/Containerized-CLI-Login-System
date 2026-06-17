package mfa

import (
	"github.com/pquerna/otp/totp"
)

type Setup struct {
	Secret string
	URL    string
}

func Generate(username string) (*Setup, error) {

	key, err := totp.Generate(
		totp.GenerateOpts{
			Issuer:      "CLI Login System",
			AccountName: username,
		},
	)

	if err != nil {
		return nil, err
	}

	return &Setup{
		Secret: key.Secret(),
		URL:    key.URL(),
	}, nil
}

func Verify(code string, secret string) bool {
	return totp.Validate(code, secret)
}
