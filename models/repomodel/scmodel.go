package repomodel

import "github.com/jonfast565/continuous-platform/models/miscmodel"

type SourceControlEndpoint interface {
	GetRepositories() ([]RepositoryMetadata, error)
	GetFile(file RepositoryFileMetadata) (miscmodel.FilePayload, error)
}
