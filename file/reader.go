// Package file contains common functions needed to work with files.
package file

import (
    "os"
    "io"
    "io/ioutil"
    "bufio"
    "strings"
)

// ReadAll reads an entire file from a path and returns the contents as a string.
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

// ReadLines reads an entire file from a path and returns
// the contents as a slice of lines, with each line being a string.
func ReadLines(path string) ([]string, error) {

    f, err := os.Open(path)
    defer f.Close()
    if err != nil {
        return []string{}, err
    }

    lines := make([]string, 0, 1)
    bufReaderPtr := bufio.NewReader(f)

	// Reads until EOF or error
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
