package campfire

import (
    "fmt"
)

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

func (c *Client) Account() (*Account, error) {
    var result AccountResult
    err := c.conn.Get("/account.json", &result)

    if err != nil {
        return nil, err
    }

    result.Account.conn = c.conn
    return result.Account, nil
}

func (c *Client) Me() (*User, error) {
    var result UserResult
    err := c.conn.Get("/users/me.json", &result)

    if err != nil {
        return nil, err
    }

    result.User.conn = c.conn
    return result.User, nil
}

func (c *Client) UserForId(id int) (*User, error) {
    var result UserResult
    err := c.conn.Get(fmt.Sprintf("/users/%d.json", id), &result)

    if err != nil {
        return nil, err
    }

    result.User.conn = c.conn
    return result.User, nil
}

func (c *Client) RoomForId(id int) (*Room, error) {
    var result RoomResult
    err := c.conn.Get(fmt.Sprintf("/room/%d.json", id), &result)

    if err != nil {
        return nil, err
    }

    result.Room.conn = c.conn
    return result.Room, nil
}

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
