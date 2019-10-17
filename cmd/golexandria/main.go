package main

import (
	"flag"
	"fmt"
	"github.com/thoj/go-ircevent"
)

func main() {
	namePtr := flag.String("name", "DefaultName", "Your name.")
	flag.Parse()

	fmt.Printf("My name is %s.\n", *namePtr)

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
