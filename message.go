package campfire

import (
    "fmt"
)

type Message struct {
    conn *Connection

    Id        int    `json:"id,omitempty"`
    Type      string `json:"type"`
    UserId    int    `json:"user_id,omitempty"`
    RoomId    int    `json:"room_id,omitempty"`
    Body      string `json:"body"`
    Starred   bool   `json:"starred,omitempty"`
    CreatedAt string `json:"created_at,omitempty"`
}

type MessageResult struct {
    Message *Message `json:"message"`
}

// Star favorites a message
func (m *Message) Star() error {
    return m.conn.Post(fmt.Sprintf("/messages/%d/star", m.Id), nil)
}

// Unstar unfavorites a message
func (m *Message) Unstar() error {
    return m.conn.Delete(fmt.Sprintf("/messages/%d/unstar", m.Id))
}
