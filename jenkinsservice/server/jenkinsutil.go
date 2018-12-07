package server

import "bytes"

func trapBOM(fileBytes []byte) []byte {
	trimmedBytes := bytes.Trim(fileBytes, "\xef\xbb\xbf")
	return trimmedBytes
}

func jenkinsByteReaderFromString(requestContents string) *bytes.Reader {
	byteBuf := []byte(requestContents)
	bomLessByteBuf := trapBOM(byteBuf)
	readerBuf := bytes.NewReader(bomLessByteBuf)
	return readerBuf
}
