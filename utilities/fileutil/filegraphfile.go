package fileutil

import "errors"

type FileGraphFile struct {
	Name     string
	Contents []byte
	Parent   *FileGraphFolder
}

func (f FileGraphFile) NavigateParent() (*FileGraphItem, error) {
	if f.Parent == nil {
		return nil, errors.New("Parent does not exist for '" + f.GetPathString() + "'")
	}
	item := FileGraphItem(f.Parent)
	return &item, nil
}

func (f FileGraphFile) GetName() string {
	return f.Name
}

func (f FileGraphFile) NavigateChildFolder(name string) (*FileGraphItem, error) {
	return nil, errors.New(f.Name + " is not a folder. Cannot navigate to it's children")
}

func (f FileGraphFile) NavigateChildFile(name string) (*FileGraphItem, error) {
	return nil, errors.New(f.Name + " is not a folder. Cannot navigate to it's children")
}

func (f FileGraphFile) GetPathString() string {
	return f.Parent.GetPathString() + "/" + f.Name
}
