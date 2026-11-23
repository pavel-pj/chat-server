package main

import (
	"context"
	"log"

	a "github.com/pavel-pj/go-chat/internal/app"
)

func main() {

	ctx := context.Background()

	app, err := a.NewApp(ctx)
	if err != nil {
		log.Fatalf("Fail to init app:%v", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("Fail to start app:%v", err)
	}

}
