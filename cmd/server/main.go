package main

import (
	"context"
	"log"

	_ "github.com/whosonfirst/go-reader-http"

	"github.com/whosonfirst/go-whosonfirst-derivatives/app/server"
)

func main() {

	ctx := context.Background()
	err := server.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run server, %v", err)
	}
}
