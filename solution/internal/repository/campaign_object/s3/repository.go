package s3

import (
	"context"
	"server/internal/config"
	"server/internal/model"
	"server/internal/repository"

	"github.com/minio/minio-go/v7"
)

type campaignObjectRepository struct {
	s3Client *minio.Client
	cfg      config.S3Config
}

func NewCampaignObjectRepository(s3Client *minio.Client, s config.S3Config) repository.CampaignObjectRepository {
	return &campaignObjectRepository{
		s3Client: s3Client,
		cfg:      s,
	}
}

func (r *campaignObjectRepository) SaveImage(ctx context.Context, image *model.Image) (*model.Image, error) {
	_, err := r.s3Client.PutObject(ctx, r.cfg.Bucket(), image.Path, image.Data, -1, minio.PutObjectOptions{ContentType: image.Type})
	if err != nil {
		return nil, err
	}
	image.URL = r.cfg.Domain() + "/" + image.Path
	return image, nil
}
