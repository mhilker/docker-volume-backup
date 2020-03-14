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

	log.Println(id)
	log.Println(secret)
	log.Println(region)
	log.Println(bucket)

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
		filename := dir.Name + ".tar.gz" 
		file, err := dockervolumebackup.CreateArchive(dir.Path, filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		key := now + "/" + filename
		path, err := backup.UploadFile(bucket, key, file)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(path)
	}
}
