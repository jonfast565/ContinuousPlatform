package fileutil

import (
	"errors"
	"github.com/ahmetb/go-linq"
)

// A folder, implementation of an item
type FileGraphFolder struct {
	Name         string
	ChildFolders []*FileGraphFolder
	ChildFiles   []*FileGraphFile
	Parent       *FileGraphFolder
}

// Creates a new child folder
func (f *FileGraphFolder) NewChildFolder(name string) *FileGraphFolder {
	childFolder := FileGraphFolder{
		Name:         name,
		Parent:       f,
		ChildFolders: make([]*FileGraphFolder, 0),
		ChildFiles:   make([]*FileGraphFile, 0),
	}
	f.ChildFolders = append(f.ChildFolders, &childFolder)
	return &childFolder
}

// Navigates to a parent
func (f FileGraphFolder) NavigateParent() (*FileGraphItem, error) {
	if f.Parent == nil {
		return nil, errors.New("Parent does not exist for '" + f.GetPathString() + "'")
	}
	item := FileGraphItem(f.Parent)
	return &item, nil
}

// Get a name from the folder
func (f FileGraphFolder) GetName() string {
	return f.Name
}

// Navigate to a child folder
func (f FileGraphFolder) NavigateChildFolder(name string) (*FileGraphItem, error) {
	for _, folder := range f.ChildFolders {
		if folder.Name == name {
			folderItem := *folder
			item := FileGraphItem(folderItem)
			return &item, nil
		}
	}
	return nil, errors.New("Item '" + name + "' not found in '" + f.GetPathString() + "'")
}

// Navigate to a child file of a folder
func (f FileGraphFolder) NavigateChildFile(name string) (*FileGraphItem, error) {
	for _, file := range f.ChildFiles {
		if file.Name == name {
			fileItem := *file
			item := FileGraphItem(fileItem)
			return &item, nil
		}
	}
	return nil, errors.New("Item '" + name + "' not found in '" + f.GetPathString() + "'")
}

// Get a path string
func (f FileGraphFolder) GetPathString() string {
	var result string
	currentNode := &f
	for {
		// obvs, but this edge case may not exist
		if currentNode == nil {
			break
		}
		// result can't be nil either
		if result != "" {
			result = currentNode.Name + "/" + result
		} else {
			result = currentNode.Name
		}
		// parent can't be nil, or else string will run forever
		if currentNode.Parent == nil {
			break
		} else {
			currentNode = currentNode.Parent
		}
	}
	return result
}

// Create a new child folder and navigate to it automatically
func (f *FileGraphFolder) NewChildFolderNavigate(name string) *FileGraphFolder {
	// handle the two edge cases gloriously (not)
	if name == "." {
		return f
	}
	if name == ".." {
		if f.Parent == nil {
			panic("Parent of node '" + f.Name + "' does not exist")
		}
		return f.Parent
	}

	childFolderFilterFunc := func(f2 *FileGraphFolder) bool {
		return f2.Name == name
	}

	existingChildFolder := linq.From(f.ChildFolders).
		FirstWithT(childFolderFilterFunc)

	if existingChildFolder != nil {
		existingChildFolderImpl := existingChildFolder.(*FileGraphFolder)
		return existingChildFolderImpl
	}

	childFolder := f.NewChildFolder(name)
	return childFolder
}

// Create a new chain of child folders and navigate to the last folder in the chain
func (f *FileGraphFolder) NewChildFolderChain(pathFragments []string) *FileGraphFolder {
	currentFolder := f
	for _, fragment := range pathFragments {
		currentFolder = currentFolder.NewChildFolderNavigate(fragment)
	}
	return currentFolder
}

// Create a new child file
func (f *FileGraphFolder) NewChildFile(name string, contents []byte) *FileGraphFile {
	childFile := FileGraphFile{
		Name:     name,
		Parent:   f,
		Contents: contents,
	}
	f.ChildFiles = append(f.ChildFiles, &childFile)
	return &childFile
}
