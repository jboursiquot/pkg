package things

import "fmt"

// Thing is a thing.
type Thing struct {
	Name string
}

func (t Thing) String() string {
	return fmt.Sprintf("Hello, I'm %s", t.Name)
}
