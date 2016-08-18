package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	count := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)

	for {
		rune_, nbytes, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if rune_ == unicode.ReplacementChar && nbytes == 1 {
			invalid++
			continue
		}
		count[rune_]++
		utflen[nbytes]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range count {
		fmt.Printf("%d\t%d\n", c, n)
	}
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 chars\n", invalid)
	}
}
