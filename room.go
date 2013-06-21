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

func (r *Room) Stream(messages chan *Message) *Stream {
    return NewStream(r, messages)
}

func (r *Room) Join() error {
    return r.conn.Post(fmt.Sprintf("/room/%d/join", r.Id), nil)
}

func (r *Room) Leave() error {
    return r.conn.Post(fmt.Sprintf("/room/%d/leave", r.Id), nil)
}

func (r *Room) Lock() error {
    return r.conn.Post(fmt.Sprintf("/room/%d/lock", r.Id), nil)
}

func (r *Room) Unlock() error {
    return r.conn.Post(fmt.Sprintf("/room/%d/unlock", r.Id), nil)
}

func (r *Room) SendText(message string) error {
    return r.message("TextMessage", message)
}

func (r *Room) SendPaste(content string) error {
    return r.message("PasteMessage", content)
}

func (r *Room) SendSound(name string) error {
    return r.message("SoundMessage", name)
}

func (r *Room) SendTweet(url string) error {
    return r.message("TweetMessage", url)
}

func (r *Room) message(typ, body string) error {
    var result MessageResult;
    result.Message.Body = strings.Replace(body, "\n", "&#xA;", -1)
    result.Message.Type = typ
    return r.conn.Post(fmt.Sprintf("/room/%d/speak", r.Id), result)
}
