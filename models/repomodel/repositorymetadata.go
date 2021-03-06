package repomodel

import (
	"fmt"
	"github.com/jonfast565/continuous-platform/models/filesysmodel"
	"github.com/jonfast565/continuous-platform/utilities/fileutil"
	"github.com/jonfast565/continuous-platform/utilities/pathutil"
	"github.com/jonfast565/continuous-platform/utilities/stringutil"
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
		if file.Type == filesysmodel.FolderType {
			folderChain := pp.GetAllItems()
			fg.Root.NewChildFolderChain(folderChain)
		} else if file.Type == filesysmodel.FileType {
			fileItem := pp.GetLastItem()
			stringFrag := pp.GetPreviousItems()
			newFolder := fg.Root.NewChildFolderChain(stringFrag)
			newFolder.NewChildFile(fileItem, []byte{})
		}
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
