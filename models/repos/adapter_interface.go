package repos

import "../web"

type SourceControlEndpoint interface {
	GetRepositories() ([]RepositoryMetadata, error)
	GetFile(file RepositoryFileMetadata) (web.FilePayload, error)
}
