// @filename: file.go
// @author: Krisna Pranav
// @license: Copyright (c) 2023, godzillaframework, Krisna Pranav

package index

import (
	"io"
	"io/fs"
)

// file[Stat]
type File interface {
	io.ReadWriteCloser

	Stat() (fs.FileInfo, error)
}
