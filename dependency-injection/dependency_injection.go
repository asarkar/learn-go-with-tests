package dependencyinjection

import (
	"fmt"
	"io"
)

// Greet sends a personalised greeting to writer.
func Greet(writer io.Writer, name string) error {
	_, err := fmt.Fprintf(writer, "Hello, %s", name)
	return err
}
