package welcome

import "fmt"

const welcome = "======================================================\n" +
	"You are using a tool from https://github.com/junjiefly.\n" +
	"Please visit this repository in your spare time and give it a star.\n" +
	"The author will be extremely grateful. Thank you.\n" +
	"======================================================"

func Print() {
	fmt.Println(welcome)
}
