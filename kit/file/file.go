package utilFile

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"pandax/kit/biz"
	"path/filepath"
	"strconv"
	"sync"
)

const (
	MaxConcurrency = 16 // 最大并发数
)

type DownloadTask struct {
	URL      string
	FilePath string
}

func DownloadFileWithConcurrency(url, filepath string) error {
	resp, err := http.Head(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fileSize, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// 检查本地文件大小
	localFileSize, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}

	// 计算剩余未下载的文件大小
	remainingSize := fileSize - int(localFileSize)

	// 计算每个片段的大小
	chunkSize := remainingSize / MaxConcurrency

	// 创建等待组，用于等待所有goroutine完成
	var wg sync.WaitGroup
	wg.Add(MaxConcurrency)

	// 创建并发下载任务
	for i := 0; i < MaxConcurrency; i++ {
		start := localFileSize + int64(i*chunkSize)
		end := start + int64(chunkSize) - 1

		// 最后一个片段的结束位置可能超过文件大小，需要修正
		if i == MaxConcurrency-1 {
			end = int64(fileSize) - 1
		}

		go func(index int, start, end int64) {
			defer wg.Done()

			err := downloadChunk(url, filepath, start, end)
			if err != nil {
				biz.NewBizErr("文件下载失败")
				// 处理下载错误
			}
		}(i, start, end)
	}

	// 等待所有goroutine完成
	wg.Wait()

	return nil
}

func downloadChunk(url, filepath string, start, end int64) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// 设置Range头部
	req.Header.Set("Range", "bytes="+strconv.FormatInt(start, 10)+"-"+strconv.FormatInt(end, 10))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Seek(start, io.SeekStart)
	if err != nil {
		return err
	}

	_, err = io.CopyN(file, resp.Body, end-start+1)
	if err != nil {
		return err
	}

	return nil
}

func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func IsExcl(fpath string) bool {
	ext := filepath.Ext(fpath)
	switch ext {
	case "xls", "xlsx":
		return true
	default:
		return false
	}
}
