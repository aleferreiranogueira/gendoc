package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aleferreiranogueira/gendoc/document"
	"github.com/urfave/cli"
)

var register = make(document.DocTypes)

func init() {
	register.Set("cpf", document.Cpf{})
}

func main() {
	app := &cli.App{
		Name:    "gendoc",
		Version: "v0.1",
		Usage:   "Generate fake documents for development purposes",
		Action: func(c *cli.Context) error {
			docType := c.Args().First()

			if docType == "" {
				fmt.Printf("Required argument Type missing \n")
				return nil
			}

			doc, err := register.New(docType)

			if err != nil {
				fmt.Printf("Could not generate document for type %v\n", docType)
				return nil

			}

			b, err := json.Marshal(doc)

			if err != nil {
				fmt.Printf("Could not present document")
			}

			fmt.Println(string(b))

			return nil
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
