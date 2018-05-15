package str

import (
	"fmt"
	"io"
)

// Printer represents a string factory with the hability to serve it on
// differents outputs.
type Printer interface {
	fmt.Stringer
	Error() error
	Print(...io.Writer) (int, error)
}
