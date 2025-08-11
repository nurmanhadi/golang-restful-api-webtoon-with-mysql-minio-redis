package s3

import (
	"context"
	"os"
	"welltoon/internal/dto"
	"welltoon/internal/repository"

	"github.com/minio/minio-go/v7"
)

type s3Repository struct {
	ctx context.Context
	s3  *minio.Client
}

func NewS3(ctx context.Context, s3 *minio.Client) repository.S3Repository {
	return &s3Repository{
		ctx: ctx,
		s3:  s3,
	}
}

func (r *s3Repository) PutObject(object *dto.WebpFile) error {
	bucket := os.Getenv("MINIO_BUCKETS")

	_, err := r.s3.PutObject(r.ctx, bucket, object.Filename, object.Content, object.Size, minio.PutObjectOptions{ContentType: "image/webp"})
	if err != nil {
		return err
	}
	return nil
}
func (r *s3Repository) RemoveObject(filename string) error {
	bucket := os.Getenv("MINIO_BUCKETS")
	err := r.s3.RemoveObject(r.ctx, bucket, filename, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
