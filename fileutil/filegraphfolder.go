package fileutil

import (
	"errors"
	"github.com/ahmetb/go-linq"
)

type FileGraphFolder struct {
	Name         string
	ChildFolders []*FileGraphFolder
	ChildFiles   []*FileGraphFile
	Parent       *FileGraphFolder
}

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

func (f FileGraphFolder) GetParent() *FileGraphItem {
	item := FileGraphItem(f.Parent)
	return &item
}

func (f FileGraphFolder) GetName() string {
	return f.Name
}

func (f FileGraphFolder) NavigateChildFolder(name string) (*FileGraphItem, error) {
	for _, folder := range f.ChildFolders {
		if folder.Name == name {
			item := FileGraphItem(folder)
			return &item, nil
		}
	}
	return nil, errors.New("Item '" + name + "' not found in '" + f.GetPathString() + "'")
}

func (f FileGraphFolder) NavigateChildFile(name string) (*FileGraphItem, error) {
	for _, file := range f.ChildFiles {
		if file.Name == name {
			item := FileGraphItem(file)
			return &item, nil
		}
	}
	return nil, errors.New("Item '" + name + "' not found in '" + f.GetPathString() + "'")
}

func (f FileGraphFolder) GetPathString() string {
	var result string
	currentNode := &f
	for {
		if result != "" {
			result = currentNode.Name + "/" + result
		} else {
			result = currentNode.Name
		}
		if f.Parent == nil {
			break
		} else {
			currentNode = f.Parent
		}
	}
	return result
}

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
		return f2.Name == f.Name
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

func (f *FileGraphFolder) NewChildFolderChain(pathFragments []string) *FileGraphFolder {
	currentFolder := f
	for _, fragment := range pathFragments {
		currentFolder = currentFolder.NewChildFolderNavigate(fragment)
	}
	return currentFolder
}

func (f *FileGraphFolder) NewChildFile(name string, contents []byte) *FileGraphFile {
	childFile := FileGraphFile{
		Name:     name,
		Parent:   f,
		Contents: contents,
	}
	f.ChildFiles = append(f.ChildFiles, &childFile)
	return &childFile
}
