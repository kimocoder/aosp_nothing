
package main

import (
	"io"

	"reflect"
	"regexp"
	"testing"
	"time"

	pkg "android/soong/bp2build"
)

var t = []testing.InternalTest{

	{"TestAllowlistingBp2buildTargetsExplicitly", pkg.TestAllowlistingBp2buildTargetsExplicitly},

	{"TestAllowlistingBp2buildTargetsWithConfig", pkg.TestAllowlistingBp2buildTargetsWithConfig},

	{"TestAndroidAppAllSupportedFields", pkg.TestAndroidAppAllSupportedFields},

	{"TestAndroidAppArchVariantSrcs", pkg.TestAndroidAppArchVariantSrcs},

	{"TestAndroidAppCertificateSimple", pkg.TestAndroidAppCertificateSimple},

	{"TestApexBundleCompileMultilib32", pkg.TestApexBundleCompileMultilib32},

	{"TestApexBundleCompileMultilib64", pkg.TestApexBundleCompileMultilib64},

	{"TestApexBundleCompileMultilibBoth", pkg.TestApexBundleCompileMultilibBoth},

	{"TestApexBundleCompileMultilibFirst", pkg.TestApexBundleCompileMultilibFirst},

	{"TestApexBundleDefaultPropertyValues", pkg.TestApexBundleDefaultPropertyValues},

	{"TestApexBundleHasBazelModuleProps", pkg.TestApexBundleHasBazelModuleProps},

	{"TestApexBundleSimple", pkg.TestApexBundleSimple},

	{"TestApexKeySimple", pkg.TestApexKeySimple},

	{"TestBasicCcBinary", pkg.TestBasicCcBinary},

	{"TestCCLibraryNoCrtArchVariant", pkg.TestCCLibraryNoCrtArchVariant},

	{"TestCCLibraryNoCrtFalse", pkg.TestCCLibraryNoCrtFalse},

	{"TestCCLibraryNoCrtTrue", pkg.TestCCLibraryNoCrtTrue},

	{"TestCCLibraryNoLibCrtArchAndTargetVariant", pkg.TestCCLibraryNoLibCrtArchAndTargetVariant},

	{"TestCCLibraryNoLibCrtArchAndTargetVariantConflict", pkg.TestCCLibraryNoLibCrtArchAndTargetVariantConflict},

	{"TestCCLibraryNoLibCrtArchVariant", pkg.TestCCLibraryNoLibCrtArchVariant},

	{"TestCCLibraryNoLibCrtFalse", pkg.TestCCLibraryNoLibCrtFalse},

	{"TestCCLibraryNoLibCrtTrue", pkg.TestCCLibraryNoLibCrtTrue},

	{"TestCcBinaryDoNotDistinguishBetweenDepsAndImplementationDeps", pkg.TestCcBinaryDoNotDistinguishBetweenDepsAndImplementationDeps},

	{"TestCcBinaryNo_libcrtTests", pkg.TestCcBinaryNo_libcrtTests},

	{"TestCcBinaryNocrtTests", pkg.TestCcBinaryNocrtTests},

	{"TestCcBinaryPropertiesToFeatures", pkg.TestCcBinaryPropertiesToFeatures},

	{"TestCcBinarySharedProto", pkg.TestCcBinarySharedProto},

	{"TestCcBinarySplitSrcsByLang", pkg.TestCcBinarySplitSrcsByLang},

	{"TestCcBinaryStaticProto", pkg.TestCcBinaryStaticProto},

	{"TestCcBinaryVersionScript", pkg.TestCcBinaryVersionScript},

	{"TestCcBinaryWithLinkStatic", pkg.TestCcBinaryWithLinkStatic},

	{"TestCcBinaryWithSharedLdflagDisableFeature", pkg.TestCcBinaryWithSharedLdflagDisableFeature},

	{"TestCcLibraryConfiguredVersionScript", pkg.TestCcLibraryConfiguredVersionScript},

	{"TestCcLibraryCppFlagsGoesIntoCopts", pkg.TestCcLibraryCppFlagsGoesIntoCopts},

	{"TestCcLibraryCppStdWithGnuExtensions_ConvertsToFeatureAttr", pkg.TestCcLibraryCppStdWithGnuExtensions_ConvertsToFeatureAttr},

	{"TestCcLibraryDeps", pkg.TestCcLibraryDeps},

	{"TestCcLibraryDisabledArchAndTarget", pkg.TestCcLibraryDisabledArchAndTarget},

	{"TestCcLibraryDisabledArchAndTargetWithDefault", pkg.TestCcLibraryDisabledArchAndTargetWithDefault},

	{"TestCcLibraryEscapeLdflags", pkg.TestCcLibraryEscapeLdflags},

	{"TestCcLibraryExcludeLibs", pkg.TestCcLibraryExcludeLibs},

	{"TestCcLibraryExcludeSrcs", pkg.TestCcLibraryExcludeSrcs},

	{"TestCcLibraryFeatures", pkg.TestCcLibraryFeatures},

	{"TestCcLibraryHeadersArchAndTargetExportSystemIncludes", pkg.TestCcLibraryHeadersArchAndTargetExportSystemIncludes},

	{"TestCcLibraryHeadersLoadStatement", pkg.TestCcLibraryHeadersLoadStatement},

	{"TestCcLibraryHeadersNoCrtIgnored", pkg.TestCcLibraryHeadersNoCrtIgnored},

	{"TestCcLibraryHeadersOsSpecficHeaderLibsExportHeaderLibHeaders", pkg.TestCcLibraryHeadersOsSpecficHeaderLibsExportHeaderLibHeaders},

	{"TestCcLibraryHeadersOsSpecificHeader", pkg.TestCcLibraryHeadersOsSpecificHeader},

	{"TestCcLibraryHeadersSimple", pkg.TestCcLibraryHeadersSimple},

	{"TestCcLibraryNonConfiguredVersionScript", pkg.TestCcLibraryNonConfiguredVersionScript},

	{"TestCcLibraryOsSelects", pkg.TestCcLibraryOsSelects},

	{"TestCcLibraryProtoExplicitCanonicalPathFromRoot", pkg.TestCcLibraryProtoExplicitCanonicalPathFromRoot},

	{"TestCcLibraryProtoExportHeaders", pkg.TestCcLibraryProtoExportHeaders},

	{"TestCcLibraryProtoFilegroups", pkg.TestCcLibraryProtoFilegroups},

	{"TestCcLibraryProtoFull", pkg.TestCcLibraryProtoFull},

	{"TestCcLibraryProtoLite", pkg.TestCcLibraryProtoLite},

	{"TestCcLibraryProtoNoCanonicalPathFromRoot", pkg.TestCcLibraryProtoNoCanonicalPathFromRoot},

	{"TestCcLibraryProtoSimple", pkg.TestCcLibraryProtoSimple},

	{"TestCcLibrarySTaticArchMultilibSrcsExcludeSrcs", pkg.TestCcLibrarySTaticArchMultilibSrcsExcludeSrcs},

	{"TestCcLibrarySharedArchSpecificSharedLib", pkg.TestCcLibrarySharedArchSpecificSharedLib},

	{"TestCcLibrarySharedBaseArchOsSpecificSharedLib", pkg.TestCcLibrarySharedBaseArchOsSpecificSharedLib},

	{"TestCcLibrarySharedDisabled", pkg.TestCcLibrarySharedDisabled},

	{"TestCcLibrarySharedLibs", pkg.TestCcLibrarySharedLibs},

	{"TestCcLibrarySharedNoCrtArchVariant", pkg.TestCcLibrarySharedNoCrtArchVariant},

	{"TestCcLibrarySharedNoCrtFalse", pkg.TestCcLibrarySharedNoCrtFalse},

	{"TestCcLibrarySharedNoCrtTrue", pkg.TestCcLibrarySharedNoCrtTrue},

	{"TestCcLibrarySharedOsSpecificSharedLib", pkg.TestCcLibrarySharedOsSpecificSharedLib},

	{"TestCcLibrarySharedProto", pkg.TestCcLibrarySharedProto},

	{"TestCcLibrarySharedSimple", pkg.TestCcLibrarySharedSimple},

	{"TestCcLibrarySharedSimpleExcludeSrcs", pkg.TestCcLibrarySharedSimpleExcludeSrcs},

	{"TestCcLibrarySharedStaticProps", pkg.TestCcLibrarySharedStaticProps},

	{"TestCcLibrarySharedStaticPropsInArch", pkg.TestCcLibrarySharedStaticPropsInArch},

	{"TestCcLibrarySharedStaticPropsWithMixedSources", pkg.TestCcLibrarySharedStaticPropsWithMixedSources},

	{"TestCcLibrarySharedStrip", pkg.TestCcLibrarySharedStrip},

	{"TestCcLibrarySharedStubs", pkg.TestCcLibrarySharedStubs},

	{"TestCcLibrarySharedUseVersionLib", pkg.TestCcLibrarySharedUseVersionLib},

	{"TestCcLibrarySharedVersionScript", pkg.TestCcLibrarySharedVersionScript},

	{"TestCcLibrarySimple", pkg.TestCcLibrarySimple},

	{"TestCcLibrarySpacesInCopts", pkg.TestCcLibrarySpacesInCopts},

	{"TestCcLibraryStaticArchSpecificStaticLib", pkg.TestCcLibraryStaticArchSpecificStaticLib},

	{"TestCcLibraryStaticArchSrcsExcludeSrcsGeneratedFiles", pkg.TestCcLibraryStaticArchSrcsExcludeSrcsGeneratedFiles},

	{"TestCcLibraryStaticBaseArchOsSpecificStaticLib", pkg.TestCcLibraryStaticBaseArchOsSpecificStaticLib},

	{"TestCcLibraryStaticDisabledForSomeArch", pkg.TestCcLibraryStaticDisabledForSomeArch},

	{"TestCcLibraryStaticExportIncludeDir", pkg.TestCcLibraryStaticExportIncludeDir},

	{"TestCcLibraryStaticExportSystemIncludeDir", pkg.TestCcLibraryStaticExportSystemIncludeDir},

	{"TestCcLibraryStaticFourArchExcludeSrcs", pkg.TestCcLibraryStaticFourArchExcludeSrcs},

	{"TestCcLibraryStaticGeneratedHeadersAllPartitions", pkg.TestCcLibraryStaticGeneratedHeadersAllPartitions},

	{"TestCcLibraryStaticGetTargetProperties", pkg.TestCcLibraryStaticGetTargetProperties},

	{"TestCcLibraryStaticIncludeBuildDirectoryDisabled", pkg.TestCcLibraryStaticIncludeBuildDirectoryDisabled},

	{"TestCcLibraryStaticIncludeBuildDirectoryEnabled", pkg.TestCcLibraryStaticIncludeBuildDirectoryEnabled},

	{"TestCcLibraryStaticLoadStatement", pkg.TestCcLibraryStaticLoadStatement},

	{"TestCcLibraryStaticManyIncludeDirs", pkg.TestCcLibraryStaticManyIncludeDirs},

	{"TestCcLibraryStaticMultipleDepSameName", pkg.TestCcLibraryStaticMultipleDepSameName},

	{"TestCcLibraryStaticOneArchEmpty", pkg.TestCcLibraryStaticOneArchEmpty},

	{"TestCcLibraryStaticOneArchEmptyOtherSet", pkg.TestCcLibraryStaticOneArchEmptyOtherSet},

	{"TestCcLibraryStaticOneArchSrcs", pkg.TestCcLibraryStaticOneArchSrcs},

	{"TestCcLibraryStaticOneArchSrcsExcludeSrcs", pkg.TestCcLibraryStaticOneArchSrcsExcludeSrcs},

	{"TestCcLibraryStaticOneMultilibSrcsExcludeSrcs", pkg.TestCcLibraryStaticOneMultilibSrcsExcludeSrcs},

	{"TestCcLibraryStaticOsSpecificStaticLib", pkg.TestCcLibraryStaticOsSpecificStaticLib},

	{"TestCcLibraryStaticProductVariableArchSpecificSelects", pkg.TestCcLibraryStaticProductVariableArchSpecificSelects},

	{"TestCcLibraryStaticProductVariableSelects", pkg.TestCcLibraryStaticProductVariableSelects},

	{"TestCcLibraryStaticProductVariableStringReplacement", pkg.TestCcLibraryStaticProductVariableStringReplacement},

	{"TestCcLibraryStaticProto", pkg.TestCcLibraryStaticProto},

	{"TestCcLibraryStaticSimple", pkg.TestCcLibraryStaticSimple},

	{"TestCcLibraryStaticSimpleExcludeSrcs", pkg.TestCcLibraryStaticSimpleExcludeSrcs},

	{"TestCcLibraryStaticStdInFlags", pkg.TestCcLibraryStaticStdInFlags},

	{"TestCcLibraryStaticSubpackage", pkg.TestCcLibraryStaticSubpackage},

	{"TestCcLibraryStaticTwoArchExcludeSrcs", pkg.TestCcLibraryStaticTwoArchExcludeSrcs},

	{"TestCcLibraryStaticTwoMultilibSrcsExcludeSrcs", pkg.TestCcLibraryStaticTwoMultilibSrcsExcludeSrcs},

	{"TestCcLibraryStaticUseVersionLib", pkg.TestCcLibraryStaticUseVersionLib},

	{"TestCcLibraryStrip", pkg.TestCcLibraryStrip},

	{"TestCcLibraryStripWithArch", pkg.TestCcLibraryStripWithArch},

	{"TestCcLibraryStubs", pkg.TestCcLibraryStubs},

	{"TestCcLibraryTrimmedLdAndroid", pkg.TestCcLibraryTrimmedLdAndroid},

	{"TestCcLibraryWholeStaticLibsAlwaysLink", pkg.TestCcLibraryWholeStaticLibsAlwaysLink},

	{"TestCcLibrary_SystemSharedLibsBionicEmpty", pkg.TestCcLibrary_SystemSharedLibsBionicEmpty},

	{"TestCcLibrary_SystemSharedLibsLinuxBionicEmpty", pkg.TestCcLibrary_SystemSharedLibsLinuxBionicEmpty},

	{"TestCcLibrary_SystemSharedLibsRootEmpty", pkg.TestCcLibrary_SystemSharedLibsRootEmpty},

	{"TestCcLibrary_SystemSharedLibsSharedAndRoot", pkg.TestCcLibrary_SystemSharedLibsSharedAndRoot},

	{"TestCcLibrary_SystemSharedLibsSharedBionicEmpty", pkg.TestCcLibrary_SystemSharedLibsSharedBionicEmpty},

	{"TestCcLibrary_SystemSharedLibsSharedEmpty", pkg.TestCcLibrary_SystemSharedLibsSharedEmpty},

	{"TestCcLibrary_SystemSharedLibsStaticEmpty", pkg.TestCcLibrary_SystemSharedLibsStaticEmpty},

	{"TestCcLibrarystatic_SystemSharedLibUsedAsDep", pkg.TestCcLibrarystatic_SystemSharedLibUsedAsDep},

	{"TestCcObjectCcObjetDepsInObjs", pkg.TestCcObjectCcObjetDepsInObjs},

	{"TestCcObjectCflagsFourArch", pkg.TestCcObjectCflagsFourArch},

	{"TestCcObjectCflagsOneArch", pkg.TestCcObjectCflagsOneArch},

	{"TestCcObjectDefaults", pkg.TestCcObjectDefaults},

	{"TestCcObjectDepsAndLinkerScriptSelects", pkg.TestCcObjectDepsAndLinkerScriptSelects},

	{"TestCcObjectIncludeBuildDirFalse", pkg.TestCcObjectIncludeBuildDirFalse},

	{"TestCcObjectLinkerScript", pkg.TestCcObjectLinkerScript},

	{"TestCcObjectProductVariable", pkg.TestCcObjectProductVariable},

	{"TestCcObjectSelectOnLinuxAndBionicArchs", pkg.TestCcObjectSelectOnLinuxAndBionicArchs},

	{"TestCcObjectSimple", pkg.TestCcObjectSimple},

	{"TestCombineBuildFilesBp2buildTargets", pkg.TestCombineBuildFilesBp2buildTargets},

	{"TestCommonBp2BuildModuleAttrs", pkg.TestCommonBp2BuildModuleAttrs},

	{"TestConvertManyModulesFull", pkg.TestConvertManyModulesFull},

	{"TestCreateBazelFiles_Bp2Build_CreatesDefaultFiles", pkg.TestCreateBazelFiles_Bp2Build_CreatesDefaultFiles},

	{"TestCreateBazelFiles_QueryView_AddsTopLevelFiles", pkg.TestCreateBazelFiles_QueryView_AddsTopLevelFiles},

	{"TestFilegroupSameNameAsFile_MultipleFiles", pkg.TestFilegroupSameNameAsFile_MultipleFiles},

	{"TestFilegroupSameNameAsFile_OneFile", pkg.TestFilegroupSameNameAsFile_OneFile},

	{"TestGenerateBazelTargetModules", pkg.TestGenerateBazelTargetModules},

	{"TestGenerateBazelTargetModules_OneToMany_LoadedFromStarlark", pkg.TestGenerateBazelTargetModules_OneToMany_LoadedFromStarlark},

	{"TestGenerateModuleRuleShims", pkg.TestGenerateModuleRuleShims},

	{"TestGenerateSoongModuleBzl", pkg.TestGenerateSoongModuleBzl},

	{"TestGenerateSoongModuleTargets", pkg.TestGenerateSoongModuleTargets},

	{"TestGenruleBp2BuildInlinesDefaults", pkg.TestGenruleBp2BuildInlinesDefaults},

	{"TestGenruleCliVariableReplacement", pkg.TestGenruleCliVariableReplacement},

	{"TestGenruleLocationLabelShouldSubstituteFirstToolLabel", pkg.TestGenruleLocationLabelShouldSubstituteFirstToolLabel},

	{"TestGenruleLocationsAbsoluteLabel", pkg.TestGenruleLocationsAbsoluteLabel},

	{"TestGenruleLocationsLabel", pkg.TestGenruleLocationsLabel},

	{"TestGenruleLocationsLabelShouldSubstituteFirstToolLabel", pkg.TestGenruleLocationsLabelShouldSubstituteFirstToolLabel},

	{"TestGenruleSrcsLocationsAbsoluteLabel", pkg.TestGenruleSrcsLocationsAbsoluteLabel},

	{"TestGenruleWithoutToolsOrToolFiles", pkg.TestGenruleWithoutToolsOrToolFiles},

	{"TestGlobExcludeSrcs", pkg.TestGlobExcludeSrcs},

	{"TestJavaBinaryHost", pkg.TestJavaBinaryHost},

	{"TestJavaBinaryHostRuntimeDeps", pkg.TestJavaBinaryHostRuntimeDeps},

	{"TestJavaImportArchVariant", pkg.TestJavaImportArchVariant},

	{"TestJavaImportMinimal", pkg.TestJavaImportMinimal},

	{"TestJavaLibrary", pkg.TestJavaLibrary},

	{"TestJavaLibraryConvertsStaticLibsToDepsAndExports", pkg.TestJavaLibraryConvertsStaticLibsToDepsAndExports},

	{"TestJavaLibraryConvertsStaticLibsToExportsIfNoSrcs", pkg.TestJavaLibraryConvertsStaticLibsToExportsIfNoSrcs},

	{"TestJavaLibraryErrorproneJavacflagsEnabledManually", pkg.TestJavaLibraryErrorproneJavacflagsEnabledManually},

	{"TestJavaLibraryErrorproneJavacflagsErrorproneDisabledByDefault", pkg.TestJavaLibraryErrorproneJavacflagsErrorproneDisabledByDefault},

	{"TestJavaLibraryErrorproneJavacflagsErrorproneDisabledManually", pkg.TestJavaLibraryErrorproneJavacflagsErrorproneDisabledManually},

	{"TestJavaLibraryFailsToConvertLibsWithNoSrcs", pkg.TestJavaLibraryFailsToConvertLibsWithNoSrcs},

	{"TestJavaLibraryHost", pkg.TestJavaLibraryHost},

	{"TestJavaLibraryLogTags", pkg.TestJavaLibraryLogTags},

	{"TestJavaLibraryPlugins", pkg.TestJavaLibraryPlugins},

	{"TestJavaPlugin", pkg.TestJavaPlugin},

	{"TestJavaPluginNoSrcs", pkg.TestJavaPluginNoSrcs},

	{"TestJavaProto", pkg.TestJavaProto},

	{"TestJavaProtoDefault", pkg.TestJavaProtoDefault},

	{"TestLibcryptoHashInjection", pkg.TestLibcryptoHashInjection},

	{"TestLoadStatements", pkg.TestLoadStatements},

	{"TestMinimalAndroidApp", pkg.TestMinimalAndroidApp},

	{"TestModuleTypeBp2Build", pkg.TestModuleTypeBp2Build},

	{"TestPrebuiltEtcArchAndTargetVariant", pkg.TestPrebuiltEtcArchAndTargetVariant},

	{"TestPrebuiltEtcArchVariant", pkg.TestPrebuiltEtcArchVariant},

	{"TestPrebuiltEtcNoSubdir", pkg.TestPrebuiltEtcNoSubdir},

	{"TestPrebuiltEtcSimple", pkg.TestPrebuiltEtcSimple},

	{"TestPrebuiltLibraryAdditionalAttrs", pkg.TestPrebuiltLibraryAdditionalAttrs},

	{"TestPrebuiltLibrarySharedAndStaticStanzas", pkg.TestPrebuiltLibrarySharedAndStaticStanzas},

	{"TestPrebuiltLibrarySharedStanzaFails", pkg.TestPrebuiltLibrarySharedStanzaFails},

	{"TestPrebuiltLibraryStaticAndSharedSimple", pkg.TestPrebuiltLibraryStaticAndSharedSimple},

	{"TestPrebuiltLibraryStaticStanzaFails", pkg.TestPrebuiltLibraryStaticStanzaFails},

	{"TestPrebuiltLibraryWithArchVariance", pkg.TestPrebuiltLibraryWithArchVariance},

	{"TestPrebuiltUsrShareSimple", pkg.TestPrebuiltUsrShareSimple},

	{"TestPythonArchVariance", pkg.TestPythonArchVariance},

	{"TestPythonBinaryHostArchVariance", pkg.TestPythonBinaryHostArchVariance},

	{"TestPythonBinaryHostPy2", pkg.TestPythonBinaryHostPy2},

	{"TestPythonBinaryHostPy3", pkg.TestPythonBinaryHostPy3},

	{"TestPythonBinaryHostSimple", pkg.TestPythonBinaryHostSimple},

	{"TestShBinaryDefaults", pkg.TestShBinaryDefaults},

	{"TestShBinaryLoadStatement", pkg.TestShBinaryLoadStatement},

	{"TestShBinarySimple", pkg.TestShBinarySimple},

	{"TestSharedPrebuiltLibrary", pkg.TestSharedPrebuiltLibrary},

	{"TestSharedPrebuiltLibrarySharedStanzaFails", pkg.TestSharedPrebuiltLibrarySharedStanzaFails},

	{"TestSharedPrebuiltLibraryWithArchVariance", pkg.TestSharedPrebuiltLibraryWithArchVariance},

	{"TestSimplePythonLib", pkg.TestSimplePythonLib},

	{"TestSoongConfigModuleType", pkg.TestSoongConfigModuleType},

	{"TestSoongConfigModuleTypeImport", pkg.TestSoongConfigModuleTypeImport},

	{"TestSoongConfigModuleType_Defaults", pkg.TestSoongConfigModuleType_Defaults},

	{"TestSoongConfigModuleType_Defaults_Another", pkg.TestSoongConfigModuleType_Defaults_Another},

	{"TestSoongConfigModuleType_Defaults_MultipleNamespaces", pkg.TestSoongConfigModuleType_Defaults_MultipleNamespaces},

	{"TestSoongConfigModuleType_Defaults_SingleNamespace", pkg.TestSoongConfigModuleType_Defaults_SingleNamespace},

	{"TestSoongConfigModuleType_Defaults_UnusedProps", pkg.TestSoongConfigModuleType_Defaults_UnusedProps},

	{"TestSoongConfigModuleType_MultipleDefaults_SingleNamespace", pkg.TestSoongConfigModuleType_MultipleDefaults_SingleNamespace},

	{"TestSoongConfigModuleType_ProductVariableConfigOverridesEnable", pkg.TestSoongConfigModuleType_ProductVariableConfigOverridesEnable},

	{"TestSoongConfigModuleType_ProductVariableConfigWithPlatformConfig", pkg.TestSoongConfigModuleType_ProductVariableConfigWithPlatformConfig},

	{"TestSoongConfigModuleType_ProductVariableIgnoredIfEnabledByDefault", pkg.TestSoongConfigModuleType_ProductVariableIgnoredIfEnabledByDefault},

	{"TestSoongConfigModuleType_StringAndBoolVar", pkg.TestSoongConfigModuleType_StringAndBoolVar},

	{"TestSoongConfigModuleType_StringVar", pkg.TestSoongConfigModuleType_StringVar},

	{"TestSoongConfigModuleType_StringVar_LabelListDeps", pkg.TestSoongConfigModuleType_StringVar_LabelListDeps},

	{"TestStaticLibrary_SystemSharedLibsBionic", pkg.TestStaticLibrary_SystemSharedLibsBionic},

	{"TestStaticLibrary_SystemSharedLibsBionicEmpty", pkg.TestStaticLibrary_SystemSharedLibsBionicEmpty},

	{"TestStaticLibrary_SystemSharedLibsLinuxBionicEmpty", pkg.TestStaticLibrary_SystemSharedLibsLinuxBionicEmpty},

	{"TestStaticLibrary_SystemSharedLibsLinuxRootAndLinuxBionic", pkg.TestStaticLibrary_SystemSharedLibsLinuxRootAndLinuxBionic},

	{"TestStaticLibrary_SystemSharedLibsRootEmpty", pkg.TestStaticLibrary_SystemSharedLibsRootEmpty},

	{"TestStaticLibrary_SystemSharedLibsStaticEmpty", pkg.TestStaticLibrary_SystemSharedLibsStaticEmpty},

	{"TestStaticPrebuiltLibrary", pkg.TestStaticPrebuiltLibrary},

	{"TestStaticPrebuiltLibraryStaticStanzaFails", pkg.TestStaticPrebuiltLibraryStaticStanzaFails},

	{"TestStaticPrebuiltLibraryWithArchVariance", pkg.TestStaticPrebuiltLibraryWithArchVariance},

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
	return "android/soong/bp2build"
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
