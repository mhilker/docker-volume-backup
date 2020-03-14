package dockervolumebackup

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func ListFilesInDirectory(directory string) map[string]os.FileInfo {
	infos := make(map[string]os.FileInfo, 0)

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		infos[path] = info
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	return infos
}

func CreateArchive(sourcePath string, outputPath string) (*os.File, error) {
	file, err := os.Create(outputPath)
	if err != nil {
		return nil, err
	}

	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	err = filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		return addFileToArchive(path, tarWriter)
	})

	if err != nil {
		return nil, err
	}

	return file, nil
}

func addFileToArchive(filePath string, tarWriter *tar.Writer) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	header := &tar.Header{
		Name:    filePath,
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
