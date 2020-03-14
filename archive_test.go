package dockervolumebackup

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateArchive(t *testing.T) {
	file, err := CreateArchive("testdata/archive/")
	require.Nil(t, err)
	require.FileExists(t, file)
	defer os.Remove(file)

	// TODO: Test archive contains both files
}
