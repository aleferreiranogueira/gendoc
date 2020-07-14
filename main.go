package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aleferreiranogueira/gendoc/document"
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

			doc := document.CPF{}
			fmt.Printf("There you go, lad: \n %s \n", doc.Generate())

			return nil
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
