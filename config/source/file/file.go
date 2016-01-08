package file

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/micro/go-platform/config"
)

// Currently a single file reader
type file struct {
	opts config.SourceOptions
}

func (f *file) Read() (*config.ChangeSet, error) {
	fh, err := os.Open(f.opts.Name)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	b, err := ioutil.ReadAll(fh)
	if err != nil {
		return nil, err
	}
	info, err := fh.Stat()
	if err != nil {
		return nil, err
	}

	// hash the file
	h := md5.New()
	h.Write(b)
	checksum := fmt.Sprintf("%x", h.Sum(nil))

	return &config.ChangeSet{
		Source:    f.String(),
		Timestamp: info.ModTime(),
		Data:      b,
		Checksum:  checksum,
	}, nil
}

func (f *file) String() string {
	return "file"
}

func (f *file) Watch() (config.Watcher, error) {
	if _, err := os.Stat(f.opts.Name); err != nil {
		return nil, err
	}
	return &watcher{f: f}, nil
}

func NewSource(opts ...config.SourceOption) config.Source {
	var options config.SourceOptions
	for _, o := range opts {
		o(&options)
	}
	return &file{opts: options}
}
