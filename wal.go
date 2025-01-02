package wal

import (
	"bufio"
	"context"
	"os"
	"sync"
	"time"
)

const (
	syncInterval  = 200 * time.Millisecond
	segmentPrefix = "segment-"
)

type WAL struct {
	dir             string
	currentSeg      *os.File
	lock            sync.Mutex
	lsn             uint64
	bufWriter       *bufio.Writer
	syncTimer       *time.Timer
	doFsync         bool
	mxFileSize      int64
	mxSegements     int
	currentSegIndex int
	ctx             context.Context
	cancel          context.CancelFunc
}

func OpenWAL(dir string, doFsync bool, mxFileSize int64, mxSegements int) (*WAL, error) {
	ctx, cancel := context.WithCancel(context.Background())
	wal := &WAL{
		dir:         dir,
		doFsync:     doFsync,
		mxFileSize:  mxFileSize,
		mxSegements: mxSegements,
		ctx:         ctx,
		cancel:      cancel,
	}
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, err
	}
	return wal, nil

}

func (w *WAL) Close() error {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.cancel()

	if w.currentSeg != nil {
		err := w.currentSeg.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
