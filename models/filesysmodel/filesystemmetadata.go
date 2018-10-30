package filesysmodel

type FileSystemObjectType int

const (
	FileType   FileSystemObjectType = 0
	FolderType FileSystemObjectType = 1
)

type FileSystemMetadata struct {
	Path             string
	Type             FileSystemObjectType
	OptionalChangeId string
}
