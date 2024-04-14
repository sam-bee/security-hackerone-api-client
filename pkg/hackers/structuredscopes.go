package hackers

import (
	"context"
	"fmt"

	"github.com/sam-bee/security-hackerone-api-client/pkg/api"
)

type getStructuredScopesResponse struct {
	Data  []api.StructuredScope `json:"data"`
	Links api.Links             `json:"links"`
}

// GetStructuredScopes returns a list of structured scopes for a given program. If there are further pages, nextPage will be >0.
func (a *API) GetStructuredScopes(ctx context.Context, handle string, pageOptions *api.PageOptions) (programs []api.StructuredScope, nextPage int, err error) {
	var response getStructuredScopesResponse
	path := fmt.Sprintf(
		"/hackers/programs/%s/structured_scopes?page[number]=%d&page[size]=%d",
		handle,
		pageOptions.GetPageNumber(),
		pageOptions.GetPageSize(),
	)
	if err := a.client.Get(ctx, path, &response); err != nil {
		return nil, 0, err
	}
	if response.Links.Next != "" {
		nextPage = pageOptions.GetPageNumber() + 1
	}
	return response.Data, nextPage, nil
}
