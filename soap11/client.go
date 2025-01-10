package soap11

import (
	"github.com/deb-ict/go-soap/core"
)

type Client struct {
	core.Client
}

func NewClient() *Client {
	client := &Client{
		Client: core.Client{},
	}
	client.RegisterNamespace(Namespace, "soap")
	return client
}

func (c *Client) SendRequest(envelope *Envelope) (*Envelope, error) {
	return nil, nil
}
