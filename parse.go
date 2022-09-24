package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// Read lease entries from a file
func ParseLeases(r io.Reader) {
	s := bufio.NewScanner(r)
	s.Split(func(data []byte, atEOF bool) (int, []byte, error) {
		if atEOF {
			return 0, nil, bufio.ErrFinalToken
		}

		// Find line starting with "lease"
		if i := bytes.Index(data, []byte("lease")); i != -1 {
			i += 6

			// Find the closing bracket
			if j := bytes.Index(data[i:], []byte("}")); j != -1 {
				return i + j, data[i : i+j], nil
			}
		}

		return 0, nil, bufio.ErrFinalToken
	})

	for s.Scan() {
		fmt.Println(string(s.Bytes()))
	}
}
