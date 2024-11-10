// command/command-keys.go
package command

import (
	"fmt"
	"net"
	"strings"
	"LiteDB/storage"
)
// HandleKeys processes the KEYS command, finding all keys matching the pattern.
func HandleKeys(input string, c net.Conn) {
	// Split input to get the pattern part (expects pattern at index 4, following RESP protocol)
	parts := strings.Split(input, "\r\n")
	if len(parts) < 5 {
		c.Write([]byte("-ERR wrong number of arguments for 'keys' command\r\n"))
		return
	}
	pattern := parts[4]
	
	// Find matching keys
	matchingKeys := matchKeys(pattern)
	
	// Formulate RESP array response for matching keys
	resp := fmt.Sprintf("*%d\r\n", len(matchingKeys))
	for _, key := range matchingKeys {
		resp += fmt.Sprintf("$%d\r\n%s\r\n", len(key), key)
	}
	c.Write([]byte(resp))
}

// matchKeys finds all keys in the store matching the pattern (supports '*' wildcard).
func matchKeys(pattern string) []string {
	var keys []string

	// Convert wildcard '*' to an equivalent search pattern
	pattern = strings.ReplaceAll(pattern, "*", ".*")

	// Loop over the store to find matching keys
	for key := range storage.Store {
		if matched := strings.Contains(key, strings.Trim(pattern, ".*")); matched {
			keys = append(keys, key)
		}
	}
	return keys
}
