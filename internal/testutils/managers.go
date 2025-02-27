package testutils

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/termie/go-shutil"
)

// MockAssetsDumper is a mock implementation of the AssetsDumper function which
// uncompresses policy assets to a directory.
type MockAssetsDumper struct {
	Path string

	T           *testing.T
	Err         bool
	ReadOnlyErr bool
}

// SaveAssetsTo uncompresses policy assets to a directory.
// It returns an error if the Err field is set to true.
// It returns an error if Path is different than the dest exercised by the manager.
func (m MockAssetsDumper) SaveAssetsTo(ctx context.Context, relSrc, dest string, uid int, gid int) (err error) {
	if m.Err {
		return errors.New("SaveAssetsTo error")
	}

	if m.ReadOnlyErr {
		if err := shutil.CopyTree(fmt.Sprintf("testdata/sysvol-%s", m.Path), dest, nil); err != nil {
			return fmt.Errorf("SaveAssetsTo: unexpected error when dumping assets: %w", err)
		}

		// Make all child files and directories read-only, leaving execute permissions for directories
		err = filepath.WalkDir(dest, func(path string, d fs.DirEntry, err error) error {
			require.NoError(m.T, err, "SaveAssetsTo: unexpected error when walking directory")
			perm := 0400
			if d.IsDir() {
				perm = 0500
			}
			err = os.Chmod(path, os.FileMode(perm))
			require.NoError(m.T, err, "SaveAssetsTo: unexpected error when changing permissions")

			return nil
		})
		require.NoError(m.T, err, "SaveAssetsTo: unexpected error when making files read-only")
		return nil
	}

	if relSrc != m.Path {
		return fmt.Errorf("SaveAssetsTo: unexpected relSrc: %q", relSrc)
	}
	return shutil.CopyTree(fmt.Sprintf("testdata/sysvol-%s", m.Path), dest, nil)
}
