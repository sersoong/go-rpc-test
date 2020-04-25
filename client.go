package main

import (
	"bufio"
	"log"
	"net/rpc"
	"os"
)

type Reply struct {
	Data string
}

func main() {
	client, err := rpc.Dial("tcp", "192.168.6.240:12345")
	if err != nil {
		log.Fatal(err)
	}

	in := bufio.NewReader(os.Stdin)
	var re Reply
	client.Call("Say", "sersoong", &re)
	log.Println("Say: ", re.Data)
	for {
		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		var reply Reply
		err = client.Call("Listener.GetLine", line, &reply)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Reply: %v,Data :%v", reply, reply.Data)
	}
}
