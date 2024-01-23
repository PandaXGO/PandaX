package oss

import (
	"net/http"
	"os"
	"path"
	"time"

	"github.com/google/uuid"
)

func BuildName(fileName, fileType string) (string, error) {
	t := time.Now().Format("20060102")
	newUUID, _ := uuid.NewUUID()
	ext := path.Ext(fileName)
	return t + "/" + fileType + "/" + newUUID.String() + ext, nil
}

func GetFileContentType(fileName string) (string, error) {

	out, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err = out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
