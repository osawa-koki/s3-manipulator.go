package app

import (
	"os"
	"fmt"
)

func Handler() {
	os.Setenv("FOO", "1")
	fmt.Println("Hello World!!!")
}
