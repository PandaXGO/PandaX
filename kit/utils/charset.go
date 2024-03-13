package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
)

var (
	// Alias for charsets.
	charsetAlias = map[string]string{
		"HZGB2312": "HZ-GB-2312",
		"hzgb2312": "HZ-GB-2312",
		"GB2312":   "HZ-GB-2312",
		"gb2312":   "HZ-GB-2312",
	}
)

func Convert(dstCharset string, srcCharset string, src string) (dst string, err error) {
	if dstCharset == srcCharset {
		return src, nil
	}
	dst = src
	// Converting `src` to UTF-8.
	e, err := getEncoding(srcCharset)
	if err != nil {
		return "", err
	}
	if srcCharset != "UTF-8" {
		if e != nil {
			tmp, err := io.ReadAll(
				transform.NewReader(bytes.NewReader([]byte(src)), e.NewDecoder()),
			)
			if err != nil {
				return "", errors.New(fmt.Sprintf(`convert string "%s" to utf8 failed`, srcCharset))
			}
			src = string(tmp)
		} else {
			return dst, errors.New(fmt.Sprintf(`unsupported srcCharset "%s"`, srcCharset))
		}
	}
	// Do the converting from UTF-8 to `dstCharset`.
	if dstCharset != "UTF-8" {
		if e != nil {
			tmp, err := io.ReadAll(
				transform.NewReader(bytes.NewReader([]byte(src)), e.NewEncoder()),
			)
			if err != nil {
				return "", errors.New(fmt.Sprintf(`convert string from utf8 to "%s" failed`, srcCharset))
			}
			dst = string(tmp)
		} else {
			return dst, errors.New(fmt.Sprintf(`unsupported dstCharset "%s"`, srcCharset))
		}
	} else {
		dst = src
	}
	return dst, nil
}

func ToUTF8(srcCharset string, src string) (dst string, err error) {
	return Convert("UTF-8", srcCharset, src)
}

func UTF8To(dstCharset string, src string) (dst string, err error) {
	return Convert(dstCharset, "UTF-8", src)
}

func getEncoding(charset string) (encoding.Encoding, error) {
	if c, ok := charsetAlias[charset]; ok {
		charset = c
	}
	return ianaindex.MIB.Encoding(charset)
}
