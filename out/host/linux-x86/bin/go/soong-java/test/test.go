
package main

import (
	"io"

	"reflect"
	"regexp"
	"testing"
	"time"

	pkg "android/soong/java"
)

var t = []testing.InternalTest{

	{"TestAidlEnforcePermissions", pkg.TestAidlEnforcePermissions},

	{"TestAidlEnforcePermissionsException", pkg.TestAidlEnforcePermissionsException},

	{"TestAidlExportIncludeDirsFromImports", pkg.TestAidlExportIncludeDirsFromImports},

	{"TestAidlFlagsArePassedToTheAidlCompiler", pkg.TestAidlFlagsArePassedToTheAidlCompiler},

	{"TestAidlFlagsWithMinSdkVersion", pkg.TestAidlFlagsWithMinSdkVersion},

	{"TestAndroidAppImport", pkg.TestAndroidAppImport},

	{"TestAndroidAppImport_ArchVariants", pkg.TestAndroidAppImport_ArchVariants},

	{"TestAndroidAppImport_DefaultDevCert", pkg.TestAndroidAppImport_DefaultDevCert},

	{"TestAndroidAppImport_DpiVariants", pkg.TestAndroidAppImport_DpiVariants},

	{"TestAndroidAppImport_Filename", pkg.TestAndroidAppImport_Filename},

	{"TestAndroidAppImport_NoDexPreopt", pkg.TestAndroidAppImport_NoDexPreopt},

	{"TestAndroidAppImport_Presigned", pkg.TestAndroidAppImport_Presigned},

	{"TestAndroidAppImport_SigningLineage", pkg.TestAndroidAppImport_SigningLineage},

	{"TestAndroidAppImport_SigningLineageFilegroup", pkg.TestAndroidAppImport_SigningLineageFilegroup},

	{"TestAndroidAppImport_frameworkRes", pkg.TestAndroidAppImport_frameworkRes},

	{"TestAndroidAppImport_overridesDisabledAndroidApp", pkg.TestAndroidAppImport_overridesDisabledAndroidApp},

	{"TestAndroidAppImport_relativeInstallPath", pkg.TestAndroidAppImport_relativeInstallPath},

	{"TestAndroidAppLinkType", pkg.TestAndroidAppLinkType},

	{"TestAndroidAppSet", pkg.TestAndroidAppSet},

	{"TestAndroidAppSet_Variants", pkg.TestAndroidAppSet_Variants},

	{"TestAndroidMkEntriesForApex", pkg.TestAndroidMkEntriesForApex},

	{"TestAndroidResources", pkg.TestAndroidResources},

	{"TestAndroidTestHelperApp_LocalDisableTestConfig", pkg.TestAndroidTestHelperApp_LocalDisableTestConfig},

	{"TestAndroidTestImport", pkg.TestAndroidTestImport},

	{"TestAndroidTestImport_NoJinUncompressForPresigned", pkg.TestAndroidTestImport_NoJinUncompressForPresigned},

	{"TestAndroidTestImport_Preprocessed", pkg.TestAndroidTestImport_Preprocessed},

	{"TestAndroidTestImport_UncompressDex", pkg.TestAndroidTestImport_UncompressDex},

	{"TestAndroidTest_FixTestConfig", pkg.TestAndroidTest_FixTestConfig},

	{"TestApp", pkg.TestApp},

	{"TestAppJavaResources", pkg.TestAppJavaResources},

	{"TestAppMissingCertificateAllowMissingDependencies", pkg.TestAppMissingCertificateAllowMissingDependencies},

	{"TestAppSdkVersion", pkg.TestAppSdkVersion},

	{"TestAppSdkVersionByPartition", pkg.TestAppSdkVersionByPartition},

	{"TestAppSplits", pkg.TestAppSplits},

	{"TestArchSpecific", pkg.TestArchSpecific},

	{"TestBinary", pkg.TestBinary},

	{"TestBootclasspathFragmentInconsistentArtConfiguration_ApexMixture", pkg.TestBootclasspathFragmentInconsistentArtConfiguration_ApexMixture},

	{"TestBootclasspathFragmentInconsistentArtConfiguration_Platform", pkg.TestBootclasspathFragmentInconsistentArtConfiguration_Platform},

	{"TestBootclasspathFragment_Coverage", pkg.TestBootclasspathFragment_Coverage},

	{"TestBootclasspathFragment_StubLibs", pkg.TestBootclasspathFragment_StubLibs},

	{"TestBootclasspathFragment_UnknownImageName", pkg.TestBootclasspathFragment_UnknownImageName},

	{"TestCertificates", pkg.TestCertificates},

	{"TestClasspath", pkg.TestClasspath},

	{"TestCodelessApp", pkg.TestCodelessApp},

	{"TestCollectJavaLibraryPropertiesAddAidlIncludeDirs", pkg.TestCollectJavaLibraryPropertiesAddAidlIncludeDirs},

	{"TestCollectJavaLibraryPropertiesAddJarjarRules", pkg.TestCollectJavaLibraryPropertiesAddJarjarRules},

	{"TestCollectJavaLibraryPropertiesAddLibsDeps", pkg.TestCollectJavaLibraryPropertiesAddLibsDeps},

	{"TestCollectJavaLibraryPropertiesAddScrs", pkg.TestCollectJavaLibraryPropertiesAddScrs},

	{"TestCollectJavaLibraryPropertiesAddStaticLibsDeps", pkg.TestCollectJavaLibraryPropertiesAddStaticLibsDeps},

	{"TestCompilerFlags", pkg.TestCompilerFlags},

	{"TestD8", pkg.TestD8},

	{"TestDataDeviceBinsBuildsDeviceBinary", pkg.TestDataDeviceBinsBuildsDeviceBinary},

	{"TestDataNativeBinaries", pkg.TestDataNativeBinaries},

	{"TestDefaultInstallable", pkg.TestDefaultInstallable},

	{"TestDefaults", pkg.TestDefaults},

	{"TestDeviceForHost", pkg.TestDeviceForHost},

	{"TestDex2oatToolDeps", pkg.TestDex2oatToolDeps},

	{"TestDexpreoptBcp", pkg.TestDexpreoptBcp},

	{"TestDexpreoptBootJars", pkg.TestDexpreoptBootJars},

	{"TestDexpreoptBootZip", pkg.TestDexpreoptBootZip},

	{"TestDexpreoptBuiltInstalledForApex", pkg.TestDexpreoptBuiltInstalledForApex},

	{"TestDexpreoptEnabled", pkg.TestDexpreoptEnabled},

	{"TestDroiddoc", pkg.TestDroiddoc},

	{"TestDroiddocArgsAndFlagsCausesError", pkg.TestDroiddocArgsAndFlagsCausesError},

	{"TestDroidstubs", pkg.TestDroidstubs},

	{"TestDroidstubsSandbox", pkg.TestDroidstubsSandbox},

	{"TestDroidstubsWithSystemModules", pkg.TestDroidstubsWithSystemModules},

	{"TestEnforceRRO_propagatesToDependencies", pkg.TestEnforceRRO_propagatesToDependencies},

	{"TestErrorproneDisabled", pkg.TestErrorproneDisabled},

	{"TestErrorproneEnabled", pkg.TestErrorproneEnabled},

	{"TestErrorproneEnabledOnlyByEnvironmentVariable", pkg.TestErrorproneEnabledOnlyByEnvironmentVariable},

	{"TestExcludeFileGroupInSrcs", pkg.TestExcludeFileGroupInSrcs},

	{"TestExportedPlugins", pkg.TestExportedPlugins},

	{"TestExportedProguardFlagFiles", pkg.TestExportedProguardFlagFiles},

	{"TestGeneratedSources", pkg.TestGeneratedSources},

	{"TestGetOverriddenPackages", pkg.TestGetOverriddenPackages},

	{"TestHiddenAPIEncoding_JavaSdkLibrary", pkg.TestHiddenAPIEncoding_JavaSdkLibrary},

	{"TestHiddenAPISingleton", pkg.TestHiddenAPISingleton},

	{"TestHiddenAPISingletonSdks", pkg.TestHiddenAPISingletonSdks},

	{"TestHiddenAPISingletonWithPrebuilt", pkg.TestHiddenAPISingletonWithPrebuilt},

	{"TestHiddenAPISingletonWithPrebuiltCsvFile", pkg.TestHiddenAPISingletonWithPrebuiltCsvFile},

	{"TestHiddenAPISingletonWithPrebuiltOverrideSource", pkg.TestHiddenAPISingletonWithPrebuiltOverrideSource},

	{"TestHiddenAPISingletonWithPrebuiltUseSource", pkg.TestHiddenAPISingletonWithPrebuiltUseSource},

	{"TestHiddenAPISingletonWithSourceAndPrebuiltPreferredButNoDex", pkg.TestHiddenAPISingletonWithSourceAndPrebuiltPreferredButNoDex},

	{"TestHostBinaryNoJavaDebugInfoOverride", pkg.TestHostBinaryNoJavaDebugInfoOverride},

	{"TestHostForDevice", pkg.TestHostForDevice},

	{"TestHostdex", pkg.TestHostdex},

	{"TestHostdexRequired", pkg.TestHostdexRequired},

	{"TestHostdexSpecificRequired", pkg.TestHostdexSpecificRequired},

	{"TestImportSoongDexJar", pkg.TestImportSoongDexJar},

	{"TestIncludeSrcs", pkg.TestIncludeSrcs},

	{"TestInstrumentationTargetOverridden", pkg.TestInstrumentationTargetOverridden},

	{"TestInstrumentationTargetPrebuilt", pkg.TestInstrumentationTargetPrebuilt},

	{"TestJNIABI", pkg.TestJNIABI},

	{"TestJNIPackaging", pkg.TestJNIPackaging},

	{"TestJNISDK", pkg.TestJNISDK},

	{"TestJacocoFilterToSpecs", pkg.TestJacocoFilterToSpecs},

	{"TestJacocoFiltersToZipCommand", pkg.TestJacocoFiltersToZipCommand},

	{"TestJavaImport", pkg.TestJavaImport},

	{"TestJavaLibrary", pkg.TestJavaLibrary},

	{"TestJavaLibraryWithSystemModules", pkg.TestJavaLibraryWithSystemModules},

	{"TestJavaLinkType", pkg.TestJavaLinkType},

	{"TestJavaLint", pkg.TestJavaLint},

	{"TestJavaLintBypassUpdatableChecks", pkg.TestJavaLintBypassUpdatableChecks},

	{"TestJavaLintDatabaseSelectionFull", pkg.TestJavaLintDatabaseSelectionFull},

	{"TestJavaLintDatabaseSelectionPublicFiltered", pkg.TestJavaLintDatabaseSelectionPublicFiltered},

	{"TestJavaLintRequiresCustomLintFileToExist", pkg.TestJavaLintRequiresCustomLintFileToExist},

	{"TestJavaLintStrictUpdatabilityLinting", pkg.TestJavaLintStrictUpdatabilityLinting},

	{"TestJavaLintUsesCorrectBpConfig", pkg.TestJavaLintUsesCorrectBpConfig},

	{"TestJavaLintWithoutBaseline", pkg.TestJavaLintWithoutBaseline},

	{"TestJavaSdkLibrary", pkg.TestJavaSdkLibrary},

	{"TestJavaSdkLibraryDist", pkg.TestJavaSdkLibraryDist},

	{"TestJavaSdkLibraryEnforce", pkg.TestJavaSdkLibraryEnforce},

	{"TestJavaSdkLibraryImport", pkg.TestJavaSdkLibraryImport},

	{"TestJavaSdkLibraryImport_AccessOutputFiles", pkg.TestJavaSdkLibraryImport_AccessOutputFiles},

	{"TestJavaSdkLibraryImport_AccessOutputFiles_Invalid", pkg.TestJavaSdkLibraryImport_AccessOutputFiles_Invalid},

	{"TestJavaSdkLibraryImport_Preferred", pkg.TestJavaSdkLibraryImport_Preferred},

	{"TestJavaSdkLibraryImport_WithSource", pkg.TestJavaSdkLibraryImport_WithSource},

	{"TestJavaSdkLibrary_AccessOutputFiles", pkg.TestJavaSdkLibrary_AccessOutputFiles},

	{"TestJavaSdkLibrary_AccessOutputFiles_MissingScope", pkg.TestJavaSdkLibrary_AccessOutputFiles_MissingScope},

	{"TestJavaSdkLibrary_AccessOutputFiles_NoAnnotations", pkg.TestJavaSdkLibrary_AccessOutputFiles_NoAnnotations},

	{"TestJavaSdkLibrary_DefaultToStubs", pkg.TestJavaSdkLibrary_DefaultToStubs},

	{"TestJavaSdkLibrary_Deps", pkg.TestJavaSdkLibrary_Deps},

	{"TestJavaSdkLibrary_DoNotAccessImplWhenItIsNotBuilt", pkg.TestJavaSdkLibrary_DoNotAccessImplWhenItIsNotBuilt},

	{"TestJavaSdkLibrary_FallbackScope", pkg.TestJavaSdkLibrary_FallbackScope},

	{"TestJavaSdkLibrary_InvalidScopes", pkg.TestJavaSdkLibrary_InvalidScopes},

	{"TestJavaSdkLibrary_MissingScope", pkg.TestJavaSdkLibrary_MissingScope},

	{"TestJavaSdkLibrary_ModuleLib", pkg.TestJavaSdkLibrary_ModuleLib},

	{"TestJavaSdkLibrary_RequireXmlPermissionFile", pkg.TestJavaSdkLibrary_RequireXmlPermissionFile},

	{"TestJavaSdkLibrary_SdkVersion_ForScope", pkg.TestJavaSdkLibrary_SdkVersion_ForScope},

	{"TestJavaSdkLibrary_StubOrImplOnlyLibs", pkg.TestJavaSdkLibrary_StubOrImplOnlyLibs},

	{"TestJavaSdkLibrary_SystemServer", pkg.TestJavaSdkLibrary_SystemServer},

	{"TestJavaSdkLibrary_UpdatableLibrary", pkg.TestJavaSdkLibrary_UpdatableLibrary},

	{"TestJavaSdkLibrary_UpdatableLibrary_Validation_AtLeastTAttributes", pkg.TestJavaSdkLibrary_UpdatableLibrary_Validation_AtLeastTAttributes},

	{"TestJavaSdkLibrary_UpdatableLibrary_Validation_MinAndMaxDeviceSdk", pkg.TestJavaSdkLibrary_UpdatableLibrary_Validation_MinAndMaxDeviceSdk},

	{"TestJavaSdkLibrary_UpdatableLibrary_Validation_MinAndMaxDeviceSdkAndModuleMinSdk", pkg.TestJavaSdkLibrary_UpdatableLibrary_Validation_MinAndMaxDeviceSdkAndModuleMinSdk},

	{"TestJavaSdkLibrary_UpdatableLibrary_Validation_ValidVersion", pkg.TestJavaSdkLibrary_UpdatableLibrary_Validation_ValidVersion},

	{"TestJavaSdkLibrary_UpdatableLibrary_usesNewTag", pkg.TestJavaSdkLibrary_UpdatableLibrary_usesNewTag},

	{"TestJavaSystemModules", pkg.TestJavaSystemModules},

	{"TestJavaSystemModulesImport", pkg.TestJavaSystemModulesImport},

	{"TestJavaSystemModulesMixSourceAndPrebuilt", pkg.TestJavaSystemModulesMixSourceAndPrebuilt},

	{"TestKapt", pkg.TestKapt},

	{"TestKaptEncodeFlags", pkg.TestKaptEncodeFlags},

	{"TestKotlin", pkg.TestKotlin},

	{"TestKotlinCompose", pkg.TestKotlinCompose},

	{"TestLibraryAssets", pkg.TestLibraryAssets},

	{"TestModuleLibDroidstubs", pkg.TestModuleLibDroidstubs},

	{"TestNoPlugin", pkg.TestNoPlugin},

	{"TestOverrideAndroidApp", pkg.TestOverrideAndroidApp},

	{"TestOverrideAndroidAppDependency", pkg.TestOverrideAndroidAppDependency},

	{"TestOverrideAndroidAppOverrides", pkg.TestOverrideAndroidAppOverrides},

	{"TestOverrideAndroidAppStem", pkg.TestOverrideAndroidAppStem},

	{"TestOverrideAndroidAppWithPrebuilt", pkg.TestOverrideAndroidAppWithPrebuilt},

	{"TestOverrideAndroidTest", pkg.TestOverrideAndroidTest},

	{"TestOverrideRuntimeResourceOverlay", pkg.TestOverrideRuntimeResourceOverlay},

	{"TestPackageNameOverride", pkg.TestPackageNameOverride},

	{"TestPatchModule", pkg.TestPatchModule},

	{"TestPlatformAPIs", pkg.TestPlatformAPIs},

	{"TestPlatformBootclasspath", pkg.TestPlatformBootclasspath},

	{"TestPlatformBootclasspathModule_AndroidMkEntries", pkg.TestPlatformBootclasspathModule_AndroidMkEntries},

	{"TestPlatformBootclasspathVariant", pkg.TestPlatformBootclasspathVariant},

	{"TestPlatformBootclasspath_ClasspathFragmentPaths", pkg.TestPlatformBootclasspath_ClasspathFragmentPaths},

	{"TestPlatformBootclasspath_Dist", pkg.TestPlatformBootclasspath_Dist},

	{"TestPlatformBootclasspath_HiddenAPIMonolithicFiles", pkg.TestPlatformBootclasspath_HiddenAPIMonolithicFiles},

	{"TestPlatformCompatConfig", pkg.TestPlatformCompatConfig},

	{"TestPlatformSystemServerClasspathModule_AndroidMkEntries", pkg.TestPlatformSystemServerClasspathModule_AndroidMkEntries},

	{"TestPlatformSystemServerClasspathVariant", pkg.TestPlatformSystemServerClasspathVariant},

	{"TestPlatformSystemServerClasspath_ClasspathFragmentPaths", pkg.TestPlatformSystemServerClasspath_ClasspathFragmentPaths},

	{"TestPlugin", pkg.TestPlugin},

	{"TestPluginGeneratesApi", pkg.TestPluginGeneratesApi},

	{"TestPrebuiltApis_SystemModulesCreation", pkg.TestPrebuiltApis_SystemModulesCreation},

	{"TestPrebuiltApis_WithExtensions", pkg.TestPrebuiltApis_WithExtensions},

	{"TestPrebuiltBootclasspathFragment_UnknownImageName", pkg.TestPrebuiltBootclasspathFragment_UnknownImageName},

	{"TestPrebuiltStubsSources", pkg.TestPrebuiltStubsSources},

	{"TestPrebuilts", pkg.TestPrebuilts},

	{"TestPrepareForTestWithJavaDefaultModules", pkg.TestPrepareForTestWithJavaDefaultModules},

	{"TestPublicDroidstubs", pkg.TestPublicDroidstubs},

	{"TestR8", pkg.TestR8},

	{"TestRequestV4SigningFlag", pkg.TestRequestV4SigningFlag},

	{"TestRequired", pkg.TestRequired},

	{"TestResourceDirs", pkg.TestResourceDirs},

	{"TestResources", pkg.TestResources},

	{"TestRuntimeResourceOverlay", pkg.TestRuntimeResourceOverlay},

	{"TestRuntimeResourceOverlayPartition", pkg.TestRuntimeResourceOverlayPartition},

	{"TestRuntimeResourceOverlay_JavaDefaults", pkg.TestRuntimeResourceOverlay_JavaDefaults},

	{"TestSdkLibrary_CheckMinSdkVersion", pkg.TestSdkLibrary_CheckMinSdkVersion},

	{"TestSdkVersionByPartition", pkg.TestSdkVersionByPartition},

	{"TestSharding", pkg.TestSharding},

	{"TestSimple", pkg.TestSimple},

	{"TestStl", pkg.TestStl},

	{"TestSystemDroidstubs", pkg.TestSystemDroidstubs},

	{"TestSystemServerClasspathFragmentWithoutContents", pkg.TestSystemServerClasspathFragmentWithoutContents},

	{"TestTargetSdkVersionManifestFixer", pkg.TestTargetSdkVersionManifestFixer},

	{"TestTest", pkg.TestTest},

	{"TestTurbine", pkg.TestTurbine},

	{"TestUncompressDex", pkg.TestUncompressDex},

	{"TestUpdatableApps", pkg.TestUpdatableApps},

	{"TestUpdatableApps_ErrorIfDepMinSdkVersionIsHigher", pkg.TestUpdatableApps_ErrorIfDepMinSdkVersionIsHigher},

	{"TestUpdatableApps_ErrorIfJniLibDoesntSupportMinSdkVersion", pkg.TestUpdatableApps_ErrorIfJniLibDoesntSupportMinSdkVersion},

	{"TestUpdatableApps_JniLibShouldBeBuiltAgainstMinSdkVersion", pkg.TestUpdatableApps_JniLibShouldBeBuiltAgainstMinSdkVersion},

	{"TestUpdatableApps_JniLibsShouldShouldSupportMinSdkVersion", pkg.TestUpdatableApps_JniLibsShouldShouldSupportMinSdkVersion},

	{"TestUpdatableApps_TransitiveDepsShouldSetMinSdkVersion", pkg.TestUpdatableApps_TransitiveDepsShouldSetMinSdkVersion},

	{"TestUsesLibraries", pkg.TestUsesLibraries},

	{"TestVendorAppSdkVersion", pkg.TestVendorAppSdkVersion},

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
	return "android/soong/java"
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
