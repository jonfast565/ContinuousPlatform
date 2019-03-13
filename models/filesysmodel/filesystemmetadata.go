package filesysmodel

// object type of filesystem item
type FileSystemObjectType int

// file & folder types
const (
	FileType   FileSystemObjectType = 0
	FolderType FileSystemObjectType = 1
)

// Metadata describing the Path, Type, and Commit id of a file in the
type FileSystemMetadata struct {
	Path             string
	Type             FileSystemObjectType
	OptionalChangeId string
}
