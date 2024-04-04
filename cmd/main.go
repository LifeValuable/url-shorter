package main

import (
	"flag"
	"url-shorter/internal/app"
)

func main() {
	withDB := flag.Bool("d", false, "")
	flag.Parse()
	app.Run(*withDB)
}
