package file_reader

import (
	"os"
    "io"
	"io/ioutil"
    "bufio"
    "strings"
)

func ReadAll(path string) (string, error) {
	f, err := os.Open(path)
    defer f.Close()
	if err != nil {
		return "", err
	}

	contents, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(contents), nil
}

func Lines(path string) ([]string, error) {

    f, err := os.Open(path)
    defer f.Close()
    if err != nil {
        return []string{}, err
    }

    lines := make([]string, 0, 1)
    bufReaderPtr := bufio.NewReader(f)
    for {
        line, err := bufReaderPtr.ReadString(byte('\n'))
        if err == io.EOF {
            return lines, nil
        }
        if err != nil {
            return lines, err
        }

        line = strings.TrimSuffix(line, "\n")

        // remove the \r if file has DOS endings 
        line = strings.TrimSuffix(line, "\r")

        lines = append(lines, line)
    }

   return lines, nil
}
