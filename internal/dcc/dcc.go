package dcc

import (
	"fmt"
	"io"
	"net"
	"os"

	irc "github.com/thoj/go-ircevent"
)

func DCCSend(fileName string, recipientName string, irccon *irc.Connection) error {
	fmt.Printf("Sending file: %s to: %v.\n", fileName, recipientName)
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	fmt.Printf("%+v\n", fileInfo)
	return nil
	senderIPInt := 2130706433
	senderPort := 8081
	dccCommand := fmt.Sprintf("\x01DCC SEND %v %v %v %v\x01",
		fileName, senderIPInt, senderPort, fileSize)
	fmt.Printf("Sending DCC: %s\n", dccCommand)
	irccon.Privmsg(recipientName, dccCommand)

	server, err := net.Listen("tcp", "localhost:8081")
	defer server.Close()

	// Wait for a connection
	connection, err := server.Accept()
	if err != nil {
		return err
	}
	fmt.Println("Connection established.")

	// Send the data.
	const bufferSize = 1024
	sendBuffer := make([]byte, bufferSize)
	fmt.Println("Sending file...")
	for {
		_, err := file.Read(sendBuffer)
		if err == io.EOF {
			fmt.Println("Read ended.")
			break
		}
		connection.Write(sendBuffer)
	}
	fmt.Println("File has been sent.")
	return nil
}

func DCCRecv(file io.Reader, address, port int) {

}
