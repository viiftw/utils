package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"io"
	"bytes"
	"crypto/md5"
)

// CountFileInPath return number of files in path except directory
func CountFileInPath(path string) (int, error) {
	i := 0
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return 0, err
	}
	for _, file := range files {
		if !file.IsDir() {
			i++
		}
	}
	return i, nil
}

// FileIsExists return true if file is exists
func FileIsExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// PathIsDir return true if path is directory
func PathIsDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

// ListFilesInDir return all files name in path
func ListFilesInDir(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, len(files))
	for _, file := range files {
		if !file.IsDir() {
			result = append(result, file.Name())
		}
	}
	return result, nil
}

// GetFileSize returns the size of a file or zero in case of an error.
func GetFileSize(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return info.Size()
}

// CreateFile create a file with file name (default mode 0666)
func CreateFile(path string) error {
	if FileIsExists(path) {
		now := time.Now()
		return os.Chtimes(path, now, now)
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	return file.Close()
}

// AddToFile add bytes data to file
func AddToFile(path string, data []byte) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}

// ReadFile read file return bytes data
func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// DeleteFile delete a file
func DeleteFile(path string) error{
	return os.Remove(path)
}

// GetMD5File return md5 hash of file as string
func GetMD5File(path string) (string, error) {
	data, err := ReadFile(path)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, err = io.Copy(hash, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}