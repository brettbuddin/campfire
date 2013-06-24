package campfire

import (
    "fmt"
    "strings"
)

type Room struct {
    conn *Connection

    Id               int      `json:"id,omitempty"`
    Full             bool     `json:"full,omitempty"`
    MembershipLimit  int      `json:"membership_limit,omitempty"`
    Name             string   `json:"name,omitempty"`
    OpenToGuests     bool     `json:"open_to_guests,omitempty"`
    Topic            string   `json:"topic,omitempty"`
    Users            []*User  `json:"users,omitempty"`
    UpdatedAt        string   `json:"updated_at,omitempty"`
    CreatedAt        string   `json:"created_at,omitempty"`
}

type RoomResult struct {
    Room *Room `json:"room"`
}

type RoomsResult struct {
    Rooms []*Room `json:"rooms"`
}

// Stream returns a Stream for you to follow the contents of the Room
func (r *Room) Stream(messages chan *Message) *Stream {
    return NewStream(r, messages)
}

// Join joins the Room
func (r *Room) Join() error {
    return r.conn.Post(fmt.Sprintf("/room/%d/join", r.Id), nil)
}

// Leave leaves the Room
func (r *Room) Leave() error {
    return r.conn.Post(fmt.Sprintf("/room/%d/leave", r.Id), nil)
}

// Lock locks the Room
func (r *Room) Lock() error {
    return r.conn.Post(fmt.Sprintf("/room/%d/lock", r.Id), nil)
}

// Unlock unlocks the Room
func (r *Room) Unlock() error {
    return r.conn.Post(fmt.Sprintf("/room/%d/unlock", r.Id), nil)
}

// SendText sends a TextMessage to the Room
func (r *Room) SendText(message string) error {
    return r.message("TextMessage", message)
}

// SendPaste sends a PasteMessage to the Room
func (r *Room) SendPaste(content string) error {
    return r.message("PasteMessage", content)
}

// SendSound sends a SoundMessage to the Room
func (r *Room) SendSound(name string) error {
    return r.message("SoundMessage", name)
}

// SendTweet sends a TweetMessage to the Room
func (r *Room) SendTweet(url string) error {
    return r.message("TweetMessage", url)
}

func (r *Room) message(typ, body string) error {
    result := MessageResult{&Message{}}
    result.Message.Body = strings.Replace(body, "\n", "&#xA;", -1)
    result.Message.Type = typ
    return r.conn.Post(fmt.Sprintf("/room/%d/speak", r.Id), result)
}
