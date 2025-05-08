package storage

import (
	"errors"
	"log"

	"github.com/eugen-bondarev/backup-tool/config"
)

type Storage interface {
	Push(localFileName, remoteFileName string) error
}

func New(cfg config.Config) (Storage, error) {
	if cfg.GCPConfig.IsDefined() {
		log.Printf("storage: using gcp backend")
		return NewGCPStorage(cfg.GCPConfig)
	}

	if cfg.S3Config.IsDefined() {
		log.Printf("storage: using s3 backend")
		return NewS3Storage(cfg.S3Config)
	}

	return nil, errors.New("no storage config")
}
