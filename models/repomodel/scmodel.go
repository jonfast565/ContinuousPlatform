package repomodel

import "../miscmodel"

type SourceControlEndpoint interface {
	GetRepositories() ([]RepositoryMetadata, error)
	GetFile(file RepositoryFileMetadata) (miscmodel.FilePayload, error)
}
