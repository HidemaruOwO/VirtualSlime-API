package lib

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/HidemaruOwO/nuts/log"
)

func DownloadCache() error {
	log.Info("Downloading cache...")
	url := "https://v-sli.me/data.json"
	cacheDir := "cache"

	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		os.Mkdir(cacheDir, os.ModePerm)
	}

	filePath := filepath.Join(cacheDir, "data.json")
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	log.Info("Downloaded cache!")
	return nil
}

func ReadCache() []byte {
	cacheDir := "cache"
	filePath := filepath.Join(cacheDir, "data.json")
	file, err := os.Open(filePath)
	if err != nil {
		log.Error(err)
		return nil
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Error(err)
		return nil
	}
	fileSize := fileInfo.Size()

	buf := make([]byte, fileSize)

	_, err = file.Read(buf)
	if err != nil {
		log.Error(err)
		return nil
	}

	return buf
}
