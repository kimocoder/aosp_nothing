
package main

import (
	"io"

	"reflect"
	"regexp"
	"testing"
	"time"

	pkg "android/soong/rust"
)

var t = []testing.InternalTest{

	{"TestAndroidDylib", pkg.TestAndroidDylib},

	{"TestAutoDeps", pkg.TestAutoDeps},

	{"TestBinaryFlags", pkg.TestBinaryFlags},

	{"TestBinaryLinkage", pkg.TestBinaryLinkage},

	{"TestBinaryPreferRlib", pkg.TestBinaryPreferRlib},

	{"TestBindgenDisallowedFlags", pkg.TestBindgenDisallowedFlags},

	{"TestBootstrap", pkg.TestBootstrap},

	{"TestCargoCompat", pkg.TestCargoCompat},

	{"TestCfgsToFlags", pkg.TestCfgsToFlags},

	{"TestClippy", pkg.TestClippy},

	{"TestCoverageDeps", pkg.TestCoverageDeps},

	{"TestCoverageFlags", pkg.TestCoverageFlags},

	{"TestDataLibs", pkg.TestDataLibs},

	{"TestDataLibsRelativeInstallPath", pkg.TestDataLibsRelativeInstallPath},

	{"TestDepsTracking", pkg.TestDepsTracking},

	{"TestDylibPreferDynamic", pkg.TestDylibPreferDynamic},

	{"TestEnforceMissingSourceFiles", pkg.TestEnforceMissingSourceFiles},

	{"TestEnforceSingleSourceFile", pkg.TestEnforceSingleSourceFile},

	{"TestFeaturesToFlags", pkg.TestFeaturesToFlags},

	{"TestForbiddenVendorLinkage", pkg.TestForbiddenVendorLinkage},

	{"TestHostToolPath", pkg.TestHostToolPath},

	{"TestImageVndkCfgFlag", pkg.TestImageVndkCfgFlag},

	{"TestInstallDir", pkg.TestInstallDir},

	{"TestLibrarySizes", pkg.TestLibrarySizes},

	{"TestLibraryVariants", pkg.TestLibraryVariants},

	{"TestLibstdLinkage", pkg.TestLibstdLinkage},

	{"TestLinkObjects", pkg.TestLinkObjects},

	{"TestLinkPathFromFilePath", pkg.TestLinkPathFromFilePath},

	{"TestLints", pkg.TestLints},

	{"TestManualLinkageRejection", pkg.TestManualLinkageRejection},

	{"TestMultilib", pkg.TestMultilib},

	{"TestNoStdlibs", pkg.TestNoStdlibs},

	{"TestProcMacroDeviceDeps", pkg.TestProcMacroDeviceDeps},

	{"TestProjectJsonBinary", pkg.TestProjectJsonBinary},

	{"TestProjectJsonBindGen", pkg.TestProjectJsonBindGen},

	{"TestProjectJsonDep", pkg.TestProjectJsonDep},

	{"TestProjectJsonFeature", pkg.TestProjectJsonFeature},

	{"TestProjectJsonMultiVersion", pkg.TestProjectJsonMultiVersion},

	{"TestProjectJsonProcMacroDep", pkg.TestProjectJsonProcMacroDep},

	{"TestRecoverySnapshotCapture", pkg.TestRecoverySnapshotCapture},

	{"TestRecoverySnapshotDirected", pkg.TestRecoverySnapshotDirected},

	{"TestRecoverySnapshotExclude", pkg.TestRecoverySnapshotExclude},

	{"TestRustBenchmark", pkg.TestRustBenchmark},

	{"TestRustBenchmarkLinkage", pkg.TestRustBenchmarkLinkage},

	{"TestRustBindgen", pkg.TestRustBindgen},

	{"TestRustBindgenCustomBindgen", pkg.TestRustBindgenCustomBindgen},

	{"TestRustBindgenStdVersions", pkg.TestRustBindgenStdVersions},

	{"TestRustFuzz", pkg.TestRustFuzz},

	{"TestRustGrpc", pkg.TestRustGrpc},

	{"TestRustProcMacro", pkg.TestRustProcMacro},

	{"TestRustProtoErrors", pkg.TestRustProtoErrors},

	{"TestRustProtobuf", pkg.TestRustProtobuf},

	{"TestRustTest", pkg.TestRustTest},

	{"TestRustTestLinkage", pkg.TestRustTestLinkage},

	{"TestSanitizeMemtagHeap", pkg.TestSanitizeMemtagHeap},

	{"TestSanitizeMemtagHeapWithSanitizeDevice", pkg.TestSanitizeMemtagHeapWithSanitizeDevice},

	{"TestSanitizeMemtagHeapWithSanitizeDeviceDiag", pkg.TestSanitizeMemtagHeapWithSanitizeDeviceDiag},

	{"TestSharedLibrary", pkg.TestSharedLibrary},

	{"TestSharedLibraryToc", pkg.TestSharedLibraryToc},

	{"TestSourceProviderCollision", pkg.TestSourceProviderCollision},

	{"TestSourceProviderDeps", pkg.TestSourceProviderDeps},

	{"TestSourceProviderRequiredFields", pkg.TestSourceProviderRequiredFields},

	{"TestSourceProviderTargetMismatch", pkg.TestSourceProviderTargetMismatch},

	{"TestStaticBinaryFlags", pkg.TestStaticBinaryFlags},

	{"TestStaticLibraryLinkage", pkg.TestStaticLibraryLinkage},

	{"TestStdDeviceLinkage", pkg.TestStdDeviceLinkage},

	{"TestStrippedBinary", pkg.TestStrippedBinary},

	{"TestStrippedLibrary", pkg.TestStrippedLibrary},

	{"TestValidateLibraryStem", pkg.TestValidateLibraryStem},

	{"TestVendorLinkage", pkg.TestVendorLinkage},

	{"TestVendorRamdiskLinkage", pkg.TestVendorRamdiskLinkage},

	{"TestVendorSnapshotCapture", pkg.TestVendorSnapshotCapture},

	{"TestVendorSnapshotDirected", pkg.TestVendorSnapshotDirected},

	{"TestVendorSnapshotExclude", pkg.TestVendorSnapshotExclude},

	{"TestVendorSnapshotUse", pkg.TestVendorSnapshotUse},

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
	return "android/soong/rust"
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
