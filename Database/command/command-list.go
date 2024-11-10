// command/command-list.go
package command

import (
	"fmt"
	"net"
	"strings"
	"time"
	"LiteDB/storage"
)

// HandleList processes the LIST command
func HandleList(c net.Conn) {
	// Header of the table
	header := fmt.Sprintf("%-20s %-20s %-30s %-15s\n", "Key", "Value", "Expiry Time", "Expiry Exists")
	table := header + strings.Repeat("-", len(header)) + "\n"

	// Iterate over each entry and add to the table string
	for key, entry := range storage.Store {
		// Determine expiry information
		expiryInfo := "No expiry set"
		if entry.ExpiryTimeExists {
			if entry.ExpiryTime.Before(time.Now()) {
				expiryInfo = fmt.Sprintf("Expired at %s", entry.ExpiryTime.Format(time.RFC3339))
			} else {
				expiryInfo = fmt.Sprintf("Expires at %s", entry.ExpiryTime.Format(time.RFC3339))
			}
		}

		// Format each row with aligned columns
		row := fmt.Sprintf("%-20s %-20s %-30s %-15t\n", key, entry.Value, expiryInfo, entry.ExpiryTimeExists)
		table += row
	}

	// Send the table as a bulk string response to the client
	clientResponse := fmt.Sprintf("$%d\r\n%s\r\n", len(table), table)
	c.Write([]byte(clientResponse))
}
