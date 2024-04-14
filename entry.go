package hackerone

import (
	"github.com/sam-bee/security-hackerone-api-client/pkg/client"
	"github.com/sam-bee/security-hackerone-api-client/pkg/hackers"
)

type API struct {
	Hackers *hackers.API
}

func New(username, apiKey string) *API {

	c := client.New(username, apiKey)

	return &API{
		Hackers: hackers.New(c),
	}
}
