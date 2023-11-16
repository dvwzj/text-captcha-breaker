package tcb

import (
	"golang.org/x/net/websocket"
)

type Client struct {
	ws *websocket.Conn
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Connect() error {
	var err error
	c.ws, err = websocket.Dial("wss://docparser-text-captcha-breaker.hf.space/queue/join", "", "https://docparser-text-captcha-breaker.hf.space")
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Close() error {
	return c.ws.Close()
}

func (c *Client) Send(data []byte) error {
	_, err := c.ws.Write(data)
	return err
}

func (c *Client) Receive() ([]byte, error) {
	var msg = make([]byte, 512)
	if n, err := c.ws.Read(msg); err != nil {
		return nil, err
	} else {
		return msg[:n], nil
	}
}
