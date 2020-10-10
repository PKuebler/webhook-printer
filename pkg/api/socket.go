package api

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"github.com/pkuebler/webhook-printer/pkg/hub"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (a *API) serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Upgrade Websocket")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	c := hub.NewConnection(ws)
	s := hub.NewSubscription(a.hub, c, uuid.New().String())

	a.hub.Register <- s
	go s.WritePump()
	go s.ReadPump()
}
