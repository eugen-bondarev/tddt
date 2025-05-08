package storage

import (
	"context"
	"os"

	"github.com/eugen-bondarev/backup-tool/config"
	"github.com/eugen-bondarev/backup-tool/util"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type S3Storage struct {
	client *minio.Client
}

func NewS3Storage(c config.S3Config) (*S3Storage, error) {
	endpoint := util.SanitizeEndpoint(c.Endpoint)
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.AccessKeyID, c.SecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	return &S3Storage{
		client: client,
	}, nil
}

func (s *S3Storage) Push(localFileName, remoteFileName string) error {
	bucketName, key := util.GetBucketName(remoteFileName)

	file, err := os.Open(localFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	_, err = s.client.PutObject(context.Background(), bucketName, key, file, stat.Size(), minio.PutObjectOptions{})

	return err
}
