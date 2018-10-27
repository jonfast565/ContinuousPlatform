package repomodel

import (
	"../../pathutil"
	"../filesysmodel"
	"fmt"
)

type SourceControlProviderType int

const (
	AzureDevOps SourceControlProviderType = 0
	Github      SourceControlProviderType = 1
)

type RepositoryMetadata struct {
	Name        string
	Branch      string
	OptionalUrl string
	Files       []filesysmodel.FileSystemMetadata
}

func (rm RepositoryMetadata) String() string {
	return fmt.Sprintf("Repo: %s\nBranch: %s\nUrl: %s\n",
		rm.Name,
		rm.Branch,
		rm.OptionalUrl)
}

type RepositoryMetadataGraphPair struct {
	Metadata RepositoryMetadata
	//Graph utilities.SourceControlGraph
}

type RepositoryFileMetadata struct {
	Name   string
	Repo   string
	Branch string
	File   filesysmodel.FileSystemMetadata
}

func NewRepositoryFileMetadata(
	Name string,
	Repo string,
	Branch string) RepositoryFileMetadata {
	return RepositoryFileMetadata{
		Name:   Name,
		Repo:   Repo,
		Branch: Branch,
	}
}

func (rfm RepositoryFileMetadata) String() string {
	return fmt.Sprintf("Repo: %s\nBranch: %s\nFilePath: %s\n",
		rfm.Name,
		rfm.Branch,
		rfm.File.Path)
}

func MapToRepositoryMetadata(metadata filesysmodel.FileSystemMetadata,
	repositoryMetadata RepositoryMetadata) RepositoryFileMetadata {
	return RepositoryFileMetadata{
		Repo:   repositoryMetadata.Name,
		Branch: repositoryMetadata.Branch,
		File:   metadata,
		Name:   pathutil.GetLastPathComponent(metadata.Path),
	}
}

type RepositoryPackage struct {
	Metadata []RepositoryMetadata
	Type     SourceControlProviderType
}

func NewRepositoryPackage() *RepositoryPackage {
	return &RepositoryPackage{Metadata: make([]RepositoryMetadata, 0), Type: AzureDevOps}
}

type RepositoryAmalgamation struct {
	Packages []RepositoryPackage
}
