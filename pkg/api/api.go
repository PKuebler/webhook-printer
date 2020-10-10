package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkuebler/webhook-printer/pkg/hub"
)

// API Endpoint
type API struct {
	hub *hub.Hub
}

// NewAPI return a endpoint
func NewAPI(hub *hub.Hub) *API {
	return &API{
		hub: hub,
	}
}

// ListenAndServe on specific port
func (a *API) ListenAndServe(port string) error {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		a.serveWs(w, r)
	})

	http.HandleFunc("/hook/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		defer func() {
			r := recover()
			if r != nil {
				var err error
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}
				fmt.Println(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 3 || len(parts[2]) <= 10 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusInternalServerError)
		}
		hookID := parts[2]

		body, _ := ioutil.ReadAll(r.Body)
		a.hub.Push(hookID, body)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, http.StatusText(http.StatusOK))
	})

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	if err := http.ListenAndServe(port, nil); err != nil {
		return err
	}

	return nil
}
