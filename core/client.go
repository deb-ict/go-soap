package core

import (
	"github.com/deb-ict/go-xml"
)

type Client struct {
	nametable        map[string]string
	typeConstructors map[xml.Name]xml.TypeConstructor
}

func (c *Client) GetNametable() map[string]string {
	if c.nametable == nil {
		c.nametable = make(map[string]string)
	}
	return c.nametable
}

func (c *Client) RegisterNamespace(namespace string, prefix string) {
	if c.nametable == nil {
		c.nametable = make(map[string]string)
	}
	c.nametable[namespace] = prefix
}

func (c *Client) RegisterType(space string, local string, constructor xml.TypeConstructor) {
	if c.typeConstructors == nil {
		c.typeConstructors = make(map[xml.Name]xml.TypeConstructor)
	}
	c.typeConstructors[xml.Name{Space: space, Local: local}] = constructor
}
