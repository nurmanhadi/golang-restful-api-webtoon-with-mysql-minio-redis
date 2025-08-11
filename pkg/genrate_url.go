package pkg

import (
	"fmt"
	"os"
	"strconv"
)

func S3GenerateUrl(filename string) (string, error) {
	entpoint := os.Getenv("MINIO_ENDPOINT")
	bucket := os.Getenv("MINIO_BUCKETS")
	ssl, err := strconv.ParseBool(os.Getenv("MINIO_SSL"))
	if err != nil {
		return "", err
	}
	if ssl {
		return fmt.Sprintf("https://%s/%s/%s", entpoint, bucket, filename), nil
	} else {
		return fmt.Sprintf("http://%s/%s/%s", entpoint, bucket, filename), nil
	}
}
