package docker_volume_backup

import (
	"context"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func listLocalBackupDirectories() []BackupDirectory {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	dirs := make([]BackupDirectory, 0)

	args := filters.NewArgs()
	args.Add("label", "com.github.mhilker.docker-volume-backup")

	response, err := cli.VolumeList(ctx, args)
	for _, v := range response.Volumes {
		dirs = append(dirs, BackupDirectory{
			Name: v.Name,
			Path: v.Mountpoint,
		})
	}

	return dirs
}
