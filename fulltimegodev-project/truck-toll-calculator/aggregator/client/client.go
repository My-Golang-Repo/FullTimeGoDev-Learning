package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PorcoGalliard/truck-toll-calculator/types"
	"net/http"
)

type Client struct {
	Endpoint string
}

func NewClient(endpoint string) *Client {
	return &Client{
		Endpoint: endpoint,
	}
}

func (c *Client) AggregateInvoice(distance types.Distance) error {
	b, err := json.Marshal(distance)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", c.Endpoint, bytes.NewReader(b))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("the service responded with non 200 status code %d", resp.StatusCode)
	}
	return nil
}
