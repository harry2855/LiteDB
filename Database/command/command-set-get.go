package command

import (
	"fmt"
	"net"
	"strings"
	"time"
	"strconv"
	"LiteDB/storage"
)



func HandleSet(input string, c net.Conn){
	parts := strings.Split(input,"\r\n")
	key := parts[4]
	expiryTime := time.Time{}
	fmt.Println("Key",key)
	value := parts[6]
	fmt.Println(input)
	if(strings.Contains(input,"PX")){
		duration,_ := strconv.Atoi(parts[10])
		fmt.Println(duration)
		expiryTime = time.Now().Add(time.Millisecond*time.Duration(duration))
		storage.Store[key] = storage.Entry{Value: value,ExpiryTime: expiryTime,ExpiryTimeExists: true}
	} else{
		storage.Store[key] = storage.Entry{Value: value,ExpiryTime: expiryTime,ExpiryTimeExists: false}
	}
	fmt.Println("Value",value)
	c.Write([]byte("+OK\r\n"))
}

func HandleGet(input string, c net.Conn){
	parts := strings.Split(input,"\r\n")
	key := parts[4]
	fmt.Println("Key",key)
	entry,exists := storage.Store[key]
	if !exists{
		c.Write([]byte("$-1\r\n"))
	} else{
		c.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(entry.Value),entry.Value)))
	}
}

func CheckForExpiry(){
	for{
		time.Sleep(100 * time.Millisecond)
		for key, value := range storage.Store{
			if value.ExpiryTimeExists && value.ExpiryTime.Before(time.Now()){
				delete(storage.Store,key)
			}
		}
	}
}