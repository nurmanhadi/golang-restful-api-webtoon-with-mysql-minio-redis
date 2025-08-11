package repository

import "welltoon/internal/dto"

type S3Repository interface {
	PutObject(object *dto.WebpFile) error
	RemoveObject(filename string) error
}
