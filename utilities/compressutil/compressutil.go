// a small package for providing compression functionality
package compressutil

import (
	"bytes"
	"compress/gzip"
	"io"
)

// Uncompress a byte array and return an error if failed
func Uncompress(inBytes []byte) ([]byte, error) {
	b := bytes.NewBuffer(inBytes)

	var reader io.Reader
	reader, err := gzip.NewReader(b)
	if err != nil {
		return nil, err
	}

	var bufferedResult bytes.Buffer
	_, err = bufferedResult.ReadFrom(reader)
	if err != nil {
		return nil, err
	}

	uncompressedData := bufferedResult.Bytes()
	return uncompressedData, nil
}

// Compress a byte array and return an error if failed
func Compress(inBytes []byte) ([]byte, error) {
	var b bytes.Buffer
	writer := gzip.NewWriter(&b)

	_, err := writer.Write(inBytes)
	if err != nil {
		return nil, err
	}

	if err = writer.Flush(); err != nil {
		return nil, err
	}

	if err = writer.Close(); err != nil {
		return nil, err
	}

	compressedData := b.Bytes()
	return compressedData, nil
}
