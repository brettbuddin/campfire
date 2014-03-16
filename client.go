package campfire

import (
    "fmt"
)

// NewCLient returns a Client
func NewClient(subdomain, token string) *Client {
    return &Client{
        conn: NewConnection(subdomain, token),
        subdomain: subdomain,
        token: token,
    }
}

type Client struct {
    conn *Connection
    subdomain, token string
}

// Account returns the Account related to the token
func (c *Client) Account() (*Account, error) {
    var result AccountResult
    err := c.conn.Get("/account.json", &result)

    if err != nil {
        return nil, err
    }

    result.Account.conn = c.conn
    return result.Account, nil
}

// Me returns a User that represents You
func (c *Client) Me() (*User, error) {
    var result UserResult
    err := c.conn.Get("/users/me.json", &result)

    if err != nil {
        return nil, err
    }

    result.User.conn = c.conn
    return result.User, nil
}

// UserForID returns a User that has the specific ID
func (c *Client) UserForID(id int) (*User, error) {
    var result UserResult
    err := c.conn.Get(fmt.Sprintf("/users/%d.json", id), &result)

    if err != nil {
        return nil, err
    }

    result.User.conn = c.conn
    return result.User, nil
}

// RoomForID returns a Room that has the specific ID
func (c *Client) RoomForID(id int) (*Room, error) {
    var result RoomResult
    err := c.conn.Get(fmt.Sprintf("/room/%d.json", id), &result)

    if err != nil {
        return nil, err
    }

    result.Room.conn = c.conn
    return result.Room, nil
}

// Rooms returns all Rooms listed on the Account
func (c *Client) Rooms() ([]*Room, error) {
    var result RoomsResult
    err := c.conn.Get("/rooms.json", &result)

    if err != nil {
        return nil, err
    }

    rooms := make([]*Room, len(result.Rooms))
    for i, room := range result.Rooms {
        rooms[i] = room
        rooms[i].conn = c.conn
    }

    return rooms, nil
}
