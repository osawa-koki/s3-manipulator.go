package app

import (
	"fmt"
	"os"
)

func Handler() {
	os.Setenv("FOO", "1")
	fmt.Println("Hello World!!!")
}
