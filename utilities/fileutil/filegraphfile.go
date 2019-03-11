package fileutil

import "errors"

// A file, implementation of an item
type FileGraphFile struct {
	Name     string
	Contents []byte
	Parent   *FileGraphFolder
}

// Get the parent
func (f FileGraphFile) NavigateParent() (*FileGraphItem, error) {
	if f.Parent == nil {
		return nil, errors.New("Parent does not exist for '" + f.GetPathString() + "'")
	}
	item := FileGraphItem(f.Parent)
	return &item, nil
}

// Get the name
func (f FileGraphFile) GetName() string {
	return f.Name
}

// Cannot navigate to a child folder of a file
func (f FileGraphFile) NavigateChildFolder(name string) (*FileGraphItem, error) {
	return nil, errors.New(f.Name + " is not a folder. Cannot navigate to it's children")
}

// Cannot navigate to a child file of a file
func (f FileGraphFile) NavigateChildFile(name string) (*FileGraphItem, error) {
	return nil, errors.New(f.Name + " is not a folder. Cannot navigate to it's children")
}

// Get the path string of this item
func (f FileGraphFile) GetPathString() string {
	return f.Parent.GetPathString() + "/" + f.Name
}
