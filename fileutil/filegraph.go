package fileutil

import "github.com/ahmetb/go-linq"

type FileGraph struct {
	Root FileGraphFolder
}

func NewFileGraph() *FileGraph {
	return &FileGraph{
		Root: FileGraphFolder{
			Name:         ".",
			ChildFiles:   []*FileGraphFile{},
			ChildFolders: []*FileGraphFolder{},
			Parent:       nil}}
}

func (f *FileGraph) NewChildFolders(pathFragments []string) {
	f.Root.NewChildFolderChain(pathFragments)
}

type FileGraphFile struct {
	Name     string
	Contents []byte
	Parent   *FileGraphFolder
}

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

func (f *FileGraphFolder) NewChildFolderNavigate(name string) *FileGraphFolder {
	// handle the two edge cases gloriously (not)
	if name == "." {
		return f
	}
	if name == ".." {
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
