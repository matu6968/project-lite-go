package discord

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const apiBaseURL = "https://discord.com/api/v10"

type Client struct {
	token string
	http  *http.Client
}

type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Bot           bool   `json:"bot"`
}

type Channel struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    int    `json:"type"`
	GuildID string `json:"guild_id"`
}

type Guild struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewClient(token string) *Client {
	return &Client{
		token: token,
		http: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (c *Client) request(method, endpoint string, body interface{}) (*http.Response, error) {
	req, err := http.NewRequest(method, apiBaseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.token)
	req.Header.Set("Content-Type", "application/json")

	return c.http.Do(req)
}

func (c *Client) ValidateToken(token string) error {
	resp, err := c.request("GET", "/users/@me", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid token: status %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) GetUser() (*User, error) {
	resp, err := c.request("GET", "/users/@me", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// Add more Discord API methods as needed...
