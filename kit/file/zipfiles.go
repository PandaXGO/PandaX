package utilFile

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func ZipFiles(filename string, files []string, oldForm, newForm string) error {
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	for _, file := range files {
		err = compressFile(file, oldForm, newForm, zipWriter)
		if err != nil {
			return err
		}
	}

	return nil
}

func compressFile(file, oldForm, newForm string, zipWriter *zip.Writer) error {
	zipFile, err := os.Open(file)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	info, err := zipFile.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = strings.Replace(file, oldForm, newForm, -1)
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, zipFile)
	if err != nil {
		return err
	}

	return nil
}