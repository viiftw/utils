# Utils - A lightweight go utility library

## Use

```go
package main

import (
  "fmt"
  "github.com/viiftw/utils"
)

func main() {
  filename := "test.txt"
  err := utils.CreateFile(filename)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(utils.FileIsExists(filename))
}
// result: true
```

## Built-in Functions

### File

```go
  CreateFile(path string) error
  AddToFile(path string, data []byte) error
  ReadFile(path string) ([]byte, error)
  DeleteFile(path string) error
  CountFileInPath(path string) (int, error)
  ListFilesInDir(path string) ([]string, error)
  FileIsExists(path string) bool
  PathIsDir(path string) bool
  GetFileSize(path string) int64
  GetMD5File(path string) (string, error)
```
