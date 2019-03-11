// A utility package for working with files and paths related to them
package fileutil

import (
	"github.com/go-errors/errors"
	"github.com/jonfast565/continuous-platform/utilities/logging"
	"github.com/jonfast565/continuous-platform/utilities/pathutil"
	"reflect"
)

// A file graph struct, with root folder child
type FileGraph struct {
	Root FileGraphFolder
}

// Creates a new file graph
func NewFileGraph() *FileGraph {
	return &FileGraph{
		Root: FileGraphFolder{
			Name:         ".",
			ChildFiles:   []*FileGraphFile{},
			ChildFolders: []*FileGraphFolder{},
			Parent:       nil}}
}

// Validate a path from the root folder, determining if it is reachable or not
func (f *FileGraph) ValidatePathsFromRoot(paths []string, soft bool) ([]string, error) {
	var result []string
	for _, path := range paths {
		itemRoot, err := f.GetItemByRootPath(path)
		if err != nil {
			if !soft {
				return nil, err
			} else {
				logging.LogSoftError("Path validation failed:", err)
				continue
			}
		}
		result = append(result, (*itemRoot).GetPathString())
	}
	return result, nil
}

// Validate a path from another path that comes from the root folder
func (f *FileGraph) ValidatePathsFromBasePath(basePath string, paths []string, soft bool) ([]string, error) {
	var result []string
	for _, path := range paths {
		itemPath, err := f.GetItemByRelativePath(basePath, path)
		if err != nil {
			if !soft {
				return nil, err
			} else {
				logging.LogSoftError("Path validation failed:", err)
				continue
			}
		}
		result = append(result, (*itemPath).GetPathString())
	}
	return result, nil
}

// Gets the parent path of a given path
func (f *FileGraph) GetParentPath(path string) (*string, error) {
	itemPath, err := f.GetItemByRootPath(path)
	if err != nil {
		return nil, err
	}
	parent, err := (*itemPath).NavigateParent()
	if err != nil {
		return nil, err
	}
	result := (*parent).GetPathString()
	return &result, nil
}

// Creates a new chain of child folders from a given array of path fragments
func (f *FileGraph) NewChildFolders(pathFragments []string) {
	f.Root.NewChildFolderChain(pathFragments)
}

// Gets a FileGraphItem by its path from the root
func (f *FileGraph) GetItemByRootPath(basePath string) (*FileGraphItem, error) {
	item := FileGraphItem(f.Root)
	itemByRelativePath, err := GetItemByRelativePath(&item, basePath)
	return itemByRelativePath, err
}

// Gets a FileGraphItem by its relative path from another FileGraphItem
func GetItemByRelativePath(item *FileGraphItem, basePath string) (*FileGraphItem, error) {
	pp := pathutil.NewPathParserFromString(basePath)
	currentNode := item
	for _, action := range *pp.ActionSeries {
		if action.Name == "." {
			continue
		} else if action.Name == ".." {
			parent, err := getParentOfCurrentNode(currentNode)
			if err != nil {
				return nil, errors.New("Navigating to parent of '" + (*currentNode).GetName() +
					"' goes off the root of the graph")
			}
			currentNode = parent
		} else {
			file, folder := navigateChildFolderFile(currentNode, action)
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

func navigateChildFolderFile(currentNode *FileGraphItem, action pathutil.PathAction) (*FileGraphItem, *FileGraphItem) {
	file, _ := (*currentNode).NavigateChildFile(action.Name)
	folder, _ := (*currentNode).NavigateChildFolder(action.Name)
	return file, folder
}

func getParentOfCurrentNode(currentNode *FileGraphItem) (*FileGraphItem, error) {
	if currentNode == nil {
		return nil, errors.New("Current node cannot be null")
	}
	nodeForParent := *currentNode
	return nodeForParent.NavigateParent()
}

// Adds a folder item to the graph by relative path from an existing item (could be the root)
func AddFolderByRelativePath(item *FileGraphItem, basePath string) (*FileGraphItem, error) {
	pp := pathutil.NewPathParserFromString(basePath)
	currentNode := item
	for _, action := range *pp.ActionSeries {
		if action.Name == "." {
			continue
		} else if action.Name == ".." {
			parent, err := getParentOfCurrentNode(currentNode)
			if err != nil {
				return nil, errors.New("Navigating to parent of '" + (*currentNode).GetName() +
					"' goes off the root of the graph")
			}
			currentNode = parent
		} else {
			_, folder := navigateChildFolderFile(currentNode, action)
			if folder == nil {
				parentNode := *currentNode
				// TODO: This is a hack, needs to be fixed
				if reflect.TypeOf(parentNode).Kind() == reflect.Ptr {
					parentFolder := parentNode.(*FileGraphFolder)
					parentFolder.NewChildFolder(action.Name)
					item := FileGraphItem(parentFolder)
					currentNode = &item
				} else if reflect.TypeOf(parentNode).Kind() == reflect.Struct {
					parentFolder := parentNode.(FileGraphFolder)
					parentFolder.NewChildFolder(action.Name)
					item := FileGraphItem(parentFolder)
					currentNode = &item
				}
			} else {
				currentNode = folder
			}
		}
	}
	return currentNode, nil
}

// Gets an item by first getting the root, then getting the item relative to the root
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

// Adds a folder relative to a root path
func (f *FileGraph) AddFolderByRelativePath(basePath string, relativePath string) (*FileGraphItem, error) {
	item, err := f.GetItemByRootPath(basePath)
	if err != nil {
		return nil, err
	}
	item, err = AddFolderByRelativePath(item, relativePath)
	if err != nil {
		return nil, err
	}
	return item, nil
}
