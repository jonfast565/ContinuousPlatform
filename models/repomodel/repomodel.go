package repomodel

import (
	"../../fileutil"
	"../../pathutil"
	"../../stringutil"
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

func (rm RepositoryMetadata) BuildGraph() *fileutil.FileGraph {
	fg := fileutil.NewFileGraph()
	for _, file := range rm.Files {
		pp := pathutil.NewPathParserFromString(file.Path)
		fileItem := pp.GetLastItem()
		stringFrag := pp.GetPreviousItems()
		newFolder := fg.Root.NewChildFolderChain(stringFrag)
		newFolder.NewChildFile(fileItem, []byte{})
	}
	return fg
}

func (rm RepositoryMetadata) GetMatchingFiles(regexStrings []string) ([]string, error) {
	regexes, err := stringutil.CompileStringsAsRegexes(regexStrings)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, file := range rm.Files {
		if file.Type == filesysmodel.FileType {
			pp := pathutil.NewPathParserFromString(file.Path)
			lastItem := pp.GetLastItem()
			if stringutil.StringMatchesOneOf(lastItem, regexes) {
				result = append(result, file.Path)
			}
		}
	}
	return result, nil
}

func (rm RepositoryMetadata) String() string {
	return fmt.Sprintf("Repo: %s\nBranch: %s\nUrl: %s\n",
		rm.Name,
		rm.Branch,
		rm.OptionalUrl)
}

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
