package command

import (
	"fmt"
	"net"
	"strings"

	"LiteDB/config"
)

func HandleConfigGet(input string, c net.Conn) {
	param := extractConfigParam(input)
	switch strings.ToLower(param) {
	case "dir":
		sendConfigResponse(c, "dir", config.RDBFileStoragePath)
	case "dbfilename":
		sendConfigResponse(c, "dbfilename", config.RDBFilename)
	default:
		c.Write([]byte("-ERR unknown parameter\r\n"))
	}
}

func sendConfigResponse(c net.Conn, param, value string) {
	resp := fmt.Sprintf("*2\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n", len(param), param, len(value), value)
	c.Write([]byte(resp))
}

func extractConfigParam(input string) string {
	parts := strings.Split(input, "\r\n")
	if len(parts) > 6 {
		return parts[6]
	}
	return ""
}
