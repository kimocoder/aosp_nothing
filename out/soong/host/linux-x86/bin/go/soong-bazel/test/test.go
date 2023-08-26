
package main

import (
	"io"

	"os"

	"reflect"
	"regexp"
	"testing"
	"time"

	pkg "android/soong/bazel"
)

var t = []testing.InternalTest{

	{"TestAqueryMultiArchGenrule", pkg.TestAqueryMultiArchGenrule},

	{"TestDeduplicateAxesFromBase", pkg.TestDeduplicateAxesFromBase},

	{"TestDepfiles", pkg.TestDepfiles},

	{"TestFirstUniqueBazelLabelList", pkg.TestFirstUniqueBazelLabelList},

	{"TestInvalidInputArtifactId", pkg.TestInvalidInputArtifactId},

	{"TestInvalidInputDepsetId", pkg.TestInvalidInputDepsetId},

	{"TestInvalidOutputId", pkg.TestInvalidOutputId},

	{"TestInvalidPathFragmentId", pkg.TestInvalidPathFragmentId},

	{"TestMiddlemenAction", pkg.TestMiddlemenAction},

	{"TestMultipleDepfiles", pkg.TestMultipleDepfiles},

	{"TestPartitionLabelListAttribute", pkg.TestPartitionLabelListAttribute},

	{"TestPythonZipperActionNoInput", pkg.TestPythonZipperActionNoInput},

	{"TestPythonZipperActionNoOutput", pkg.TestPythonZipperActionNoOutput},

	{"TestPythonZipperActionSuccess", pkg.TestPythonZipperActionSuccess},

	{"TestResolveExcludes", pkg.TestResolveExcludes},

	{"TestSimpleSymlink", pkg.TestSimpleSymlink},

	{"TestSubtractBazelLabelList", pkg.TestSubtractBazelLabelList},

	{"TestSubtractStrings", pkg.TestSubtractStrings},

	{"TestSymlinkMultipleInputs", pkg.TestSymlinkMultipleInputs},

	{"TestSymlinkMultipleOutputs", pkg.TestSymlinkMultipleOutputs},

	{"TestSymlinkQuotesPaths", pkg.TestSymlinkQuotesPaths},

	{"TestTemplateExpandActionNoOutput", pkg.TestTemplateExpandActionNoOutput},

	{"TestTemplateExpandActionSubstitutions", pkg.TestTemplateExpandActionSubstitutions},

	{"TestTransitiveInputDepsets", pkg.TestTransitiveInputDepsets},

	{"TestUniqueBazelLabels", pkg.TestUniqueBazelLabels},

	{"TestUniqueSortedBazelLabelList", pkg.TestUniqueSortedBazelLabelList},

}

var e = []testing.InternalExample{

}

var matchPat string
var matchRe *regexp.Regexp

type matchString struct{}

func MatchString(pat, str string) (result bool, err error) {
	if matchRe == nil || matchPat != pat {
		matchPat = pat
		matchRe, err = regexp.Compile(matchPat)
		if err != nil {
			return
		}
	}
	return matchRe.MatchString(str), nil
}

func (matchString) MatchString(pat, str string) (bool, error) {
	return MatchString(pat, str)
}

func (matchString) StartCPUProfile(w io.Writer) error {
	panic("shouldn't get here")
}

func (matchString) StopCPUProfile() {
}

func (matchString) WriteHeapProfile(w io.Writer) error {
    panic("shouldn't get here")
}

func (matchString) WriteProfileTo(string, io.Writer, int) error {
    panic("shouldn't get here")
}

func (matchString) ImportPath() string {
	return "android/soong/bazel"
}

func (matchString) StartTestLog(io.Writer) {
	panic("shouldn't get here")
}

func (matchString) StopTestLog() error {
	panic("shouldn't get here")
}

func (matchString) SetPanicOnExit0(bool) {
	panic("shouldn't get here")
}

func (matchString) CoordinateFuzzing(time.Duration, int64, time.Duration, int64, int, []corpusEntry, []reflect.Type, string, string) error {
	panic("shouldn't get here")
}

func (matchString) RunFuzzWorker(func(corpusEntry) error) error {
	panic("shouldn't get here")
}

func (matchString) ReadCorpus(string, []reflect.Type) ([]corpusEntry, error) {
	panic("shouldn't get here")
}

func (matchString) CheckCorpus([]interface{}, []reflect.Type) error {
	panic("shouldn't get here")
}

func (matchString) ResetCoverage() {
	panic("shouldn't get here")
}

func (matchString) SnapshotCoverage() {
	panic("shouldn't get here")
}

type corpusEntry = struct {
	Parent     string
	Path       string
	Data       []byte
	Values     []interface{}
	Generation int
	IsSeed     bool
}

func main() {

	m := testing.MainStart(matchString{}, t, nil, nil, e)


	os.Exit(m.Run())

}
