package command

import (
	"fmt"
	"net"
	"strings"
	"strconv"
	"github.com/codecrafters-io/redis-starter-go/storage"
)

func HandleINCR(input string, c net.Conn) {
	parts := strings.Split(input, "\r\n")
	key := parts[4]

	// Check if the key exists in the storage
	if _, exists := storage.Store[key]; exists {
		// Key exists, increment the value
		entry := storage.Store[key]
		intValue, err := strconv.Atoi(entry.Value)
		if err != nil {
            c.Write([]byte(fmt.Sprintf("-ERR value at '%s' is not an integer\r\n", key)))
            return
        }
		intValue++
		entry.Value = strconv.Itoa(intValue)
        storage.Store[key] = entry
		fmt.Printf("Key %s incremented successfully\n", key)
		c.Write([]byte(fmt.Sprintf(":%d\r\n", intValue)))
	} else {
		// Key does not exist, return error
		fmt.Printf("Key %s not found for increment\n", key)
		c.Write([]byte("-ERR key not found\r\n"))
	}
}