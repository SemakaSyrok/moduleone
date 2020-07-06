package moduleone

import "fmt"

// One ...
const One = "HEllo from module one!"

// Bird ...
type Bird struct {
	name string
}

// Tweet ...
func (b *Bird) Tweet() {
	fmt.Println("Hello! My name is ", b.name)
}

// GetBird ...
func GetBird() Bird {
	return Bird{name: "Durak"}
}

// Run ...
func Run(ch <-chan string) {

	for {
		select {
		case msg := <-ch:
			fmt.Println("Hello My name is", msg)
		}
	}

}
