package oss

import (
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
