package storage

import (
	"context"
	"encoding/base64"
	"io"
	"os"

	"cloud.google.com/go/storage"
	"github.com/eugen-bondarev/backup-tool/config"
	"github.com/eugen-bondarev/backup-tool/util"
	"google.golang.org/api/option"
)

type GCPStorage struct {
	client *storage.Client
}

func NewGCPStorage(c config.GCPConfig) (*GCPStorage, error) {
	b64Decoded, err := base64.StdEncoding.DecodeString(c.Credentials)
	if err != nil {
		return nil, err
	}

	client, err := storage.NewClient(context.Background(), option.WithCredentialsJSON(b64Decoded))
	if err != nil {
		return nil, err
	}

	return &GCPStorage{
		client: client,
	}, nil
}

func (s *GCPStorage) Push(localFileName, remoteFileName string) error {
	bucketName, path := util.GetBucketName(remoteFileName)
	bucket := s.client.Bucket(bucketName)
	object := bucket.Object(path)
	writer := object.NewWriter(context.Background())

	reader, err := os.Open(localFileName)
	if err != nil {
		return err
	}
	defer reader.Close()

	_, err = io.Copy(writer, reader)
	if err != nil {
		return err
	}

	return writer.Close()
}
