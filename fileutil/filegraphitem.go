package fileutil

type FileGraphItem interface {
	GetName() string
	NavigateParent() (*FileGraphItem, error)
	NavigateChildFolder(name string) (*FileGraphItem, error)
	NavigateChildFile(name string) (*FileGraphItem, error)
	GetPathString() string
}
