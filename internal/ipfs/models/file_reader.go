package models

import "io"

type FileReader struct {
	reader io.Reader
	size   int64
}

func MultipartFileToFileReader(reader io.Reader, size int64) *FileReader {
	return &FileReader{reader: reader, size: size}
}

func (fr *FileReader) Read(p []byte) (n int, err error) {
	return fr.reader.Read(p)
}

func (fr *FileReader) Size() int64 {
	return fr.size
}
