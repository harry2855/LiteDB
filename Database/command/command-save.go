package command

import(
	"fmt"
	"net"
	"os"
	"encoding/json"
	"LiteDB/storage"
)

func HandleSave(c net.Conn){
	const filePath = "./backup.json"
	file,err := os.Create(filePath)
	if err!=nil{
		fmt.Println("Error creating file")
		c.Write([]byte("-ERR error creating file\r\n"))
		return
	}
	defer file.Close()

	data,err := json.Marshal(storage.Store)
	if err!=nil{
		fmt.Println("Error marshalling data")
		c.Write([]byte("-ERR error marshalling data\r\n"))
		return
	}
	fmt.Println("Data",string(data))

    _,err = file.Write(data)
	if err!=nil{
		fmt.Println("Error writing data to file")
		c.Write([]byte("-ERR error writing data to file\r\n"))
		return
	}

	fmt.Println("Save command")
	c.Write([]byte("+OK\r\n"))
}
