package main

import (
	"errors"
	"io/ioutil"
	"os"
)

func readStdin() (string, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return "", err
	}

	if stat.Mode()&os.ModeNamedPipe == 0 && stat.Size() == 0 {
		return "", errors.New("no cards provided")
	}

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
