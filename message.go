package campfire

import (
    "fmt"
)

type Message struct {
    conn *Connection

    ID        int    `json:"id,omitempty"`
    Type      string `json:"type"`
    UserID    int    `json:"user_id,omitempty"`
    RoomID    int    `json:"room_id,omitempty"`
    Body      string `json:"body"`
    Starred   bool   `json:"starred,omitempty"`
    CreatedAt string `json:"created_at,omitempty"`
}

type MessageResult struct {
    Message *Message `json:"message"`
}

// Star favorites a message
func (m *Message) Star() error {
    return m.conn.Post(fmt.Sprintf("/messages/%d/star", m.ID), nil)
}

// Unstar unfavorites a message
func (m *Message) Unstar() error {
    return m.conn.Delete(fmt.Sprintf("/messages/%d/unstar", m.ID))
}
