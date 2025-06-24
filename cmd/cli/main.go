package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func run() {
	cmd := &cli.Command{
		Name:  "Verbesserer",
		Usage: "makes it easier to make incremental improvements to your codebase!",
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func main() {
	run()
}
