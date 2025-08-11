package image

import (
	"bytes"
	"errors"
	"fmt"
	img "image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"path"
	"slices"
	"strings"
	"time"
	"welltoon/internal/dto"

	"github.com/chai2010/webp"
)

func Validate(filename string) error {
	imgExt := []string{".jpg", ".png"}
	imgFilename := strings.ToLower(filename)
	ext := path.Ext(imgFilename)
	if !slices.Contains(imgExt, ext) {
		return errors.New("extention image most be .jpg, .png")
	}
	return nil
}

func CompressToWebp(image *multipart.FileHeader) (*dto.WebpFile, error) {
	file, err := image.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var decode img.Image
	filename := strings.ToLower(image.Filename)
	imgExt := path.Ext(filename)

	switch imgExt {
	case ".jpg":
		dcd, err := jpeg.Decode(file)
		if err != nil {
			return nil, err
		}
		decode = dcd
	case ".png":
		dcd, err := png.Decode(file)
		if err != nil {
			return nil, err
		}
		decode = dcd
	default:
		return nil, errors.New("image extention most be .jpg, .png")
	}

	var buf bytes.Buffer
	if err := webp.Encode(&buf, decode, &webp.Options{Lossless: false, Quality: 75}); err != nil {
		return nil, err
	}

	newFilename := fmt.Sprintf("%d.webp", time.Now().UnixNano())
	content := bytes.NewReader(buf.Bytes())
	size := buf.Len()
	newFile := &dto.WebpFile{
		Filename: newFilename,
		Content:  content,
		Size:     int64(size),
	}
	return newFile, nil
}
