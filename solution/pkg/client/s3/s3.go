package s3

import (
	"server/internal/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewS3Cliet(config config.S3Config) (*minio.Client, error) {
	client, err := minio.New(config.Endpoint(), &minio.Options{
		Creds:  credentials.NewStaticV4(config.Access(), config.Secret(), ""),
		Secure: true,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}
