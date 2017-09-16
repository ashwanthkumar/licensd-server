package packagemgrs

// Maven represents the maven's license format parsing the XML
type Maven struct{}

func (m *Maven) GetDependencies() *[]Dependency {
	return nil
}
