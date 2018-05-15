package str

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

// chained represents the central object in this package, a string type
// with Factory pattern behaviour.
type chained string

// Chainer it's a string factory. It will mount strings through various
// methods, and then outputted in desired form.
type Chainer interface {
	Printer
	Split(string) Splitter
}

// New creates the Chainer using fmt.Sprintf to ensures string
// formatting.
func New(s interface{}, args ...interface{}) (c Chainer) {
	if len(args) > 0 {
		sfmt := New(s).String()
		c = chained(fmt.Sprintf(sfmt, args...))
	} else {
		c = chained(fmt.Sprintf("%v", s))
	}

	return
}

// Split separates string to a Splitter, an array of strings. It uses
// separator received in strings.Split to do the conversion.
func (c chained) Split(sep string) (s Splitter) {
	s = splitted(strings.Split(string(c), sep))
	return
}

// String returns the Chainer as a string.
func (c chained) String() (s string) {
	s = string(c)
	return
}

// Error returns a new error with Chainer as the message.
func (c chained) Error() (err error) {
	err = errors.New(string(c))
	return
}

// Print will log Chainer content to writer received, or os.Stdout as
// default writer.
func (c chained) Print(wa ...io.Writer) (n int, err error) {
	if len(wa) == 0 {
		n, err = fmt.Print(string(c))
	} else {
		n, err = fmt.Fprint(wa[0], string(c))
	}

	return
}
