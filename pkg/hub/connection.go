package hub

import (
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
)

// Connection from websocket
type Connection struct {
	ws       *websocket.Conn
	send     chan []byte
	clientID string
}

// NewConnection from websocket
func NewConnection(ws *websocket.Conn) *Connection {
	return &Connection{
		send:     make(chan []byte, 256),
		ws:       ws,
		clientID: uuid.New().String(),
	}
}

func (c *Connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}
