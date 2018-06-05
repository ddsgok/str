package str_test

// testWriter serves as a io.Writer mock. It will stores the content
// asked to Write on a buf value, that can be consulted.
type testWriter struct {
	buf []byte
}

// String return the buf value as string.
func (tw *testWriter) String() (s string) {
	s = string(tw.buf)
	return
}

// Clean buf value.
func (tw *testWriter) Clean() {
	tw.buf = []byte{}
}

// Write implements the io.Writer interface, stores the value received
// into a buf value.
func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.buf = append(tw.buf, p...)
	n, err = len(tw.buf), nil
	return
}

// newTestWriter builds the testWriter object, returnint to be of use
// in tests with all functions.
func newTestWriter() (tw *testWriter) {
	tw = &testWriter{
		buf: []byte{},
	}
	return
}
