package main

import (
	"log"

	"github.com/jpillora/cloud-torrent/server"
	"github.com/jpillora/opts"
)

var VERSION = "0.1.0-src" //set with ldflags

func main() {
	s := server.Server{
		Title:      "TorrentSaga",
		Port:       3000,
		ConfigPath: "settings.json",
	}

	o := opts.New(&s)
	o.Version(VERSION)
	o.PkgRepo()
	o.LineWidth = 96
	o.Parse()

	if err := s.Run(VERSION); err != nil {
		log.Fatal(err)
	}
}
