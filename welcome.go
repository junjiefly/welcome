package github.com/junjiefly/welcome

import "fmt"

const welcome = "======================================================" +
	"You are using a tool from https://github.com/junjiefly. " +
	"Please visit this repository in your spare time and give it a star. " +
	"The author will be extremely grateful. Thank you." +
	"======================================================"

func Print() {
	fmt.Println(welcome)
}
