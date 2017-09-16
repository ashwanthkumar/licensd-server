package packagemgrs

type Dependency struct {
	Name     string
	Version  string
	Licenses []*License
}

type License struct {
	Name string
	URL  string
}

// PackageManager
type PackageManager interface {
	GetDependencies() *[]Dependency
}
