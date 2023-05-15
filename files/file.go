package files

import (
	"os"
)

func CreateDir(path string) error {
	if err := os.Mkdir(path, os.ModePerm); err != nil {
		return err
	}
	os.Chmod(path, os.ModePerm)
	return nil
}

func OpenFile(f string) (*os.File, error) {
	return os.OpenFile(f, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
}
