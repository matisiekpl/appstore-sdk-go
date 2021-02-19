package appstore

import (
	"io/ioutil"
	"os"
)

//Check file is exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//Read file
func readFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	return data, err
}
