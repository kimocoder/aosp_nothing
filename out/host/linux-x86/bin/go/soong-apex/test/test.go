
package main

import (
	"io"

	"reflect"
	"regexp"
	"testing"
	"time"

	pkg "android/soong/apex"
)

var t = []testing.InternalTest{

	{"TestAllowedFiles", pkg.TestAllowedFiles},

	{"TestAndroidMkWritesCommonProperties", pkg.TestAndroidMkWritesCommonProperties},

	{"TestAndroidMk_DexpreoptBuiltInstalledForApex", pkg.TestAndroidMk_DexpreoptBuiltInstalledForApex},

	{"TestAndroidMk_DexpreoptBuiltInstalledForApex_Prebuilt", pkg.TestAndroidMk_DexpreoptBuiltInstalledForApex_Prebuilt},

	{"TestAndroidMk_RequiredDeps", pkg.TestAndroidMk_RequiredDeps},

	{"TestAndroidMk_RequiredModules", pkg.TestAndroidMk_RequiredModules},

	{"TestAndroidMk_VendorApexRequired", pkg.TestAndroidMk_VendorApexRequired},

	{"TestApexAvailable_CheckForPlatform", pkg.TestApexAvailable_CheckForPlatform},

	{"TestApexAvailable_CreatedForApex", pkg.TestApexAvailable_CreatedForApex},

	{"TestApexAvailable_DirectDep", pkg.TestApexAvailable_DirectDep},

	{"TestApexAvailable_IndirectDep", pkg.TestApexAvailable_IndirectDep},

	{"TestApexAvailable_InvalidApexName", pkg.TestApexAvailable_InvalidApexName},

	{"TestApexCanUsePrivateApis", pkg.TestApexCanUsePrivateApis},

	{"TestApexDependsOnLLNDKTransitively", pkg.TestApexDependsOnLLNDKTransitively},

	{"TestApexInVariousPartition", pkg.TestApexInVariousPartition},

	{"TestApexJavaCoverage", pkg.TestApexJavaCoverage},

	{"TestApexKeyFromOtherModule", pkg.TestApexKeyFromOtherModule},

	{"TestApexKeysTxt", pkg.TestApexKeysTxt},

	{"TestApexKeysTxtOverrides", pkg.TestApexKeysTxtOverrides},

	{"TestApexManifest", pkg.TestApexManifest},

	{"TestApexManifestMinSdkVersion", pkg.TestApexManifestMinSdkVersion},

	{"TestApexMinSdkVersion_DefaultsToLatest", pkg.TestApexMinSdkVersion_DefaultsToLatest},

	{"TestApexMinSdkVersion_ErrorIfDepIsNewer", pkg.TestApexMinSdkVersion_ErrorIfDepIsNewer},

	{"TestApexMinSdkVersion_ErrorIfDepIsNewer_Java", pkg.TestApexMinSdkVersion_ErrorIfDepIsNewer_Java},

	{"TestApexMinSdkVersion_ErrorIfIncompatibleVersion", pkg.TestApexMinSdkVersion_ErrorIfIncompatibleVersion},

	{"TestApexMinSdkVersion_NativeModulesShouldBeBuiltAgainstStubs", pkg.TestApexMinSdkVersion_NativeModulesShouldBeBuiltAgainstStubs},

	{"TestApexMinSdkVersion_Okay", pkg.TestApexMinSdkVersion_Okay},

	{"TestApexMinSdkVersion_OkayEvenWhenDepIsNewer_IfItSatisfiesApexMinSdkVersion", pkg.TestApexMinSdkVersion_OkayEvenWhenDepIsNewer_IfItSatisfiesApexMinSdkVersion},

	{"TestApexMinSdkVersion_SupportsCodeNames", pkg.TestApexMinSdkVersion_SupportsCodeNames},

	{"TestApexMinSdkVersion_SupportsCodeNames_JavaLibs", pkg.TestApexMinSdkVersion_SupportsCodeNames_JavaLibs},

	{"TestApexMinSdkVersion_WorksWithActiveCodenames", pkg.TestApexMinSdkVersion_WorksWithActiveCodenames},

	{"TestApexMinSdkVersion_WorksWithSdkCodename", pkg.TestApexMinSdkVersion_WorksWithSdkCodename},

	{"TestApexMinSdkVersion_crtobjectInVendorApex", pkg.TestApexMinSdkVersion_crtobjectInVendorApex},

	{"TestApexMutatorsDontRunIfDisabled", pkg.TestApexMutatorsDontRunIfDisabled},

	{"TestApexName", pkg.TestApexName},

	{"TestApexOutputFileProducer", pkg.TestApexOutputFileProducer},

	{"TestApexPermittedPackagesRules", pkg.TestApexPermittedPackagesRules},

	{"TestApexPropertiesShouldBeDefaultable", pkg.TestApexPropertiesShouldBeDefaultable},

	{"TestApexSet", pkg.TestApexSet},

	{"TestApexSetApksModuleAssignment", pkg.TestApexSetApksModuleAssignment},

	{"TestApexSetFilenameOverride", pkg.TestApexSetFilenameOverride},

	{"TestApexSet_NativeBridge", pkg.TestApexSet_NativeBridge},

	{"TestApexStrictUpdtabilityLint", pkg.TestApexStrictUpdtabilityLint},

	{"TestApexStrictUpdtabilityLintBcpFragmentDeps", pkg.TestApexStrictUpdtabilityLintBcpFragmentDeps},

	{"TestApexWithAppImportBuildId", pkg.TestApexWithAppImportBuildId},

	{"TestApexWithAppImports", pkg.TestApexWithAppImports},

	{"TestApexWithAppImportsPrefer", pkg.TestApexWithAppImportsPrefer},

	{"TestApexWithApps", pkg.TestApexWithApps},

	{"TestApexWithArch", pkg.TestApexWithArch},

	{"TestApexWithExplicitStubsDependency", pkg.TestApexWithExplicitStubsDependency},

	{"TestApexWithJavaImport", pkg.TestApexWithJavaImport},

	{"TestApexWithJniLibs", pkg.TestApexWithJniLibs},

	{"TestApexWithRuntimeLibsDependency", pkg.TestApexWithRuntimeLibsDependency},

	{"TestApexWithShBinary", pkg.TestApexWithShBinary},

	{"TestApexWithStubs", pkg.TestApexWithStubs},

	{"TestApexWithStubsWithMinSdkVersion", pkg.TestApexWithStubsWithMinSdkVersion},

	{"TestApexWithSystemLibsStubs", pkg.TestApexWithSystemLibsStubs},

	{"TestApexWithTarget", pkg.TestApexWithTarget},

	{"TestApexWithTestHelperApp", pkg.TestApexWithTestHelperApp},

	{"TestApexWithTests", pkg.TestApexWithTests},

	{"TestApex_PlatformUsesLatestStubFromApex", pkg.TestApex_PlatformUsesLatestStubFromApex},

	{"TestApex_withPrebuiltFirmware", pkg.TestApex_withPrebuiltFirmware},

	{"TestAppBundle", pkg.TestAppBundle},

	{"TestAppSetBundle", pkg.TestAppSetBundle},

	{"TestAppSetBundlePrebuilt", pkg.TestAppSetBundlePrebuilt},

	{"TestBasicApex", pkg.TestBasicApex},

	{"TestBasicZipApex", pkg.TestBasicZipApex},

	{"TestBootDexJarsFromSourcesAndPrebuilts", pkg.TestBootDexJarsFromSourcesAndPrebuilts},

	{"TestBootFragmentNotInApex", pkg.TestBootFragmentNotInApex},

	{"TestBootJarNotInApex", pkg.TestBootJarNotInApex},

	{"TestBootclasspathFragmentContentsNoName", pkg.TestBootclasspathFragmentContentsNoName},

	{"TestBootclasspathFragmentInArtApex", pkg.TestBootclasspathFragmentInArtApex},

	{"TestBootclasspathFragmentInPrebuiltArtApex", pkg.TestBootclasspathFragmentInPrebuiltArtApex},

	{"TestBootclasspathFragment_AndroidNonUpdatable", pkg.TestBootclasspathFragment_AndroidNonUpdatable},

	{"TestBootclasspathFragment_AndroidNonUpdatable_AlwaysUsePrebuiltSdks", pkg.TestBootclasspathFragment_AndroidNonUpdatable_AlwaysUsePrebuiltSdks},

	{"TestBootclasspathFragment_HiddenAPIList", pkg.TestBootclasspathFragment_HiddenAPIList},

	{"TestBootclasspathFragments", pkg.TestBootclasspathFragments},

	{"TestBootclasspathFragments_FragmentDependency", pkg.TestBootclasspathFragments_FragmentDependency},

	{"TestCarryRequiredModuleNames", pkg.TestCarryRequiredModuleNames},

	{"TestCertificate", pkg.TestCertificate},

	{"TestCompatConfig", pkg.TestCompatConfig},

	{"TestCompressedApex", pkg.TestCompressedApex},

	{"TestCreateClasspathElements", pkg.TestCreateClasspathElements},

	{"TestDefaults", pkg.TestDefaults},

	{"TestDependenciesInApexManifest", pkg.TestDependenciesInApexManifest},

	{"TestDexpreoptAccessDexFilesFromPrebuiltApex", pkg.TestDexpreoptAccessDexFilesFromPrebuiltApex},

	{"TestDuplicateButEquivalentDeapexersFromPrebuiltApexes", pkg.TestDuplicateButEquivalentDeapexersFromPrebuiltApexes},

	{"TestDuplicateDeapexersFromPrebuiltApexes", pkg.TestDuplicateDeapexersFromPrebuiltApexes},

	{"TestErrorsIfDepsAreNotEnabled", pkg.TestErrorsIfDepsAreNotEnabled},

	{"TestExcludeDependency", pkg.TestExcludeDependency},

	{"TestFileContexts_FindInDefaultLocationIfNotSet", pkg.TestFileContexts_FindInDefaultLocationIfNotSet},

	{"TestFileContexts_ProductSpecificApexes", pkg.TestFileContexts_ProductSpecificApexes},

	{"TestFileContexts_SetViaFileGroup", pkg.TestFileContexts_SetViaFileGroup},

	{"TestFileContexts_ShouldBeUnderSystemSepolicyForSystemApexes", pkg.TestFileContexts_ShouldBeUnderSystemSepolicyForSystemApexes},

	{"TestFilesInSubDir", pkg.TestFilesInSubDir},

	{"TestFilesInSubDirWhenNativeBridgeEnabled", pkg.TestFilesInSubDirWhenNativeBridgeEnabled},

	{"TestHeaderLibsDependency", pkg.TestHeaderLibsDependency},

	{"TestHostApexInHostOnlyBuild", pkg.TestHostApexInHostOnlyBuild},

	{"TestIndirectTestFor", pkg.TestIndirectTestFor},

	{"TestInstallExtraFlattenedApexes", pkg.TestInstallExtraFlattenedApexes},

	{"TestJavaSDKLibrary", pkg.TestJavaSDKLibrary},

	{"TestJavaSDKLibrary_CrossBoundary", pkg.TestJavaSDKLibrary_CrossBoundary},

	{"TestJavaSDKLibrary_ImportOnly", pkg.TestJavaSDKLibrary_ImportOnly},

	{"TestJavaSDKLibrary_ImportPreferred", pkg.TestJavaSDKLibrary_ImportPreferred},

	{"TestJavaSDKLibrary_WithinApex", pkg.TestJavaSDKLibrary_WithinApex},

	{"TestJavaStableSdkVersion", pkg.TestJavaStableSdkVersion},

	{"TestKeys", pkg.TestKeys},

	{"TestLegacyAndroid10Support", pkg.TestLegacyAndroid10Support},

	{"TestMacro", pkg.TestMacro},

	{"TestMinSdkVersionOverride", pkg.TestMinSdkVersionOverride},

	{"TestMinSdkVersionOverrideToLowerVersionNoOp", pkg.TestMinSdkVersionOverrideToLowerVersionNoOp},

	{"TestNoStaticLinkingToStubsLib", pkg.TestNoStaticLinkingToStubsLib},

	{"TestNoUpdatableJarsInBootImage", pkg.TestNoUpdatableJarsInBootImage},

	{"TestNonBootJarInFragment", pkg.TestNonBootJarInFragment},

	{"TestNonPreferredPrebuiltDependency", pkg.TestNonPreferredPrebuiltDependency},

	{"TestNonTestApex", pkg.TestNonTestApex},

	{"TestOverrideApex", pkg.TestOverrideApex},

	{"TestPlatformBootclasspathDependencies", pkg.TestPlatformBootclasspathDependencies},

	{"TestPlatformBootclasspath_AlwaysUsePrebuiltSdks", pkg.TestPlatformBootclasspath_AlwaysUsePrebuiltSdks},

	{"TestPlatformBootclasspath_Fragments", pkg.TestPlatformBootclasspath_Fragments},

	{"TestPlatformBootclasspath_IncludesRemainingApexJars", pkg.TestPlatformBootclasspath_IncludesRemainingApexJars},

	{"TestPlatformBootclasspath_LegacyPrebuiltFragment", pkg.TestPlatformBootclasspath_LegacyPrebuiltFragment},

	{"TestPlatformUsesLatestStubsFromApexes", pkg.TestPlatformUsesLatestStubsFromApexes},

	{"TestPrebuilt", pkg.TestPrebuilt},

	{"TestPrebuiltApexName", pkg.TestPrebuiltApexName},

	{"TestPrebuiltApexNameWithPlatformBootclasspath", pkg.TestPrebuiltApexNameWithPlatformBootclasspath},

	{"TestPrebuiltExportDexImplementationJars", pkg.TestPrebuiltExportDexImplementationJars},

	{"TestPrebuiltFilenameOverride", pkg.TestPrebuiltFilenameOverride},

	{"TestPrebuiltMissingSrc", pkg.TestPrebuiltMissingSrc},

	{"TestPrebuiltOverrides", pkg.TestPrebuiltOverrides},

	{"TestPrebuiltStandaloneSystemserverclasspathFragmentContents", pkg.TestPrebuiltStandaloneSystemserverclasspathFragmentContents},

	{"TestPrebuiltStubLibDep", pkg.TestPrebuiltStubLibDep},

	{"TestPrebuiltSystemserverclasspathFragmentContents", pkg.TestPrebuiltSystemserverclasspathFragmentContents},

	{"TestPreferredPrebuiltSharedLibDep", pkg.TestPreferredPrebuiltSharedLibDep},

	{"TestProductVariant", pkg.TestProductVariant},

	{"TestProhibitStaticExecutable", pkg.TestProhibitStaticExecutable},

	{"TestQApexesUseLatestStubsInBundledBuildsAndHWASAN", pkg.TestQApexesUseLatestStubsInBundledBuildsAndHWASAN},

	{"TestQTargetApexUsesStaticUnwinder", pkg.TestQTargetApexUsesStaticUnwinder},

	{"TestRejectNonInstallableJavaLibrary", pkg.TestRejectNonInstallableJavaLibrary},

	{"TestRuntimeApexShouldInstallHwasanIfHwaddressSanitized", pkg.TestRuntimeApexShouldInstallHwasanIfHwaddressSanitized},

	{"TestRuntimeApexShouldInstallHwasanIfLibcDependsOnIt", pkg.TestRuntimeApexShouldInstallHwasanIfLibcDependsOnIt},

	{"TestSdkLibraryCanHaveHigherMinSdkVersion", pkg.TestSdkLibraryCanHaveHigherMinSdkVersion},

	{"TestStaticLinking", pkg.TestStaticLinking},

	{"TestSymlinksFromApexToSystem", pkg.TestSymlinksFromApexToSystem},

	{"TestSymlinksFromApexToSystemRequiredModuleNames", pkg.TestSymlinksFromApexToSystemRequiredModuleNames},

	{"TestSystemServerClasspathFragmentWithContentNotInMake", pkg.TestSystemServerClasspathFragmentWithContentNotInMake},

	{"TestSystemserverclasspathFragmentContents", pkg.TestSystemserverclasspathFragmentContents},

	{"TestSystemserverclasspathFragmentNoGeneratedProto", pkg.TestSystemserverclasspathFragmentNoGeneratedProto},

	{"TestSystemserverclasspathFragmentStandaloneContents", pkg.TestSystemserverclasspathFragmentStandaloneContents},

	{"TestTestApex", pkg.TestTestApex},

	{"TestTestFor", pkg.TestTestFor},

	{"TestTestForForLibInOtherApex", pkg.TestTestForForLibInOtherApex},

	{"TestUpdatabilityLintSkipLibcore", pkg.TestUpdatabilityLintSkipLibcore},

	{"TestUpdatableApexEnforcesAppUpdatability", pkg.TestUpdatableApexEnforcesAppUpdatability},

	{"TestUpdatableDefault_should_set_min_sdk_version", pkg.TestUpdatableDefault_should_set_min_sdk_version},

	{"TestUpdatable_cannot_be_vendor_apex", pkg.TestUpdatable_cannot_be_vendor_apex},

	{"TestUpdatable_should_not_set_generate_classpaths_proto", pkg.TestUpdatable_should_not_set_generate_classpaths_proto},

	{"TestUpdatable_should_set_min_sdk_version", pkg.TestUpdatable_should_set_min_sdk_version},

	{"TestVendorApex", pkg.TestVendorApex},

	{"TestVendorApex_use_vndk_as_stable", pkg.TestVendorApex_use_vndk_as_stable},

	{"TestVendorApex_use_vndk_as_stable_TryingToIncludeVNDKLib", pkg.TestVendorApex_use_vndk_as_stable_TryingToIncludeVNDKLib},

	{"TestVndkApexCurrent", pkg.TestVndkApexCurrent},

	{"TestVndkApexDoesntSupportNativeBridgeSupported", pkg.TestVndkApexDoesntSupportNativeBridgeSupported},

	{"TestVndkApexForVndkLite", pkg.TestVndkApexForVndkLite},

	{"TestVndkApexNameRule", pkg.TestVndkApexNameRule},

	{"TestVndkApexShouldNotProvideNativeLibs", pkg.TestVndkApexShouldNotProvideNativeLibs},

	{"TestVndkApexSkipsNativeBridgeSupportedModules", pkg.TestVndkApexSkipsNativeBridgeSupportedModules},

	{"TestVndkApexUsesVendorVariant", pkg.TestVndkApexUsesVendorVariant},

	{"TestVndkApexVersion", pkg.TestVndkApexVersion},

	{"TestVndkApexWithBinder32", pkg.TestVndkApexWithBinder32},

	{"TestVndkApexWithPrebuilt", pkg.TestVndkApexWithPrebuilt},

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
	return "android/soong/apex"
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
