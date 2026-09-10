package erlc

import (
	"context"
	"encoding/json"
	"fmt"
    "net/http"
	"net/url"
	"strings"
	"time"
)

const baseURL = "https://api.erlc.gg"

type Client struct {
	ServerKey string
	http      *http.Client
}

func New(serverKey string) *Client {
	return &Client{ServerKey: serverKey, http: &http.Client{Timeout: 10 * time.Second}}
}

type Include struct {
	Players, Staff, Queue, Vehicles bool
}

func (c *Client) GetServer(ctx context.Context, inc Include) (*ServerInfo, error) {
	q := url.Values{}
	if inc.Players {
		q.Set("Players", "true")
	}
	if inc.Staff {
		q.Set("Staff", "true")
	}
	if inc.Queue {
		q.Set("Queue", "true")
	}
	if inc.Vehicles {
		q.Set("Vehicles", "true")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+"/v2/server?"+q.Encode(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("server-key", c.ServerKey)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erlc api: status %d", resp.StatusCode)
	}

	var info ServerInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	return &info, nil
}

func (c *Client) RunCommand(ctx context.Context, command string) (string, error) {
	body := strings.NewReader(fmt.Sprintf(`{"command":%q}`, command))
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+"/v2/server/command", body)
	if err != nil {
		return "", err
	}
	req.Header.Set("server-key", c.ServerKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var out struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("erlc api: %s", out.Message)
	}
	return out.Message, nil
}