package utils

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"
)

var uploadMutex sync.Mutex

func UploadImage(c *gin.Context) (string, error) {
	uploadDir := "./uploads"

	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	form, err := c.MultipartForm()
	if err != nil {
		return "", err
	}

	files := form.File["poster"]
	if len(files) == 0 {
		return "", errors.New("no image file found")
	}

	file := files[0]
	fileName := randstr.String(8) + filepath.Ext(file.Filename)
	filePath := filepath.Join(uploadDir, fileName)

	uploadMutex.Lock()
	defer uploadMutex.Unlock()

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		log.Println("Error saving uploaded file:", err)
		return "", err
	}

	filename := "http://localhost:3000/uploads/" + fileName 

	return filename, nil
}
