package repomodel

import "github.com/jonfast565/continuous-platform/models/filesysmodel"

type RepositoryFileMetadata struct {
	Repo   string
	Branch string
	Name   string
	File   filesysmodel.FileSystemMetadata
}

func NewRepositoryFileMetadata(
	repo string,
	branch string,
	name string,
	file filesysmodel.FileSystemMetadata) RepositoryFileMetadata {
	return RepositoryFileMetadata{Repo: repo, Branch: branch, Name: name, File: file}
}
