package parser

import "testing"
import "github.com/stretchr/testify/assert"

func TestParseLineForAnEntry(t *testing.T) {
	sbt := LicenseFinder{}
	line := `cascading:cascading-core, 2.5.5, "Apache 2.0"`
	expectedLicense := &License{Name: "Apache 2.0"}
	expectedDep := &Dependency{
		Name:     "cascading:cascading-core",
		Version:  "2.5.5",
		Licenses: []*License{expectedLicense},
	}
	dep, err := sbt.parseLine(line)

	assert.NoError(t, err)
	assert.EqualValues(t, expectedDep.Licenses, dep.Licenses)
}
