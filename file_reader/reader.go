package file_reader

import (
	"os"
	"io"
	"strings"
)

func ReadAll(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	var contents strings.Builder
	readTo := make([]byte, 50)
	for n, err := f.Read(readTo); (n != 0) && (err != io.EOF); n, err = f.Read(readTo) {
		toWrite := readTo[:n-1] // to remove last newline
		contents.WriteString(string(toWrite))
	}

	return contents.String(), nil
}
