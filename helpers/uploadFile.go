package helpers

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

func FileUploadHandler(file *multipart.FileHeader, path string) (string, error) {
	// UPLOAD IMAGE
	var file_name string
	var fullPath string
	src, err := file.Open()
	if err != nil {
		return file_name, err
	}
	defer src.Close()

	file.Filename = strings.ReplaceAll(file.Filename, " ", "")
	name := file.Filename
	file_name = fmt.Sprintf(os.Getenv("APP_URL")+"/api/v1/public/%s/%s", path, name)
	dst, err := os.Create(fmt.Sprintf("public/%s/%s", path, name))
	if err != nil {
		return file_name, err
	}

	defer dst.Close()

	// // Copy
	if _, err = io.Copy(dst, src); err != nil {
		return file_name, err
	}

	fullPath = file_name

	return fullPath, nil
}
