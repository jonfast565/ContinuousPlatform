package inframodel

type Environment struct {
	BusinessLine string
	Name         string
	Servers      []ServerType
}

func (e Environment) GetEnvironmentName() string {
	return e.BusinessLine + " " + e.Name
}

type EnvironmentPart struct {
	BusinessLine string
	Name         string
}

func (e EnvironmentPart) GetEnvironmentName() string {
	return e.BusinessLine + " " + e.Name
}
