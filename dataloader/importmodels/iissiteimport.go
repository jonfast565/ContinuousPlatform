package importmodels

type IisSiteImport struct {
	IisApplicationImport
	Applications []string
	Environments []EnvironmentImportPart
}
