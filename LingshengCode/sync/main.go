package main

import (
	"context"
	_case "github.com/LingshengCode/sync/case"
	"os"
	"os/signal"
)

func main() {

	_case.OnceCase()
	ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer stop()
	<-ctx.Done()
}
