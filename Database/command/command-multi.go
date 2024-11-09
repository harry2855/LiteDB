package command

import (
	"net"
	"strings"
)

var IsInTransaction bool
var transactionQueue []string

// HandleMulti processes the MULTI command (start of transaction).
func HandleMulti(input string, c net.Conn) {
	// Ensure MULTI command is valid and can only be executed once per transaction.
	if IsInTransaction {
		c.Write([]byte("-ERR MULTI can only be called once per transaction\r\n"))
		return
	}

	// Begin the transaction and clear the queue.
	IsInTransaction = true
	transactionQueue = []string{}

	// Respond with OK
	c.Write([]byte("+OK\r\n"))
}

// HandleExec processes the EXEC command (commit the transaction).
func HandleExec(input string, c net.Conn) {
	// Check if we are in a transaction.
	if !IsInTransaction {
		c.Write([]byte("-ERR EXEC without MULTI\r\n"))
		return
	}

	// Execute all commands in the transaction queue.
	for _, cmd := range transactionQueue {
		// Parse and execute the queued commands.
		if strings.Contains(cmd,"ECHO"){
			HandleEcho(input,c)
		}else if strings.Contains(cmd,"CONFIG"){
			HandleConfigGet(input,c)
		}else if strings.Contains(cmd,"SET"){
			HandleSet(input,c)
		} else if strings.Contains(cmd,"GET"){
			HandleGet(input,c)
		}else if strings.Contains(cmd,"PING"){
			c.Write([]byte("+PONG\r\n"))
		} else if strings.Contains(cmd,"SAVE"){
			HandleSave(c)
		}else if strings.Contains(cmd,"KEYS"){
			HandleKeys(input,c)
		}else if strings.Contains(cmd,"LIST"){
			HandleList(c)
		}else if strings.Contains(cmd,"LOAD"){
			HandleLoad(c)
		} else if strings.Contains(input,"DELETE"){
			HandleDelete(input,c)
		} else{
			c.Write([]byte("-ERR unknown command\r\n"))
		}
	}

	// Clear the transaction queue and reset the transaction state.
	IsInTransaction = false
	transactionQueue = nil

	// Respond with OK
	c.Write([]byte("+OK\r\n"))
}

// HandleDiscard processes the DISCARD command (abort the transaction).
func HandleDiscard(input string, c net.Conn) {
	// Check if we are in a transaction.
	if !IsInTransaction {
		c.Write([]byte("-ERR DISCARD without MULTI\r\n"))
		return
	}

	// Discard all queued commands.
	IsInTransaction = false
	transactionQueue = nil

	// Respond with OK
	c.Write([]byte("+OK\r\n"))
}

// QueueCommand adds commands to the transaction queue if we're in a transaction.
func QueueCommand(input string) {
	if IsInTransaction {
		transactionQueue = append(transactionQueue, input)
	}
}
