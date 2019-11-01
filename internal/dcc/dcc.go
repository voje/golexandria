package dcc

import (
	"fmt"
	"io"

	irc "github.com/thoj/go-ircevent"
)

func DCCSend(file *io.Writer, recipientName string, irccon *irc.Connection) {
	var filename, senderIpInt, senderPort, fileSize string // Todo
	irccon.Privmsg(recipientName, fmt.Sprintf("\x01DCC SEND %v %v %v %v\x01",
		filename, senderIpInt, senderPort, fileSize))
}

func DCCRecv(file io.Reader, address, port int) {

}
