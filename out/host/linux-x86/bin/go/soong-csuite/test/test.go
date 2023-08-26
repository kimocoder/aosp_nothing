
package main

import (
	"io"

	"reflect"
	"regexp"
	"testing"
	"time"

	pkg "android/soong/csuite"
)

var t = []testing.InternalTest{

	{"TestBpContainsManifestThrowsError", pkg.TestBpContainsManifestThrowsError},

	{"TestBpContainsTestHostPropsThrowsError", pkg.TestBpContainsTestHostPropsThrowsError},

	{"TestBpExtraTemplateUnexpectedFileExtensionThrowsError", pkg.TestBpExtraTemplateUnexpectedFileExtensionThrowsError},

	{"TestBpMissingNameThrowsError", pkg.TestBpMissingNameThrowsError},

	{"TestBpMissingTemplatePathThrowsError", pkg.TestBpMissingTemplatePathThrowsError},

	{"TestBpTemplatePathUnexpectedFileExtensionThrowsError", pkg.TestBpTemplatePathUnexpectedFileExtensionThrowsError},

	{"TestBpValidExtraTemplateDoesNotThrowError", pkg.TestBpValidExtraTemplateDoesNotThrowError},

	{"TestExtraTemplateFileCopyRuleExists", pkg.TestExtraTemplateFileCopyRuleExists},

	{"TestGeneratedTestPlanContainsExtraTemplatePath", pkg.TestGeneratedTestPlanContainsExtraTemplatePath},

	{"TestGeneratedTestPlanContainsPlanInclude", pkg.TestGeneratedTestPlanContainsPlanInclude},

	{"TestGeneratedTestPlanContainsPlanName", pkg.TestGeneratedTestPlanContainsPlanName},

	{"TestGeneratedTestPlanContainsTemplatePath", pkg.TestGeneratedTestPlanContainsTemplatePath},

	{"TestGeneratedTestPlanDoesNotContainExtraTemplatePath", pkg.TestGeneratedTestPlanDoesNotContainExtraTemplatePath},

	{"TestPlanIncludeFileCopyRuleExists", pkg.TestPlanIncludeFileCopyRuleExists},

	{"TestTemplateFileCopyRuleExists", pkg.TestTemplateFileCopyRuleExists},

	{"TestValidBpMissingPlanIncludeDoesNotThrowError", pkg.TestValidBpMissingPlanIncludeDoesNotThrowError},

	{"TestValidBpMissingPlanIncludeGeneratesPlanXmlWithoutPlaceholders", pkg.TestValidBpMissingPlanIncludeGeneratesPlanXmlWithoutPlaceholders},

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
	return "android/soong/csuite"
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


	pkg.TestMain(m)

}
