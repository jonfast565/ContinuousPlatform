package fileutil

import (
	"github.com/go-errors/errors"
	"github.com/jonfast565/continuous-platform/logging"
	"github.com/jonfast565/continuous-platform/pathutil"
	"reflect"
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
