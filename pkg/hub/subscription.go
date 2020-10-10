package hub

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// Subscription between connection and dashboard
type Subscription struct {
	hub       *Hub
	conn      *Connection
	channelID string
}

// NewSubscription with hub and channelID
func NewSubscription(hub *Hub, conn *Connection, channelID string) Subscription {
	return Subscription{
		hub:       hub,
		conn:      conn,
		channelID: channelID,
	}
}

// ReadPump message
func (s Subscription) ReadPump() {
	c := s.conn
	defer func() {
		s.hub.Unregister <- s
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}

		var cmd WSCommand
		if err := json.Unmarshal(msg, &cmd); err != nil {
			fmt.Println(err)
			continue
		}
	}
}

// WritePump message
func (s *Subscription) WritePump() {
	c := s.conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
