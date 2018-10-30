package repomodel

type RepositoryPackage struct {
	Metadata []RepositoryMetadata
	Type     SourceControlProviderType
}

func NewRepositoryPackage() *RepositoryPackage {
	return &RepositoryPackage{Metadata: make([]RepositoryMetadata, 0), Type: AzureDevOps}
}
