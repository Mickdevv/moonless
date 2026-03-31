package utils

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func CreateStaticFile(serverCfg *ServerCfg, subdirectoryPath string, r *http.Request) (string, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return "", err
	}

	file, handler, err := r.FormFile("file")
	if file == nil {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	defer file.Close()

	filename := strconv.Itoa(int(time.Now().UnixMilli())) + filepath.Base(handler.Filename)
	destinationPath := filepath.Join(serverCfg.STATIC_FILES_DIR, subdirectoryPath, filename)

	dst, err := os.Create(destinationPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	return destinationPath, nil
}
