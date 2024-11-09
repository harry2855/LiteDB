package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
	"sync"
	"github.com/codecrafters-io/redis-starter-go/command"
)

var _ = net.Listen
var _ = os.Exit

var autoSave bool = false
var autoSaveMutex sync.Mutex // To ensure safe access to autoSave across goroutines
var autoSaveSignal = make(chan struct{})

func main() {
	fmt.Println("Logs from your program will appear here!")

	// Start a goroutine that listens for auto-save signals
	go autoSaveRoutine()

	go command.CheckForExpiry()

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Accepted connection", c.RemoteAddr().String())
		go handleConnection(c)
	}
}

func autoSaveRoutine() {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			autoSaveMutex.Lock()
			if autoSave {
				fmt.Println("Executing periodic save...")
				dummyConn := &net.TCPConn{}
				command.HandleSave(dummyConn)
			}
			autoSaveMutex.Unlock()
		case <-autoSaveSignal: // Listen for changes to autoSave
			// Continue to the next iteration to check if autoSave was toggled
		}
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	buf := make([]byte, 1024)

	for {
		_, err := c.Read(buf)
		if err != nil {
			fmt.Println("Error reading: ", err.Error())
			os.Exit(1)
		}
		input := strings.TrimSpace(string(buf))
		fmt.Println("Received:", input)

		// Transaction handling
		if command.IsInTransaction {
			if strings.Contains(input, "EXEC") {
				command.HandleExec(input, c)
			} else if strings.Contains(input, "DISCARD") {
				command.HandleDiscard(input, c)
			} else {
				command.QueueCommand(input)
			}
		} else {
			if strings.Contains(input, "ECHO") {
				command.HandleEcho(input, c)
			} else if strings.Contains(input, "AUTOSAVE-ON") {
				autoSaveMutex.Lock()
				autoSave = true
				autoSaveMutex.Unlock()
				autoSaveSignal <- struct{}{} // Notify the autoSaveRoutine
				c.Write([]byte("+OK\r\n"))
			} else if strings.Contains(input, "AUTOSAVE-OFF") {
				autoSaveMutex.Lock()
				autoSave = false
				autoSaveMutex.Unlock()
				autoSaveSignal <- struct{}{} // Notify the autoSaveRoutine
				c.Write([]byte("+OK\r\n"))
			} else if strings.Contains(input, "CONFIG") {
				command.HandleConfigGet(input, c)
			} else if strings.Contains(input, "SET") {
				command.HandleSet(input, c)
			} else if strings.Contains(input, "GET") {
				command.HandleGet(input, c)
			} else if strings.Contains(input, "PING") {
				c.Write([]byte("+PONG\r\n"))
			} else if strings.Contains(input, "SAVE") {
				command.HandleSave(c)
			} else if strings.Contains(input, "KEYS") {
				command.HandleKeys(input, c)
			} else if strings.Contains(input, "LIST") {
				command.HandleList(c)
			} else if strings.Contains(input, "LOAD") {
				command.HandleLoad(c)
			} else if strings.Contains(input, "DELETE") {
				command.HandleDelete(input, c)
			} else if strings.Contains(input, "MULTI") {
				command.HandleMulti(input, c)
			} else if strings.Contains(input, "EXEC") {
				command.HandleExec(input, c)
			} else if strings.Contains(input, "DISCARD") {
				command.HandleDiscard(input, c)
			}else if strings.Contains(input, "INCR") {
				command.HandleINCR(input, c)
			} else {
				c.Write([]byte("-ERR unknown command\r\n"))
			}
		}
	}
}
