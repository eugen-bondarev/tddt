package util

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const (
	tmpDir = "tmp"
)

func createTmpName() string {
	return uuid.New().String()
}

func CreateTmpSqlFilePath() string {
	return filepath.Join(tmpDir, fmt.Sprintf("%s.sql", createTmpName()))
}

func EnsureTmpDir() {
	if _, err := os.Stat(tmpDir); os.IsNotExist(err) {
		os.MkdirAll(tmpDir, 0755)
	}
}
