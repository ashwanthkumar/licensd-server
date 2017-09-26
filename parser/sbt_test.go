package parser

import "testing"
import "github.com/stretchr/testify/assert"

func TestParseLineForAnEntryInSbtLicenseReportFile(t *testing.T) {
	sbt := SBTLicenseReport{}
	line := `Apache,"The Apache Software License, Version 2.0 (http://www.apache.org/licenses/LICENSE-2.0.txt)",org.apache.thrift # libthrift # 0.8.0,`
	expectedLicense := &License{
		Name: "The Apache Software License, Version 2.0",
		URL:  "http://www.apache.org/licenses/LICENSE-2.0.txt",
	}
	expectedDep := &Dependency{
		Name:     "org.apache.thrift:libthrift",
		Version:  "0.8.0",
		Licenses: []*License{expectedLicense},
	}
	dep, err := sbt.parseLine(line)

	assert.NoError(t, err)
	assert.EqualValues(t, expectedDep.Licenses, dep.Licenses)
}
