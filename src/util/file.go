package util

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func GetFilePath(filename string) (string, error) {
	if strings.HasPrefix(filename, "~") {
		u, err := user.Current()
		if err != nil {
			return filename, err
		}

		filename = strings.Replace(filename, "~", u.HomeDir, 1)
	} else {
		fn, err := filepath.Abs(filename)
		if err != nil {
			return filename, err
		}

		filename = fn
	}

	return filename, nil
}

// ReadConfig read config from filepath
func ReadFile(filename string) ([]byte, error) {
	filepath, err := GetFilePath(filename)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func MakeFolderOn(folderPath string) error {
	return os.MkdirAll(folderPath, os.ModePerm)
}
