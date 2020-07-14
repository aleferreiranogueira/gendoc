package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "doc",
				Aliases:  []string{"d"},
				Usage:    "Generates a document of type `DOCUMENT`",
				Required: true,
			},
		},
		Name:    "gendoc",
		Version: "v0.1",
		Usage:   "Generate fake documents for development purposes",
		Action: func(c *cli.Context) error {
			fmt.Println(c.String("doc"))
			return nil
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
