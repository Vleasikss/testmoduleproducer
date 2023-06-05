package sender

import "fmt"

func SendMessage(msg string) {
	fmt.Println("sending message...")
	fmt.Println(msg)
}

func SendMessageNTimes(msg string, n int) {
	fmt.Println("BEGIN!")
	for i := 1; i <= n; i++ {
		fmt.Printf("%d: %s\n", i, msg)
	}
	fmt.Println("END!")
}

func SayHello() {
	fmt.Println("hello")
}
