package bitclient

import (
	"github.com/dghubble/sling"
)

type BitClient struct {
	sling *sling.Sling
}

func NewBitClient(url string, username string, password string) *BitClient {

	return &BitClient{
		sling: sling.New().Base(url).Set("Accept", "application/json").SetBasicAuth(username, password),
	}
}
