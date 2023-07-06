package cmd

import "context"

var srv server

func main() {
	ctx := context.Background()
	if err := srv.LoadConfig(ctx); err != nil {
		return
	}
}
