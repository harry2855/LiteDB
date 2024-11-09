package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
	"github.com/codecrafters-io/redis-starter-go/command"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")
	go func() {
		ticker := time.NewTicker(60 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			fmt.Println("Executing periodic save...")
			// Call HandleSave here
			// Passing a dummy connection for simplicity, but adjust as necessary
			dummyConn := &net.TCPConn{}
			command.HandleSave(dummyConn)
		}
	}()

	go command.CheckForExpiry()


	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	// defer l.Close()
	for{
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Accepted connection", c.RemoteAddr().String())
		go handleConnection(c)
	}
	
	
}

func handleConnection(c net.Conn) {
	defer c.Close()
	buf := make([]byte, 1024)
	for{
	_, err := c.Read(buf)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		os.Exit(1)
	}
	input := strings.TrimSpace(string(buf))
	fmt.Println("Received",string(buf))
	if command.IsInTransaction{
		if strings.Contains(input,"EXEC"){
			command.HandleExec(input,c)
		} else if strings.Contains(input,"DISCARD"){
			command.HandleDiscard(input,c)
		} else{
			command.QueueCommand(input)
		}
	} else {
		if strings.Contains(input,"ECHO"){
			command.HandleEcho(input,c)
		}else if strings.Contains(input,"CONFIG"){
			command.HandleConfigGet(input,c)
		}else if strings.Contains(input,"SET"){
			command.HandleSet(input,c)
		} else if strings.Contains(input,"GET"){
			command.HandleGet(input,c)
		}else if strings.Contains(input,"PING"){
			c.Write([]byte("+PONG\r\n"))
		} else if strings.Contains(input,"SAVE"){
			command.HandleSave(c)
		}else if strings.Contains(input,"KEYS"){
			command.HandleKeys(input,c)
		}else if strings.Contains(input,"LIST"){
			command.HandleList(c)
		}else if strings.Contains(input,"LOAD"){
			command.HandleLoad(c)
		} else if strings.Contains(input,"DELETE"){
			command.HandleDelete(input,c)
		}else if strings.Contains(input,"MULTI"){
			command.HandleMulti(input,c)
		}else if strings.Contains(input,"EXEC"){
			command.HandleExec(input,c)
		}else if strings.Contains(input,"DISCARD"){
			command.HandleDiscard(input,c)
		}else{
			c.Write([]byte("-ERR unknown command\r\n"))
		}
	}

}

}
