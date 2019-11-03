package dcc_test

import (
	"fmt"
	"log"
	"testing"

	irc "github.com/thoj/go-ircevent"
	"github.com/voje/golexandria/internal/dcc"
)

const testUser string = "dcctest"
const testChannel string = "#golexandria"
const testServer string = "irc.demo.server:6667"

func TestDCCSend(t *testing.T) {
	irccon := irc.IRC(testUser, testUser)

	irccon.AddCallback("001", func(e *irc.Event) {
		irccon.Join(testChannel)
		irccon.Privmsgf(testChannel, "Testing: %v here!", testUser)

		err := dcc.DCCSend("./testfile.txt", "geryon", irccon)
		if err != nil {
			log.Printf("Error: %v", err)
		}
		// irccon.Privmsg("geryon", "\x01DCC SEND test.txt 2130706433 50200 2\x01")
	})

	err := irccon.Connect(testServer)
	if err != nil {
		fmt.Printf("IRC connection error: %s", err)
	}
	irccon.Loop()
}
