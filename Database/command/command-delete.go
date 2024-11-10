package command

import (
	"fmt"
	"net"
	"strings"
	"LiteDB/storage"
)

// HandleDelete processes the DELETE command in RESP format, which deletes a key-value pair.
func HandleDelete(input string, c net.Conn) {
	// Split the input by \r\n to parse RESP format
	parts := strings.Split(input,"\r\n")
	// c.Write([]byte(parts[4]+"\r\n"))
	key := parts[4] 

	// Check if the key exists in the storage
	if _, exists := storage.Store[key]; exists {
		// Key exists, delete it from storage
		delete(storage.Store, key)
		fmt.Printf("Key %s deleted successfully\n", key)
		c.Write([]byte("+OK\r\n"))
	} else {
		// Key does not exist, return error
		fmt.Printf("Key %s not found for deletion\n", key)
		c.Write([]byte("-ERR key not found\r\n"))
	}
}
