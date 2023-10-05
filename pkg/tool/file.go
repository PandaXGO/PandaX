package tool

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func GetFileMd5(filePath string) (string, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	// 创建MD5哈希对象
	hash := md5.New()
	// 将文件内容拷贝到哈希对象中
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	// 计算MD5哈希值
	hashBytes := hash.Sum(nil)
	md5String := hex.EncodeToString(hashBytes)
	return md5String, nil
}
