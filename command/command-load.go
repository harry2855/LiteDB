package command

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"github.com/codecrafters-io/redis-starter-go/storage"
)

// HandleLoad processes the LOAD command
func HandleLoad(c net.Conn) {
	const filePath = "./backup.json"

	// Check if the file exists before trying to load
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		c.Write([]byte("-ERR backup file not found\r\n"))
		return
	}

	// Open the backup.json file
	file, err := os.Open(filePath)
	if err != nil {
		c.Write([]byte("-ERR error opening backup file\r\n"))
		return
	}
	defer file.Close()

	// Unmarshal the contents into the storage.Store
	var storeData map[string]storage.Entry
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&storeData)
	if err != nil {
		c.Write([]byte("-ERR error unmarshalling backup data\r\n"))
		return
	}

	// Update the storage with the loaded data
	storage.Store = storeData

	// Respond with OK
	fmt.Println("Data loaded successfully from backup.json")
	c.Write([]byte("+OK\r\n"))
}
