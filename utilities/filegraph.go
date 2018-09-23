package utilities

import (
	"container/list"
)

type FileGraph struct {
	Root FileGraphItem
}

type FileGraphItem interface {
	getName() string
	getParent() FileGraphItem
	setName(name string)
	setParent(item FileGraphItem)
}

type FileGraphFile struct {
	Name string
}

type FileGraphFolder struct {
	Name     string
	Children *list.List
}

func (folder *FileGraphFolder) New(name string) {
	folder.Name = name
	folder.Children = list.New()
}

type SourceControlItemType int

const (
	FolderItem SourceControlItemType = 0
	FileItem   SourceControlItemType = 1
)
