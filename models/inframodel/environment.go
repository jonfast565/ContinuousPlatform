package inframodel

type Environment struct {
	BusinessLine string
	Name         string
	Servers      []ServerType
}

func (e Environment) GetEnvironmentName() string {
	return e.BusinessLine + " " + e.Name
}
