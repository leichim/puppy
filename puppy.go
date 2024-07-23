package puppy

import (
	"fmt"

	"github.com/leichim/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func Version10() {
	fmt.Println("Version 1.0.1")
}
