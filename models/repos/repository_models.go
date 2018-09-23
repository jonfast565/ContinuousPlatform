package repos

import (
	"../../utilities"
	"../filesystem"
	"fmt"
)

type SourceControlProviderType int

const (
	AzureDevOps SourceControlProviderType = 0
	Github      SourceControlProviderType = 1
)

type RepositoryMetadata struct {
	Name     string
	Branch   string
	Url      string
	Metadata []filesystem.FileSystemMetadata
}

func (rm RepositoryMetadata) String() string {
	return fmt.Sprintf("Name: %s\nBranch: %s\nUrl: %s\n",
		rm.Name,
		rm.Branch,
		rm.Url)
}

type RepositoryMetadataGraphPair struct {
	Metadata RepositoryMetadata
	//Graph utilities.SourceControlGraph
}

type RepositoryFileMetadata struct {
	Name   string
	Repo   string
	Branch string
	File   filesystem.FileSystemMetadata
}

func MapToRepositoryMetadata(metadata filesystem.FileSystemMetadata,
	repositoryMetadata RepositoryMetadata) RepositoryFileMetadata {
	return RepositoryFileMetadata{
		Repo:   repositoryMetadata.Name,
		Branch: repositoryMetadata.Branch,
		File:   metadata,
		Name:   utilities.GetLastPathComponent(metadata.Path),
	}
}

type RepositoryPackage struct {
	Metadata []RepositoryMetadata
	Type     SourceControlProviderType
}

type RepositoryAmalgamation struct {
	Packages []RepositoryPackage
}
