package chronicle

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	chronicleapiscope = "https://www.googleapis.com/auth/cloud-platform"
	chroniclebaseurl  = "chronicle.googleapis.com"
)

type Client struct {
	client   *http.Client
	version  string
	project  string
	location string
	instance string
}

func NewClient(ctx context.Context) *Client {
	ts, err := google.DefaultTokenSource(ctx, chronicleapiscope)
	if err != nil {
		return nil
	}
	client := Client{
		client: oauth2.NewClient(ctx, ts),
	}
	return &client
}

func (c *Client) baseUrl() *url.URL {
	u := url.URL{
		Scheme: "https",
		Host:   c.location + "-" + chroniclebaseurl,
		Path:   path.Join("projects", c.project, "locations", c.location, "instances", c.instance),
	}
	return &u
}

func errorResponse(resp *http.Response) error {
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		failmessagebytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("response error %s:\n%s", resp.Status, string(failmessagebytes))
	}
	return nil
}

func (c *Client) ListLogTypes(req ListLogTypesRequest) (*ListLogTypesResult, error) {
	resp, err := c.client.Get(c.baseUrl().String())
	if err != nil {
		return nil, err
	}
	if err = errorResponse(resp); err != nil {
		return nil, err
	}
	respbodybytes, err := io.ReadAll(resp.Body)
	listlogtyperesponse := ListLogTypesResult{}
	err = json.Unmarshal(respbodybytes, &listlogtyperesponse)
	if err != nil {
		return nil, err
	}
	return &listlogtyperesponse, nil
}
