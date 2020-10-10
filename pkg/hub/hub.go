package hub

import (
	"context"
	"encoding/json"
	"fmt"
)

// Hub to manage connections
type Hub struct {
	Channels   map[string]map[*Connection]bool
	Register   chan Subscription
	Unregister chan Subscription
	Outbound   chan WSEvent
}

// NewHub to manage connections
func NewHub() *Hub {
	return &Hub{
		Channels:   make(map[string]map[*Connection]bool),
		Register:   make(chan Subscription),
		Unregister: make(chan Subscription),
		Outbound:   make(chan WSEvent),
	}
}

// Push to websocket
func (h *Hub) Push(id string, body []byte) {
	h.Outbound <- WSEvent{
		ChannelID: id,
		Name:      "incoming",
		Event:     string(body),
	}
}

// Run hub
func (h *Hub) Run(ctx context.Context) error {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case evt := <-h.Outbound:
				b, err := json.Marshal(evt)
				if err != nil {
					continue
				}

				connections := h.Channels[evt.ChannelID]
				for c := range connections {
					select {
					case c.send <- b:
					default:
						close(c.send)
						delete(connections, c)
						if len(connections) == 0 {
							delete(h.Channels, evt.ChannelID)
						}
					}
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case s := <-h.Register:
				connections := h.Channels[s.channelID]
				if connections == nil {
					connections = make(map[*Connection]bool)
					h.Channels[s.channelID] = connections
				}
				h.Channels[s.channelID][s.conn] = true

				// welcome message
				evt := struct {
					HookID string `json:"hookID"`
				}{
					s.channelID,
				}
				fmt.Println(evt)
				b, _ := json.Marshal(evt)

				h.Outbound <- WSEvent{
					ChannelID: s.channelID,
					Name:      "config",
					Event:     string(b),
				}
			case s := <-h.Unregister:
				connections := h.Channels[s.channelID]
				if connections != nil {
					if _, ok := connections[s.conn]; ok {
						delete(connections, s.conn)
						close(s.conn.send)
						if len(connections) == 0 {
							delete(h.Channels, s.channelID)
						}
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}
