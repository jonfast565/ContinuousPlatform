package persistmodel

type KeyValueRequest struct {
	Key   string
	Value []byte
}

type KeyRequest struct {
	Key string
}

type KeyValueResult struct {
	Value []byte
}
