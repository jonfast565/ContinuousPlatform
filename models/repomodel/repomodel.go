package repomodel

import (
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

type RepositoryFileMetadata struct {
	Name   string
	Repo   string
	Branch string
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
