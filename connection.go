package campfire

import (
    "encoding/json"
    "net/url"
    "github.com/brettbuddin/httpie"
)

func NewConnection(subdomain, token string) *Connection {
    return &Connection{
        subdomain: subdomain,
        token: token,
        client: httpie.NewClient(httpie.BasicAuth{token, "X"}),
    }
}

type Connection struct {
    subdomain, token string
    client *httpie.Client
}

func (c *Connection) url(path string) *url.URL {
    return &url.URL{
        Scheme: "https",
        Host: c.subdomain + ".campfirenow.com:443",
        Path: path,
    }
}

// Get sends a GET request and loads the result into a value
func (c *Connection) Get(path string, value interface{}) error {
    endpoint  := httpie.Get{c.url(path)}
    resp, err := c.client.Request(endpoint)

    if err != nil {
        return err
    }

    defer resp.Body.Close()
    return json.NewDecoder(resp.Body).Decode(value)
}

// Post sends a POST request and loads the result into a value
func (c *Connection) Post(path string, value interface{}) error {
    body, err := json.Marshal(value)
    if err != nil {
        return err
    }

    endpoint  := httpie.Post{c.url(path), body, "application/json"}
    resp, err := c.client.Request(endpoint)

    if err != nil {
        return err
    }

    defer resp.Body.Close()
    return nil
}

// Put sends a PUT request and loads the result into a value
func (c *Connection) Put(path string, value interface{}) error {
    body, err := json.Marshal(value)
    if err != nil {
        return err
    }

    endpoint  := httpie.Put{c.url(path), body, "application/json"}
    resp, err := c.client.Request(endpoint)

    if err != nil {
        return err
    }

    defer resp.Body.Close()
    return nil
}

// Delete sends a DELETE request and loads the result into a value
func (c *Connection) Delete(path string) error {
    endpoint  := httpie.Delete{c.url(path)}
    resp, err := c.client.Request(endpoint)
    if err != nil {
        return err
    }

    defer resp.Body.Close()
    return nil
}
