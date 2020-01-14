package config

import "os"

var (
	// PublicURL contains the domain and HTTP schema where the service exists.
	PublicURL string
)

func init() {
	PublicURL = os.Getenv("PUBLIC_URL")
}
