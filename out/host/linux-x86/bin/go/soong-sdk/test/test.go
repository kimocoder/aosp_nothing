
package main

import (
	"io"

	"reflect"
	"regexp"
	"testing"
	"time"

	pkg "android/soong/sdk"
)

var t = []testing.InternalTest{

	{"TestAddPropertySetExisting", pkg.TestAddPropertySetExisting},

	{"TestAddPropertySetNew", pkg.TestAddPropertySetNew},

	{"TestAddPropertySimple", pkg.TestAddPropertySimple},

	{"TestAddPropertySubset", pkg.TestAddPropertySubset},

	{"TestBasicSdkWithBootclasspathFragment", pkg.TestBasicSdkWithBootclasspathFragment},

	{"TestBasicTrait_MultipleTraits", pkg.TestBasicTrait_MultipleTraits},

	{"TestBasicTrait_WithoutTrait", pkg.TestBasicTrait_WithoutTrait},

	{"TestBuildReleaseSetContains", pkg.TestBuildReleaseSetContains},

	{"TestCommonValueOptimization", pkg.TestCommonValueOptimization},

	{"TestCommonValueOptimization_InvalidArchSpecificVariants", pkg.TestCommonValueOptimization_InvalidArchSpecificVariants},

	{"TestDeviceAndHostSnapshotWithCcHeadersLibrary", pkg.TestDeviceAndHostSnapshotWithCcHeadersLibrary},

	{"TestDeviceAndHostSnapshotWithJavaHeaderLibrary", pkg.TestDeviceAndHostSnapshotWithJavaHeaderLibrary},

	{"TestDeviceAndHostSnapshotWithOsSpecificMembers", pkg.TestDeviceAndHostSnapshotWithOsSpecificMembers},

	{"TestDeviceAndHostSnapshotWithStubsLibrary", pkg.TestDeviceAndHostSnapshotWithStubsLibrary},

	{"TestHostSdkWithCc", pkg.TestHostSdkWithCc},

	{"TestHostSnapshotWithCcHeadersLibrary", pkg.TestHostSnapshotWithCcHeadersLibrary},

	{"TestHostSnapshotWithCcSharedLibrary", pkg.TestHostSnapshotWithCcSharedLibrary},

	{"TestHostSnapshotWithCcStaticLibrary", pkg.TestHostSnapshotWithCcStaticLibrary},

	{"TestHostSnapshotWithJavaHeaderLibrary", pkg.TestHostSnapshotWithJavaHeaderLibrary},

	{"TestHostSnapshotWithJavaImplLibrary", pkg.TestHostSnapshotWithJavaImplLibrary},

	{"TestHostSnapshotWithJavaSystemModules", pkg.TestHostSnapshotWithJavaSystemModules},

	{"TestHostSnapshotWithJavaTest", pkg.TestHostSnapshotWithJavaTest},

	{"TestHostSnapshotWithMultiLib64", pkg.TestHostSnapshotWithMultiLib64},

	{"TestModuleExportsSnapshot", pkg.TestModuleExportsSnapshot},

	{"TestMultipleHostOsTypesSnapshotWithCcBinary", pkg.TestMultipleHostOsTypesSnapshotWithCcBinary},

	{"TestMultipleHostOsTypesSnapshotWithCcSharedLibrary", pkg.TestMultipleHostOsTypesSnapshotWithCcSharedLibrary},

	{"TestNameToRelease", pkg.TestNameToRelease},

	{"TestNoSanitizerMembers", pkg.TestNoSanitizerMembers},

	{"TestParseBuildReleaseSet", pkg.TestParseBuildReleaseSet},

	{"TestPrebuiltVisibilityProperty_AddPrivate", pkg.TestPrebuiltVisibilityProperty_AddPrivate},

	{"TestPrebuiltVisibilityProperty_IsValidated", pkg.TestPrebuiltVisibilityProperty_IsValidated},

	{"TestPropertyPrunerByBuildRelease", pkg.TestPropertyPrunerByBuildRelease},

	{"TestPropertyPrunerInvalidTag", pkg.TestPropertyPrunerInvalidTag},

	{"TestSdkCompileMultilibOverride", pkg.TestSdkCompileMultilibOverride},

	{"TestSdkDependsOnSourceEvenWhenPrebuiltPreferred", pkg.TestSdkDependsOnSourceEvenWhenPrebuiltPreferred},

	{"TestSdkInstall", pkg.TestSdkInstall},

	{"TestSdkIsCompileMultilibBoth", pkg.TestSdkIsCompileMultilibBoth},

	{"TestSdkWithCc", pkg.TestSdkWithCc},

	{"TestSingleDeviceOsAssumption", pkg.TestSingleDeviceOsAssumption},

	{"TestSnapshotVisibility", pkg.TestSnapshotVisibility},

	{"TestSnapshotWithBootClasspathFragment_Contents", pkg.TestSnapshotWithBootClasspathFragment_Contents},

	{"TestSnapshotWithBootClasspathFragment_Fragments", pkg.TestSnapshotWithBootClasspathFragment_Fragments},

	{"TestSnapshotWithBootclasspathFragment_HiddenAPI", pkg.TestSnapshotWithBootclasspathFragment_HiddenAPI},

	{"TestSnapshotWithBootclasspathFragment_ImageName", pkg.TestSnapshotWithBootclasspathFragment_ImageName},

	{"TestSnapshotWithCcBinary", pkg.TestSnapshotWithCcBinary},

	{"TestSnapshotWithCcDuplicateHeaders", pkg.TestSnapshotWithCcDuplicateHeaders},

	{"TestSnapshotWithCcExportGeneratedHeaders", pkg.TestSnapshotWithCcExportGeneratedHeaders},

	{"TestSnapshotWithCcHeadersLibrary", pkg.TestSnapshotWithCcHeadersLibrary},

	{"TestSnapshotWithCcHeadersLibraryAndImageVariants", pkg.TestSnapshotWithCcHeadersLibraryAndImageVariants},

	{"TestSnapshotWithCcHeadersLibraryAndNativeBridgeSupport", pkg.TestSnapshotWithCcHeadersLibraryAndNativeBridgeSupport},

	{"TestSnapshotWithCcHeadersLibrary_DetectsNativeBridgeSpecificProperties", pkg.TestSnapshotWithCcHeadersLibrary_DetectsNativeBridgeSpecificProperties},

	{"TestSnapshotWithCcLibrary", pkg.TestSnapshotWithCcLibrary},

	{"TestSnapshotWithCcSharedLibrary", pkg.TestSnapshotWithCcSharedLibrary},

	{"TestSnapshotWithCcSharedLibraryCommonProperties", pkg.TestSnapshotWithCcSharedLibraryCommonProperties},

	{"TestSnapshotWithCcSharedLibrarySharedLibs", pkg.TestSnapshotWithCcSharedLibrarySharedLibs},

	{"TestSnapshotWithCcStaticLibrary", pkg.TestSnapshotWithCcStaticLibrary},

	{"TestSnapshotWithCcStaticNocrtBinary", pkg.TestSnapshotWithCcStaticNocrtBinary},

	{"TestSnapshotWithCompatConfig", pkg.TestSnapshotWithCompatConfig},

	{"TestSnapshotWithJavaBootLibrary", pkg.TestSnapshotWithJavaBootLibrary},

	{"TestSnapshotWithJavaBootLibrary_UpdatableMedia", pkg.TestSnapshotWithJavaBootLibrary_UpdatableMedia},

	{"TestSnapshotWithJavaHeaderLibrary", pkg.TestSnapshotWithJavaHeaderLibrary},

	{"TestSnapshotWithJavaImplLibrary", pkg.TestSnapshotWithJavaImplLibrary},

	{"TestSnapshotWithJavaSdkLibrary", pkg.TestSnapshotWithJavaSdkLibrary},

	{"TestSnapshotWithJavaSdkLibrary_AnnotationsZip", pkg.TestSnapshotWithJavaSdkLibrary_AnnotationsZip},

	{"TestSnapshotWithJavaSdkLibrary_AnnotationsZip_PreT", pkg.TestSnapshotWithJavaSdkLibrary_AnnotationsZip_PreT},

	{"TestSnapshotWithJavaSdkLibrary_ApiScopes", pkg.TestSnapshotWithJavaSdkLibrary_ApiScopes},

	{"TestSnapshotWithJavaSdkLibrary_CompileDex", pkg.TestSnapshotWithJavaSdkLibrary_CompileDex},

	{"TestSnapshotWithJavaSdkLibrary_DoctagFiles", pkg.TestSnapshotWithJavaSdkLibrary_DoctagFiles},

	{"TestSnapshotWithJavaSdkLibrary_ModuleLib", pkg.TestSnapshotWithJavaSdkLibrary_ModuleLib},

	{"TestSnapshotWithJavaSdkLibrary_NamingScheme", pkg.TestSnapshotWithJavaSdkLibrary_NamingScheme},

	{"TestSnapshotWithJavaSdkLibrary_SdkVersion_ForScope", pkg.TestSnapshotWithJavaSdkLibrary_SdkVersion_ForScope},

	{"TestSnapshotWithJavaSdkLibrary_SdkVersion_None", pkg.TestSnapshotWithJavaSdkLibrary_SdkVersion_None},

	{"TestSnapshotWithJavaSdkLibrary_SystemServer", pkg.TestSnapshotWithJavaSdkLibrary_SystemServer},

	{"TestSnapshotWithJavaSdkLibrary_UseSrcJar", pkg.TestSnapshotWithJavaSdkLibrary_UseSrcJar},

	{"TestSnapshotWithJavaSystemModules", pkg.TestSnapshotWithJavaSystemModules},

	{"TestSnapshotWithJavaSystemserverLibrary", pkg.TestSnapshotWithJavaSystemserverLibrary},

	{"TestSnapshotWithJavaTest", pkg.TestSnapshotWithJavaTest},

	{"TestSnapshotWithObject", pkg.TestSnapshotWithObject},

	{"TestSnapshotWithPackageDefaultLicense", pkg.TestSnapshotWithPackageDefaultLicense},

	{"TestSnapshotWithSingleHostOsType", pkg.TestSnapshotWithSingleHostOsType},

	{"TestSnapshotWithSystemServerClasspathFragment", pkg.TestSnapshotWithSystemServerClasspathFragment},

	{"TestSnapshot_EnvConfiguration", pkg.TestSnapshot_EnvConfiguration},

	{"TestStubsLibrary", pkg.TestStubsLibrary},

	{"TestSystemSharedLibPropagation", pkg.TestSystemSharedLibPropagation},

	{"TestTraitUnsupportedByMemberType", pkg.TestTraitUnsupportedByMemberType},

	{"TestTransformRemoveProperty", pkg.TestTransformRemoveProperty},

	{"TestTransformRemovePropertySet", pkg.TestTransformRemovePropertySet},

	{"TestUniqueHostSoname", pkg.TestUniqueHostSoname},

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
	return "android/soong/sdk"
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
