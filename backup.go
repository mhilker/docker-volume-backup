package dockervolumebackup

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Backup contains methods to upload files to AWS S3
type Backup struct {
	uploader *s3manager.Uploader
}

// NewBackup creates a new backup process for AWS S3
func NewBackup(id string, secret string, region string) *Backup {
	config := aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(id, secret, ""),
	}
	sess := session.Must(session.NewSession(&config))

	return &Backup{
		uploader: s3manager.NewUploader(sess),
	}
}

// UploadFile uploads a file to AWS S3
func (b *Backup) UploadFile(bucket string, key string, path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	result, err := b.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})

	if err != nil {
		return "", err
	}

	return result.Location, nil
}
