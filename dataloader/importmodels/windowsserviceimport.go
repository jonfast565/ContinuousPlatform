package importmodels

type WindowsServiceImport struct {
	Name                      string
	BinaryPath                string
	BinaryExecutableName      string
	BinaryExecutableArguments string
	Environments              []EnvironmentImportPart
}
