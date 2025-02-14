package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/caarlos0/env/v11"
	"github.com/eugen-bondarev/backup-tool/util"
	"github.com/joho/godotenv"
)

const (
	tmpDir = "tmp"
)

func main() {
	if _, err := os.Stat(tmpDir); os.IsNotExist(err) {
		os.MkdirAll(tmpDir, 0755)
	}

	godotenv.Load()
	cfg, err := env.ParseAs[Config]()
	util.CheckErr(err)
	fmt.Println(cfg)

	d := NewDump("test", filepath.Join(tmpDir, "dump.sql"))
	err = d.Create(cfg, MySQL)
	util.CheckErr(err)
}
