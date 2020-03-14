# Docker Volume Backup

[![build](https://github.com/mhilker/docker-volume-backup/workflows/build/badge.svg)](https://github.com/mhilker/docker-volume-backup/actions)
[![license](https://img.shields.io/github/license/mhilker/docker-volume-backup)](./LICENSE.md)
![go version](https://img.shields.io/github/go-mod/go-version/mhilker/docker-volume-backup)

This tool creates .tar.gz archives from an volume with the label `com.github.mhilker.docker-volume-backup` and uploads them to AWS S3.

## Example

Create a new docker volume:

``` bash
docker volume create --label com.github.mhilker.docker-volume-backup my-test-volume
```

Create some files in the volume:

```bash
export VOLUME_PATH="$(sudo docker volume inspect my-test-volume -f '{{.Mountpoint}}')"
echo "lorem" | sudo --preserve-env=VOLUME_PATH tee "$VOLUME_PATH/lorem"
echo "ipsum" | sudo --preserve-env=VOLUME_PATH tee "$VOLUME_PATH/ipsum"
```

```bash
$ git clone https://github.com/mhilker/docker-volume-backup
$ cd docker-volume-backup
$ export AWS_ID="XXXXXXXXXXXXXXXXXXXX"
$ export AWS_SECRET="XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
$ export AWS_REGION="eu-central-1"
$ export AWS_BUCKET="your-bucket-name"
$ sudo --preserve-env=AWS_ID,AWS_SECRET,AWS_REGION,AWS_BUCKET /usr/local/go/bin/go run cmd/docker-volume-backup/main.go
2020/03/14 20:26:13 Adding file /var/lib/docker/volumes/my-test-volume/_data/ipsum /ipsum
2020/03/14 20:26:13 Adding file /var/lib/docker/volumes/my-test-volume/_data/lorem /lorem
2020/03/14 20:26:13 Uploaded archive to https://your-bucket-name.s3.eu-central-1.amazonaws.com/my-test-volume/2020-03-14T19%3A26%3A13Z.tar.gz
```

## Build and run

### On your machine

```bash
go build -o build/docker-volume-backup ./cmd/docker-volume-backup/main.go
./build/docker-volume-backup
```

## Code Style

```bash
go fmt .
```

## License

The MIT License (MIT). Please see the [license file](LICENSE.md) for more information.