// *File.Close() is a tricky function call; it is useful to avoid the temptation of deferring this call
package ch5

import (
	"io"
	"os"
)

func CloseExperiment(from, to string) (int64, error) {
	in, e := os.Open(from)
	if e != nil {
		return 0, e
	}
	defer in.Close() // note: there seems to be no easy way to take care of errors from a "deferred" function call
	// here, in seems to be usable, now we can reuse the identifier e
	out, e := os.Create(to)
	if e != nil {
		return 0, e
	}
	written, e := io.Copy(out, in)
	if closeError := out.Close(); e == nil { // only if copying succeeded
		e = closeError // this is the error that we want to return only in case Copy succeeds and Close fails
	}
	// this is the pattern to follow instead of defer out.Close()
	return written, e
}
