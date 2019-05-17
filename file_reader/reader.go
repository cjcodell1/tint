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
		// REMOVE ALL THE EMPTY BYTES AT THE END OF THE SLICE
		contents.WriteString(string(readTo))
	}

	return contents.String(), nil
}
