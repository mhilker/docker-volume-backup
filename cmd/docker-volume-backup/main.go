package main

import (
	"log"
	"os"
	"time"
)

func main() {
	id := os.Getenv("AWS_ID") 
	secret := os.Getenv("AWS_SECRET")
	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("AWS_BUCKET")
	filename := "README.md"
    key := time.Now().UTC().Format(time.RFC3339) + "/" + filename

    file, err := os.Open(filename)
    if err != nil {
		log.Fatal("failed to open file")
        return
    }

	b := docker_volume_backup.NewBackup(id, secret, region)
	b.UploadFile(bucket, key, file)
}

	// dirs := listLocalBackupDirectories()
	// for _, dir := range dirs {
	// 	files := listFilesInDirectory(dir.Path)
	// 	for path, file := range files {
	// 		if !file.IsDir() {
	// 			key := strings.TrimPrefix(path, dir.Path)
	// 			key = strings.TrimLeft(key, "/")
	// 			key = dir.Name + "/" + now + "/" + key

	// 			fmt.Println(file.Mode())

	// 			uploadFile(path, region, bucket, key)
	// 		}
	// 	}
	// }