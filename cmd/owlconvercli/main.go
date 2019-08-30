package main

import (
	"fmt"
	"github.com/shful/gofp"
	"io"

	"github.com/knakk/rdf"
	"github.com/urfave/cli"
	"os"
	"path/filepath"
	"time"
)



func main() {
	app := cli.NewApp()
	app.Name = "owlconvercli"
	app.Version = "2019.5.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Vitaliy Marchenko",
			Email: "vmarchenko@gmail.com",
		},
	}
	app.Copyright = "(c) 2019 VMarchenko"
	app.HelpName = "owlconvercli"
	app.Usage = "Convert different owl format"
	app.UsageText = "owlconvercli - use for convert owl documents in to different formats"

	app.Flags = []cli.Flag{

		cli.StringFlag{
			Name:  "src",
			Value: "pizza-functional.owl",
			Usage: "--src pizza-functional.owl",
		},
		cli.StringFlag{
			Name:  "dst",
			Value: "pizza-functional.xml",
			Usage: "--dst pizza-functional.xml",
		},
	}

	app.Action = func(c *cli.Context) error {

		src := c.String("src")
		fmt.Println(" Source file", src)
		dst := c.String("dst")
		fmt.Println("Destination file", dst)

		srcExt := filepath.Ext(src)

		f, err := os.Open(src)
		if err != nil {
			fmt.Println("Open source file error!", err)
			return nil
		}
		if srcExt == ".owl" {
			dec := rdf.NewTripleDecoder(f, rdf.Turtle)
			for triple, err := dec.Decode(); err != io.EOF; triple, err = dec.Decode() {
				fmt.Println(triple)
			}
			o, err :=gofp.OntologyFromReader(f,filepath.Base(src))
			if err!=nil{
				fmt.Println("Open source file error!", err)
				return nil
			}
			fmt.Println(err, o.About())

		} else if srcExt == ".xml" {

		}

		return nil
	}
	app.Run(os.Args)
}