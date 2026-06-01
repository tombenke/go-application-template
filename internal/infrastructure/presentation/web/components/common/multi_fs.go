package common

import (
	"io/fs"
)

// MultiFS is a custom file system that combines multiple fs.FS instances.
type MultiFS struct {
	fss []fs.FS
}

// NewMultiFS creates a new MultiFS instance by combining the provided file systems.
func NewMultiFS(fss ...fs.FS) *MultiFS {
	// Combine the provided file systems
	return &MultiFS{
		fss: fss,
	}
}

// Open implements the fs.FS interface.
// It tries to open the file from each of the combined file systems in order
// until it finds it or returns an error if not found.
func (m *MultiFS) Open(name string) (fs.File, error) {
	var err error
	for _, f := range m.fss {
		file, errFS := f.Open(name)
		if errFS == nil {
			return file, nil
		}
		err = errFS
	}
	return nil, err
}
