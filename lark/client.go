package lark

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	webhookURL string
	httpClient *http.Client
}

func NewClient(webhookURL string, timeout ...time.Duration) *Client {
	d := 10 * time.Second
	if len(timeout) > 0 {
		d = timeout[0]
	}
	return &Client{
		webhookURL: webhookURL,
		httpClient: &http.Client{Timeout: d},
	}
}

func (c *Client) Send(msg Message) error {
	payload := msg.Build()
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("序列化消息失败: %w", err)
	}
	resp, err := c.httpClient.Post(c.webhookURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return errors.New("飞书返回错误: " + resp.Status + " - " + string(respBody))
	}
	return nil
}

func (c *Client) SendText(format string, args ...interface{}) error {
	return c.Send(NewText(format, args...))
}
