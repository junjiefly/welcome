package github.com/junjiefly/welcome

import "fmt"

const welcome = "======================================================" +
	"you are using a tool from https://github.com/junjiefly, " +
	"please visit this repository in your spare time and give it a star. " +
	"The author will be extremely grateful. Thank you."

func Welcome() {
	fmt.Println(welcome)
}
