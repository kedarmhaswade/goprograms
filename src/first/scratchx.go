package main

type Error string

const (
	EOF = Error("eof")
)
func main() {
	var four int
	four = func () int {
		return four
	}()
}
