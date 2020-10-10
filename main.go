package main

import (
	"context"

	"github.com/pkuebler/webhook-printer/pkg/api"
	"github.com/pkuebler/webhook-printer/pkg/hub"
)

func main() {
	ctx := context.Background()

	hubService := hub.NewHub()
	hubService.Run(ctx)

	apiService := api.NewAPI(hubService)
	apiService.ListenAndServe(":8080")
}
