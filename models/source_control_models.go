package models

type SourceControlEndpoint interface {
	GetRepositories() ([]RepositoryMetadata, error)
	GetFile(file RepositoryFileMetadata) (FilePayload, error)
}
