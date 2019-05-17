package file_reader

import (
	"os"
	"io/ioutil"
)

func ReadAll(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	contents, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(contents), nil
}
