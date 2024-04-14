package hackers

import (
	"github.com/sam-bee/security-hackerone-api-client/pkg/client"
)

type API struct {
	client *client.Client
}

func New(c *client.Client) *API {
	return &API{
		client: c,
	}
}
