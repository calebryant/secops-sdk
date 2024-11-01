package chronicleapi

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	scope = "https://www.googleapis.com/auth/cloud-platform"
)

func AuthenticatedClient() (*http.Client, error) {
	ctx := context.Background()
	ts, err := google.DefaultTokenSource(ctx, scope)
	if err != nil {
		return nil, err
	}
	return oauth2.NewClient(ctx, ts), nil
}
