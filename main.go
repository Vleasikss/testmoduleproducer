package main

import (
	"fmt"

	"github.com/Vleasikss/testmoduleproducer/sender"
)

func main() {

	sender.SendMessage("hello")
	fmt.Println("hello")

}
