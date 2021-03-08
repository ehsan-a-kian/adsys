package gdm_test

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/termie/go-shutil"
	"github.com/ubuntu/adsys/internal/policies/dconf"
	"github.com/ubuntu/adsys/internal/policies/entry"
	"github.com/ubuntu/adsys/internal/policies/gdm"
)

var update bool

func TestApplyPolicy(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		entries []entry.Entry

		wantErr bool
	}{
		// user cases
		"dconf policy": {entries: []entry.Entry{
			{Key: "dconf/com/ubuntu/category/key-s", Value: "'onekey-s-othervalue'", Meta: "s"}}},
	}

	for name, tc := range tests {
		tc := tc
		name := name
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			dconfDir := t.TempDir()

			// Apply machine configuration
			dconfManager := dconf.NewWithDconfDir(dconfDir)
			err := dconfManager.ApplyPolicy(context.Background(), "ubuntu", true, nil)
			require.NoError(t, err, "ApplyPolicy failed but shouldn't have")

			m, err := gdm.New(gdm.WithDconf(dconfManager))
			require.NoError(t, err, "Setup: can't create gdm manager")

			err = m.ApplyPolicy(context.Background(), tc.entries)
			if tc.wantErr {
				require.NotNil(t, err, "ApplyPolicy should have failed but didn't")
				return
			}
			require.NoError(t, err, "ApplyPolicy failed but shouldn't have")

			goldPath := filepath.Join("testdata", "golden", name, "etc", "dconf")
			// Update golden file
			if update {
				t.Logf("updating golden file %s", goldPath)
				require.NoError(t, os.RemoveAll(goldPath), "Cannot remove target golden directory")
				// Filter dconf generated DB files that are machine dependent
				require.NoError(t,
					shutil.CopyTree(
						dconfDir, goldPath,
						&shutil.CopyTreeOptions{Symlinks: true, Ignore: ignoreDconfDB, CopyFunction: shutil.Copy}),
					"Can’t update golden directory")
			}

			gotContent := treeContent(t, dconfDir, []byte("GVariant"))
			goldContent := treeContent(t, goldPath, nil)
			assert.Equal(t, goldContent, gotContent, "got and expected content differs")

			// Verify that each <DB>.d has a corresponding gvariant db generated by dconf update
			dbs, err := filepath.Glob(filepath.Join(dconfDir, "db", "*.d"))
			require.NoError(t, err, "Checking pattern for dconf db failed")
			for _, db := range dbs {
				_, err = os.Stat(strings.TrimSuffix(db, ".db"))
				assert.NoError(t, err, "Binary version of dconf DB should exists")
			}
		})
	}
}

// treeContent build a recursive file list of dir with their content
func treeContent(t *testing.T, dir string, ignoreHeaders []byte) map[string]string {
	t.Helper()

	r := make(map[string]string)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("couldn't access path %q: %v", path, err)
		}

		content := ""
		if !info.IsDir() {
			d, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			// ignore given header
			if ignoreHeaders != nil && bytes.HasPrefix(d, ignoreHeaders) {
				return nil
			}
			content = string(d)
		}
		r[strings.TrimPrefix(path, dir)] = content
		return nil
	})

	if err != nil {
		t.Fatalf("error while listing directory: %v", err)
	}

	return r
}

func ignoreDconfDB(src string, entries []os.FileInfo) []string {
	var r []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		d, err := os.ReadFile(filepath.Join(src, e.Name()))
		if err != nil {
			continue
		}

		if bytes.HasPrefix(d, []byte("GVariant")) {
			r = append(r, e.Name())
		}
	}
	return r
}

func TestMain(m *testing.M) {
	flag.BoolVar(&update, "update", false, "update golden files")
	flag.Parse()

	m.Run()
}
