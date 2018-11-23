package importmodels

type IisSiteRecord struct {
	IisApplicationRecord
	Applications []string
	Environments []EnvironmentRecordPart
}
