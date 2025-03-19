package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/caddyserver/certmagic"
)

func TestFileStorage(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "certmagic-storage-tests-")
	if err != nil {
		t.Fatalf("Cannot create temp directory: %s", err)
	}
	defer os.RemoveAll(tempDir)
	fs := &certmagic.FileStorage{
		Path: filepath.Join(tempDir, "filestorage"),
	}
	NewTestSuite(fs).Run(t)
}
