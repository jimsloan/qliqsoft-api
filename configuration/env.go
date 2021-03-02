package configuration

import (
	"log"
	"os"
)

// Envars ...
type Envars struct {
	AdminEmail string
	Token      string
}

// Getenv ...
func Getenv() Envars {
	var secrets Envars

	// pass secrets via environment
	email, ok := os.LookupEnv("QLIQ_ADMIN_EMAIL")
	if !ok {
		log.Fatal("QLIQ_ADMIN_EMAIL not set\n")
	}
	if len(email) == 0 {
		log.Fatal("QLIQ_ADMIN_EMAIL empty\n")
	}

	token, ok := os.LookupEnv("QLIQ_API_TOKEN")
	if !ok {
		log.Fatal("QLIQ_API_TOKEN not set\n")
	}
	if len(token) == 0 {
		log.Fatal("QLIQ_API_TOKEN empty\n")
	}

	secrets.AdminEmail = email
	secrets.Token = token
	return secrets
}
