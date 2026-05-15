package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"golang.ngrok.com/ngrok/v2"
)

const upstream = "http://127.0.0.1:18800"

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	agent, err := ngrok.NewAgent(
		ngrok.WithAuthtoken(os.Getenv("NGROK_AUTHTOKEN")),
	)
	if err != nil {
		return err
	}

	ln, err := agent.Forward(ctx, ngrok.WithUpstream(upstream), ngrok.WithURL(os.Getenv("NGROK_URL")),)
	if err != nil {
		return err
	}

	fmt.Println("Endpoint online:", ln.URL())
	fmt.Println("Forwarding to:", upstream)

	<-ln.Done()
	return nil
}
