package fileutil

import (
	"../pathutil"
	"github.com/go-errors/errors"
)

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

func (f *FileGraph) PrettifyPaths(paths []string) ([]string, error) {
	var result []string
	for _, path := range paths {
		itemRoot, err := f.GetItemByRootPath(path)
		if err != nil {
			return nil, err
		}
		result = append(result, (*itemRoot).GetPathString())
	}
	return result, nil
}

func (f *FileGraph) NewChildFolders(pathFragments []string) {
	f.Root.NewChildFolderChain(pathFragments)
}

func (f *FileGraph) GetItemByRootPath(basePath string) (*FileGraphItem, error) {
	item := FileGraphItem(f.Root)
	itemByRelativePath, err := GetItemByRelativePath(&item, basePath)
	return itemByRelativePath, err
}

func GetItemByRelativePath(item *FileGraphItem, basePath string) (*FileGraphItem, error) {
	pp := pathutil.NewPathParserFromString(basePath)
	currentNode := item
	for _, action := range *pp.ActionSeries {
		if action.Name == "." {
			continue
		} else if action.Name == ".." {
			parent := (*currentNode).GetParent()
			if parent == nil {
				return nil, errors.New("Navigating to parent of '" + (*currentNode).GetName() +
					"' goes off the root of the graph")
			}
			currentNode = parent
		} else {
			file, _ := (*currentNode).NavigateChildFile(action.Name)
			folder, _ := (*currentNode).NavigateChildFolder(action.Name)
			if file == nil && folder == nil {
				return nil, errors.New("Child '" + action.Name + "' of '" +
					(*currentNode).GetName() + "' does not exist")
			} else if file != nil && folder != nil {
				return nil, errors.New("Child '" + action.Name + "' of '" +
					(*currentNode).GetName() + "' cannot be both a file and folder.")
			} else if file != nil && folder == nil {
				currentNode = file
			} else if file == nil && folder != nil {
				currentNode = folder
			}
		}
	}
	return currentNode, nil
}

func (f *FileGraph) GetItemByRelativePath(basePath string, relativePath string) (*FileGraphItem, error) {
	item, err := f.GetItemByRootPath(basePath)
	if err != nil {
		return nil, err
	}
	item, err = GetItemByRelativePath(item, relativePath)
	if err != nil {
		return nil, err
	}
	return item, nil
}
