package inframodel

type Environment struct {
	BusinessLine string
	Name         string
	Servers      []ServerType
}

func (e Environment) GetEnvironmentName() string {
	return e.BusinessLine + " " + e.Name
}

func (e Environment) NameMatch(e2 Environment) bool {
	return e.BusinessLine == e2.BusinessLine &&
		e.Name == e2.Name
}

type EnvironmentList []Environment

func (el EnvironmentList) HasMatch(e Environment) bool {
	for _, environment := range el {
		if environment.NameMatch(e) {
			return true
		}
	}
	return false
}

type EnvironmentPart struct {
	BusinessLine string
	Name         string
}

func (e EnvironmentPart) GetEnvironmentName() string {
	return e.BusinessLine + " " + e.Name
}

func (e EnvironmentPart) NameMatch(e2 EnvironmentPart) bool {
	return e.BusinessLine == e2.BusinessLine &&
		e.Name == e2.Name
}

func (e EnvironmentPart) NameMatchEnv(e2 Environment) bool {
	return e.BusinessLine == e2.BusinessLine &&
		e.Name == e2.Name
}

type EnvironmentPartList []EnvironmentPart

func (el EnvironmentPartList) HasMatch(e EnvironmentPart) bool {
	for _, environment := range el {
		if environment.NameMatch(e) {
			return true
		}
	}
	return false
}
