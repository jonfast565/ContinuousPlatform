package fileutil

type FileGraphItem interface {
	GetParent() *FileGraphItem
	GetName() string
	NavigateChildFolder(name string) (*FileGraphItem, error)
	NavigateChildFile(name string) (*FileGraphItem, error)
	GetPathString() string
}
