package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/rbwsam/puller/internal"
)

func main() {
	app := &cli.App{
		Name:  "puller",
		Usage: "pull Docker images to generate load on a Docker registry",
		Action: func(c *cli.Context) error {
			cfg := c.Args().Get(0)

			f, err := os.Open(cfg)
			if err != nil {
				return err
			}
			defer f.Close()

			r := internal.Run{}
			if err := json.NewDecoder(f).Decode(&r); err != nil {
				return err
			}

			return r.Exec()
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
