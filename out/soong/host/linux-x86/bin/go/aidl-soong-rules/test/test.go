
package main

import (
	"io"

	"reflect"
	"regexp"
	"testing"
	"time"

	pkg "android/soong/aidl"
)

var t = []testing.InternalTest{

	{"TestAidlFlags", pkg.TestAidlFlags},

	{"TestAidlImportFlagsForImportedModules", pkg.TestAidlImportFlagsForImportedModules},

	{"TestAidlImportFlagsForUnstable", pkg.TestAidlImportFlagsForUnstable},

	{"TestAidlModuleJavaSdkVersionDeterminesMinSdkVersion", pkg.TestAidlModuleJavaSdkVersionDeterminesMinSdkVersion},

	{"TestAidlModuleNameContainsVersion", pkg.TestAidlModuleNameContainsVersion},

	{"TestAidlPreprocess", pkg.TestAidlPreprocess},

	{"TestCreatesModulesWithFrozenVersions", pkg.TestCreatesModulesWithFrozenVersions},

	{"TestCreatesModulesWithNoVersions", pkg.TestCreatesModulesWithNoVersions},

	{"TestDuplicateInterfacesWithTheSameNameInDifferentSoongNamespaces", pkg.TestDuplicateInterfacesWithTheSameNameInDifferentSoongNamespaces},

	{"TestDuplicatedVersions", pkg.TestDuplicatedVersions},

	{"TestErrorsWithDuplicateVersions", pkg.TestErrorsWithDuplicateVersions},

	{"TestErrorsWithNonIntegerVersions", pkg.TestErrorsWithNonIntegerVersions},

	{"TestErrorsWithNonPositiveVersions", pkg.TestErrorsWithNonPositiveVersions},

	{"TestErrorsWithUnsortedVersions", pkg.TestErrorsWithUnsortedVersions},

	{"TestExplicitAidlModuleImport", pkg.TestExplicitAidlModuleImport},

	{"TestFreezeApiDeps", pkg.TestFreezeApiDeps},

	{"TestImportInRelease", pkg.TestImportInRelease},

	{"TestImports", pkg.TestImports},

	{"TestNativeOutputIsAlwaysVersioned", pkg.TestNativeOutputIsAlwaysVersioned},

	{"TestNonVersionedModuleOwnedByOtherUsageInRelease", pkg.TestNonVersionedModuleOwnedByOtherUsageInRelease},

	{"TestNonVersionedModuleOwnedByTestUsageInRelease", pkg.TestNonVersionedModuleOwnedByTestUsageInRelease},

	{"TestNonVersionedModuleUsageInRelease", pkg.TestNonVersionedModuleUsageInRelease},

	{"TestRecoveryAvailable", pkg.TestRecoveryAvailable},

	{"TestRustDuplicateNames", pkg.TestRustDuplicateNames},

	{"TestSupportsGenruleAndFilegroup", pkg.TestSupportsGenruleAndFilegroup},

	{"TestUnstableChecksForAidlInterfacesInDifferentNamespaces", pkg.TestUnstableChecksForAidlInterfacesInDifferentNamespaces},

	{"TestUnstableModules", pkg.TestUnstableModules},

	{"TestUnstableVersionUsageInRelease", pkg.TestUnstableVersionUsageInRelease},

	{"TestUnstableVersionedModuleOwnedByOtherUsageInRelease", pkg.TestUnstableVersionedModuleOwnedByOtherUsageInRelease},

	{"TestUnstableVersionedModuleOwnedByTestUsageInRelease", pkg.TestUnstableVersionedModuleOwnedByTestUsageInRelease},

	{"TestUnstableVersionedModuleUsageInRelease", pkg.TestUnstableVersionedModuleUsageInRelease},

	{"TestUnstableVndkModule", pkg.TestUnstableVndkModule},

	{"TestUseVersionedPreprocessedWhenImporotedWithVersions", pkg.TestUseVersionedPreprocessedWhenImporotedWithVersions},

	{"TestUsingUnstableVersionIndirectlyInRelease", pkg.TestUsingUnstableVersionIndirectlyInRelease},

	{"TestVersionsWithInfo", pkg.TestVersionsWithInfo},

	{"TestVersionsWithInfoAndVersions", pkg.TestVersionsWithInfoAndVersions},

	{"TestVersionsWithInfoImport", pkg.TestVersionsWithInfoImport},

	{"TestVintfWithoutVersionInRelease", pkg.TestVintfWithoutVersionInRelease},

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
	return "android/soong/aidl"
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
