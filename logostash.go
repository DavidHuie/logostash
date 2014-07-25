// Implements a Logstash client via TCP or UDP.

package logostash

import (
	"bytes"
	"encoding/json"
	"io"
	"net"
)

type Client struct {
	Conn *net.Conn
}

func NewClient(network, address string) (*Client, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &Client{&conn}, nil
}

// JSON encodes the input struct and sends it to Logstash.
func (c *Client) SendJson(j interface{}) error {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	if err := encoder.Encode(j); err != nil {
		return err
	}
	if _, err := io.Copy(*c.Conn, buf); err != nil {
		return err
	}
	return nil
}
