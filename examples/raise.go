package main

import (
	"log"

	"github.com/Nadim147c/go-mpris"
	"github.com/godbus/dbus/v5"
)

func main() {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}

	names, err := mpris.List(conn)

	if err != nil {
		panic(err)
	}

	if len(names) == 0 {
		log.Fatal("No media player found.")
	}

	name := names[0]
	log.Println("Found media player:", name)

	player := mpris.New(conn, name)

	identity, err := player.GetIdentity()

	if err != nil {
		panic(err)
	}

	log.Println("Media player identity:", identity)

	err = player.Raise()
	if err != nil {
		panic(err)
	}
}
