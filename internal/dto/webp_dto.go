package dto

import "io"

type WebpFile struct {
	Filename string
	Content  io.Reader
	Size     int64
}
