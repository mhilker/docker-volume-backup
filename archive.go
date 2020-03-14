package dockervolumebackup

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateArchive(sourcePath string) (string, error) {
	file, err := ioutil.TempFile(os.TempDir(), "temp")
	if err != nil {
		return "", err
	}
	defer file.Close()

	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	err = filepath.Walk(sourcePath, func(localPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		archivePath := strings.TrimPrefix(localPath, sourcePath)
		log.Println("Adding file", localPath, archivePath)
		return addFileToArchive(localPath, archivePath, tarWriter)
	})

	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func addFileToArchive(localPath string, archivePath string, tarWriter *tar.Writer) error {
	file, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	header := &tar.Header{
		Name:    archivePath,
		Size:    stat.Size(),
		Mode:    int64(stat.Mode()),
		ModTime: stat.ModTime(),
	}

	err = tarWriter.WriteHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(tarWriter, file)
	if err != nil {
		return err
	}

	return nil
}
