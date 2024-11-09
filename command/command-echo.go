package command

import (
	"fmt"
	"net"
	"strings"
)

func HandleEcho(input string,c net.Conn) {
	parts := strings.Split(input,"\r\n")
	fmt.Println("Echoing",parts[4])
	// c.Write([]byte(parts[4]+"\r\n"))
	message := parts[4] // parts[4] should contain the message (e.g., "hey")

    // RESP bulk string response format: $<number of bytes>\r\n<string>\r\n
    resp := fmt.Sprintf("$%d\r\n%s\r\n", len(message), message)
    c.Write([]byte(resp))
}