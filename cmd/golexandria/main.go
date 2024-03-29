package main

// Proof of concept irc bot:
//   - connect to server (todo: try different servers)
//   - send msg
//   - receive msg, mux channel, sender, msg content
//   - send file
//   - receive file

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	irc "github.com/thoj/go-ircevent"
	"github.com/voje/golexandria/internal/user"
)

func main() {
	ircnick := user.GetRandomName(0)
	fmt.Printf("Your irc nick is %v.\n", ircnick)

	// Freenode worked.
	// server := "irc.freenode.net:6667"

	// server := "irc.undernet.org:6667"  // nope
	// server := "amsterdam.nl.eu.undernet.org:6667" // worked
	server := "localhost:6667"

	// my local server
	// server := "k-arch.example.com:6667"

	channel := "#golexandria"
	irccon := irc.IRC(ircnick, ircnick)
	irccon.PingFreq = time.Second * 3

	irccon.AddCallback("001", func(e *irc.Event) {
		irccon.Join(channel)
		irccon.Privmsgf(channel, "Hello! Is anybody out there!?")

		// Send CTCP message.
		// irccon.Privmsgf(channel, "CTCP geryon VERSION")
		irccon.Privmsg(channel, "\x01CTCP TIME\x01")
		irccon.Privmsg(channel, "\x01CTCP PING\x01")
		irccon.Privmsg("geryon", "\x01DCC SEND test.txt 2130706433 50200 2\x01")
	})

	/*
		irccon.AddCallback("366", func(e *irc.Event) {
			go func(e *irc.Event) {
				tick := time.NewTicker(5 * time.Second)
				i := 10
				for {
					select {
					case <-tick.C:
						irccon.Privmsgf(channel, "test-%d\n", i)
						if i == 0 {
							fmt.Printf("Timeout while wating for test message from the other thread.")
							return
						}
					}
					i -= 1
				}
			}(e)
		})
	*/

	irccon.AddCallback("PRIVMSG", func(e *irc.Event) {
		// fmt.Printf("%v\n", e)
		sender := e.Arguments[0]
		if sender != channel {
			go handlePrivateMsg(e, sender)
		}

		fmt.Printf("Channel: %s || Nick: %s || Message: %s\n", channel, e.Nick, e.Message())
		/*
			2019/10/19 21:18:51 Connected to amsterdam.nl.eu.undernet.org:6667 (45.58.135.130:6667)
			&{PRIVMSG :kristjan!~kristjan@31.15.194.157 PRIVMSG #golexandria-test :ok this works kristjan 31.15.194.157 kristjan!~kristjan@31.15.194.157 ~kristjan [#golexandria-test ok this works] map[] 0xc0000d0000 context.Background}
		*/
	})

	irccon.AddCallback("PING", func(e *irc.Event) {
		fmt.Printf("Oh my gosh! A PING!!! %v\n", e)
		go irccon.SendRaw("PONG test@test.com")
	})

	irccon.AddCallback("CTCP_VERSION", func(e *irc.Event) {
		fmt.Printf("CTCP_VERSION response: %v\n", e)
	})

	irccon.AddCallback("CTCP_TIME", func(e *irc.Event) {
		fmt.Printf("CTCP_TIME response: %v\n", e)
	})

	irccon.AddCallback("CTCP_PING", func(e *irc.Event) {
		fmt.Printf("CTCP_PING response: %v\n", e)
	})

	irccon.AddCallback("CTCP", func(e *irc.Event) {
		fmt.Printf("CTCP: %+v\n", e)
		fmt.Printf("Host: %+v\n", e.Host)
		fmt.Printf("Message: %+v\n", e.Message())
		msgSpl := strings.Split(e.Message(), " ")
		// If port is 0, that means that server can't open a port and we need to
		// initiate this connection.
		senderPort := msgSpl[4] // if port is 0
		senderAddr := fmt.Sprintf("%v:%v", e.Host, senderPort)
		fmt.Printf("Sender: %v\n", senderAddr)
		tcpConn, err := net.Dial("tcp6", senderAddr)
		if err != nil {
			fmt.Printf("%v\n", err)
			panic(err)
		}
		for {
			message, _ := bufio.NewReader(tcpConn).ReadString('\n')
			fmt.Print("->: " + message)
		}
	})

	err := irccon.Connect(server)
	if err != nil {
		fmt.Printf("Err %s", err)
	}
	irccon.Loop()
}

/*
CTCP Unknown CTCP
CTCP_VERSION Version request (Handled internaly)
CTCP_USERINFO
CTCP_CLIENTINFO
CTCP_TIME
CTCP_PING
CTCP_ACTION (/me)
*/

func handlePrivateMsg(e *irc.Event, sender string) {
	// challes is my nickname
	// nick is sender
	// File transfer message example:
	/*
		Private message :: Channel: goofy_morse || Nick: geryon || Message: SHA-256 checksum for Transferme.txt (remote): 7d9819c4648ddd6dd1192b2ff5294ae340549520f2eb77b97557ab1cb42c58d3
	*/
	fmt.Printf("Private message :: Channel: %s || Nick: %s || Message: %s\n", sender, e.Nick, e.Message())
}

/*
func sendFile(irccon *irc.Connection, user string) {
	// Create a file.
	f1 := []byte("This is a test file\nnr: 913472\n")
	irccon.CTC
}
*/

func sendCTCP(c *irc.Connection, ctcpCommand string) {
	c.SendRawf("%v\n", ctcpCommand)
}
