package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "gokyll",
		Usage: "fast & minimal blog-generator",
		Commands: []*cli.Command{
			{
				Name:  "version",
				Usage: "prints the version of gokyll",
				Action: func(context *cli.Context) error {
					fmt.Println("gokyll version 1.0.0")
					return nil
				},
			},
			{
				Name:    "init",
				Usage:   "creates blog in `DIR`",
				Aliases: []string{"generate"},
				Action: func(context *cli.Context) error {
					dir := context.Args().Get(0)
					if dir == "" {
						dir = "./docs"
					}
					fmt.Printf("Initalizing blog in %v\n", dir)
					return nil
				},
			},
			{
				Name:  "build",
				Usage: "build blog",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   "docs",
						Usage:   "set the output directory of your blog",
					},
				},
				Action: func(context *cli.Context) error {
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
