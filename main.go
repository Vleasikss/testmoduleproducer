package main

import (
	"fmt"

	"github.com/Vleasikss/testmoduleproducer/sender"
)

func main() {

	sender.SendMessage("hello")
	sender.SendMessageNTimes("good", 5)
	fmt.Println("hello")

}
