package main

import (
	"fmt"

	"github.com/Vleasikss/testmoduleproducer/v5/sender"
)

func main() {

	sender.SendMessage("hello")
	sender.SendMessageNTimes("good", 5)
	sender.SayBye()
	fmt.Println("hello")

}
