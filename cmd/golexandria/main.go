package main

import (
	"fmt"

	irc "github.com/thoj/go-ircevent"
	"github.com/voje/golexandria/internal/namesgenerator"
)

func main() {
	username := namesgenerator.GenRandomName(3)
	fmt.Println("My username is %v.", username)
	return

	server := "irc.undernet.org:6667"

	ircnick := "helmplh"
	irccon := irc.IRC(ircnick, ircnick)

	channel := "#irctest-59323"

	irccon.AddCallback("001", func(e *irc.Event) {
		fmt.Println("Connection succeeded.")
		irccon.Join(channel)
	})
	irccon.AddCallback("366", func(e *irc.Event) {
		fmt.Println("Connection failed.")
	})

	err := irccon.Connect(server)
	if err != nil {
		fmt.Printf("Err %s", err)
		panic(err)
	}
	irccon.Loop()
}
