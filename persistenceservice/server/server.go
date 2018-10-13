package server

type PersistenceServiceConfiguration struct {
	Port             int    `json:"port"`
	ConnectionString string `json:"connectionString"`
}

type PersistenceServiceEndpoint struct {
	Configuration PersistenceServiceConfiguration
}

func NewPersistenceServiceEndpoint(configuration PersistenceServiceConfiguration) *PersistenceServiceEndpoint {
	result := new(PersistenceServiceEndpoint)
	result.Configuration = configuration
	return result
}

func (p *PersistenceServiceEndpoint) SetKeyValueCache() {

}

func (p *PersistenceServiceEndpoint) GetKeyValueCache() {

}

func (p *PersistenceServiceEndpoint) GetInfrastructureMetadata() {

}

func (p *PersistenceServiceEndpoint) SetLogRecord() {

}
