package appstore_sdk

import (
	"io/ioutil"
	"os"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func readFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	return data, err
}
