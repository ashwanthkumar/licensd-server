package parser

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"strings"
)

// LicenseFinder parses the CSV files generated by
// https://github.com/pivotal/LicenseFinder
type LicenseFinder struct {
}

// NewLicenseFinderParser returns a new instance of LicenseFinder Parser
func NewLicenseFinderParser() *LicenseFinder {
	return &LicenseFinder{}
}

// Parse parses the CSV file generated by the plugin into LicenseD format
func (lf *LicenseFinder) Parse(scanner *bufio.Scanner) ([]*Dependency, error) {
	var deps []*Dependency
	for scanner.Scan() {
		line := scanner.Text()
		dep, err := lf.parseLine(line)
		if err == nil {
			deps = append(deps, dep)
		} else {
			fmt.Printf("%v\n", err)
		}
	}
	return deps, nil
}

// line -> group:name, version, license
// line -> name, version, license
func (lf *LicenseFinder) parseLine(line string) (*Dependency, error) {
	reader := csv.NewReader(strings.NewReader(line))
	reader.TrimLeadingSpace = true
	reader.LazyQuotes = true
	row, err := reader.Read()
	if err != nil {
		return nil, err
	}
	license := &License{Name: row[2]}
	dep := &Dependency{
		Name:     row[0],
		Version:  row[1],
		Licenses: []*License{license},
	}
	return dep, nil
}
