package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	chromedptool "github.com/yiGmMk/go-pdf/chromedp"
)

const (
	defaultUrl  = "https://baidu.com"
	defaultPath = "./baidu.pdf"
)

var (
	// cli app params
	url  string
	path string
	pU   = &cli.StringFlag{
		Name:        "url",
		Aliases:     []string{"u"},
		Usage:       fmt.Sprintf("url to generate pdf,if no url in input,we use as %s default", defaultUrl),
		Destination: &url,
	}
	pP = &cli.StringFlag{
		Name:        "path",
		Aliases:     []string{"p"},
		Usage:       fmt.Sprintf("output pdf file path,if no path in input,we use as %s default", defaultPath),
		Destination: &path,
	}
)

// gen pdf from url
func GenPdf(url, path string) error {
	err := chromedptool.ChromedpHandler{}.GenPdf(url, path)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	app := &cli.App{
		Name:  "go-pdf-cli",
		Usage: "generate pdf from url",
		Action: func(c *cli.Context) error {
			log.Printf("input,url: %s, path: %s", url, path)
			if url == "" {
				url = defaultUrl
				log.Printf("no url in input,use as %s default", defaultUrl)
			}
			if path == "" {
				path = defaultPath
				log.Printf("no path in input,use as %s default", defaultPath)
			}
			return GenPdf(url, path)
		},
		Flags: []cli.Flag{pU, pP},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
