package util

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	acceptedExtension = map[string]bool{
		".txt": true,
		".log": true,
	}
)

// GetOutputFile - get output file
func GetOutputFile(fp string) (*os.File, error) {
	info, err := os.Stat(fp)
	if err != nil {
		return nil, fmt.Errorf("file %s does not exist", fp)
	}
	if info.IsDir() {
		return nil, fmt.Errorf("can't write output to a directory")
	}
	fileExtension := filepath.Ext(fp)
	if !acceptedExtension[fileExtension] {
		return nil, fmt.Errorf("unsupported extension %s", fileExtension)
	}
	return os.Open(fp)
}
