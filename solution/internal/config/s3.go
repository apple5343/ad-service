package config

import "github.com/ilyakaznacheev/cleanenv"

type S3Config interface {
	Access() string
	Secret() string
	Endpoint() string
	Bucket() string
	Domain() string
}

type s3Config struct {
	AccessField   string `env:"S3_ACCESS"`
	SecretField   string `env:"S3_SECRET"`
	EndpointField string `env:"S3_ENDPOINT"`
	BucketField   string `env:"S3_BUCKET"`
	DomainField   string `env:"S3_DOMAIN"`
}

func NewS3Config() (S3Config, error) {
	cfg := s3Config{}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (s *s3Config) Access() string {
	return s.AccessField
}

func (s *s3Config) Bucket() string {
	return s.BucketField
}

func (s *s3Config) Domain() string {
	return s.DomainField
}

func (s *s3Config) Endpoint() string {
	return s.EndpointField
}

func (s *s3Config) Secret() string {
	return s.SecretField
}
