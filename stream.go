package campfire

import (
    "github.com/brettbuddin/httpie"
    "net/url"
    "fmt"
    "encoding/json"
    "log"
)

type Stream struct {
    base     *httpie.Stream
    room     *Room
    messages chan *Message
}

// NewStream returns a Stream
func NewStream(room *Room, messages chan *Message) *Stream {
    return &Stream{
        room: room,
        messages: messages,
    }
}

// Connect starts the stream
func (s *Stream) Connect() {
    url := &url.URL{
        Scheme: "https",
        Host: "streaming.campfirenow.com",
        Path: fmt.Sprintf("/room/%d/live.json", s.room.Id),
    }

    s.base = httpie.NewStream(
        httpie.Get{url},
        httpie.BasicAuth{s.room.conn.token, "X"},
        httpie.CarriageReturn,
    )

    go s.base.Connect()

    log.Println("Streaming")
    for {
        select {
        case data := <-s.base.Data():
            var m Message
            err := json.Unmarshal(data, &m)

            if err != nil {
                continue
            }

            m.conn = s.room.conn
            s.messages <- &m
        }
    }
}
