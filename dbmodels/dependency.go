package dbmodels

// Dependency represents each dependency that the project has for a particular build
type Dependency struct {
	ID           int64
	Name         string
	Version      string
	CreatedAt    int64
	BuildMatrix  string
	Project      *Project
	BuildVersion int
	CIURL        string
	Licenses     []*License
}
