// The LimitReader function in the io package accepts an io.Reader r and a number of bytes n,
// and returns another Reader that reads from r but reports an end-of-file condition after n bytes.
// Implement it.
package ch7

import "io"

type Limited struct {
	reader io.Reader
	limit  int64
	cur    int64
}

func (limited *Limited) Read(p []byte) (int, error) {
	n, err := limited.reader.Read(p)
	if limited.cur + int64(n) > limited.limit {
		return n, io.EOF
	}
	limited.cur += int64(n)
	return n, err
}
func LimitReader(r io.Reader, n int64) *Limited {
	return &Limited{reader: r, limit: n}
}
