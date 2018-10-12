package repos

import "../payload"

type SourceControlEndpoint interface {
	GetRepositories() ([]RepositoryMetadata, error)
	GetFile(file RepositoryFileMetadata) (payload.FilePayload, error)
}
