package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinio() *minio.Client {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("MINIO_SECRET_ACCESS_KEY")
	useSSL, err := strconv.ParseBool(os.Getenv("MINIO_SSL"))
	if err != nil {
		log.Fatalf("parse string to bool error most be true or false not: %s", os.Getenv("MINIO_SSL"))
	}
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("minio error: %s", err.Error())
	}
	if err := createBuckets(minioClient); err != nil {
		log.Fatalf("minio bucket error: %s", err.Error())
	}
	return minioClient
}

func createBuckets(minioClient *minio.Client) error {
	ctx := context.Background()
	bucket := os.Getenv("MINIO_BUCKETS")
	found, err := minioClient.BucketExists(ctx, bucket)
	if err != nil {
		return err
	}
	if !found {
		if err := minioClient.MakeBucket(ctx, bucket, minio.MakeBucketOptions{ObjectLocking: true}); err != nil {
			return err
		}
	}

	policy := fmt.Sprintf(`{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::%s/*"],"Sid": ""}]}`, bucket)
	if err := minioClient.SetBucketPolicy(ctx, bucket, policy); err != nil {
		return err
	}
	return nil
}
