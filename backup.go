package docker_volume_backup

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Backup struct {
	uploader *s3manager.Uploader
	region   string
}

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

func (b *Backup) UploadFile(bucket string, key string, file *os.File) (string, error) {
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
