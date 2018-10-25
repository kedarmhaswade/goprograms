// Write a program wordfreq to report the frequency of each word in an input text file.
// Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into words instead of lines.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	iScan := bufio.NewScanner(os.Stdin)
	freq := make(map[string]int)
	iScan.Split(bufio.ScanWords)
	for iScan.Scan() {
		w := iScan.Text()
		freq[w] += 1
	}
	fmt.Printf("%s\t\t%s\n", "word", "freq")
	for w, f := range freq {
		fmt.Printf("%s\t\t%d\n", w, f)
	}
}
