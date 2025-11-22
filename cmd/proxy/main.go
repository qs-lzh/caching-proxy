package main

import (
	"log"

	"github.com/go-chi/chi/v5"

	"github.com/qs-lzh/caching-proxy/internal/cli"
)

var r = chi.NewRouter()

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	if err := cli.CommandExecute(); err != nil {
		log.Fatal(err)
	}
}
