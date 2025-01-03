package tests

import (
	"os"
	"testing"

	wal "github.com/anmit007/writeAheadLogGo"
	"github.com/stretchr/testify/assert"
)

const (
	maxSegements = 3
	maxFileSize  = 64 * 1000 * 1000
)

func TestWAL_OpenClose(t *testing.T) {
	t.Parallel()
	dirPath := "TestWaL.OpenClose.log"
	defer os.RemoveAll(dirPath)
	w, err := wal.OpenWAL(dirPath, true, maxFileSize, maxSegements)
	assert.NoError(t, err, "Failed to create WAL")
	assert.NotNil(t, w, "wal is nil")

	err = w.Close()
	assert.NoError(t, err, "Failed to close WAL")
}
