package main

import (
	"log"
	"os"
	"time"

	dockervolumebackup "github.com/mhilker/docker-volume-backup"
)

func main() {
	id := os.Getenv("AWS_ID")
	secret := os.Getenv("AWS_SECRET")
	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("AWS_BUCKET")
	now := time.Now().UTC().Format(time.RFC3339)

	backup := dockervolumebackup.NewBackup(id, secret, region)

	provider, err := dockervolumebackup.NewProvider()
	if err != nil {
		log.Fatal(err)
	}

	dirs, err := provider.GetDirectories("com.github.mhilker.docker-volume-backup")
	if err != nil {
		log.Fatal(err)
	}

	for _, dir := range dirs {
		file, err := dockervolumebackup.CreateArchive(dir.Path)
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(file)

		key := dir.Name + "/" + now + ".tar.gz"
		path, err := backup.UploadFile(bucket, key, file)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Uploaded archive to " + path)
	}
}
