package models

import (
	"mime/multipart"

	"github.com/ipfs/boxo/files"
)

func MultipartFileToFileReader(file multipart.File, size int64) files.File {
	return files.NewReaderFile(&fileReader{file, size})
}

type fileReader struct {
	file multipart.File
	size int64
}

func (fr *fileReader) Read(p []byte) (n int, err error) {
	return fr.file.Read(p)
}

func (fr *fileReader) Close() error {
	return fr.file.Close()
}

func (fr *fileReader) Size() (int64, error) {
	return fr.size, nil
}
