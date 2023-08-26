
package main

import (
	"io"

	"reflect"
	"regexp"
	"testing"
	"time"

	pkg "android/soong/android"
)

var t = []testing.InternalTest{

	{"TestAndroidMkSingleton_PassesUpdatedAndroidMkDataToCustomCallback", pkg.TestAndroidMkSingleton_PassesUpdatedAndroidMkDataToCustomCallback},

	{"TestArchMutator", pkg.TestArchMutator},

	{"TestArchMutatorNativeBridge", pkg.TestArchMutatorNativeBridge},

	{"TestArchProperties", pkg.TestArchProperties},

	{"TestBaseDirForInstallPath", pkg.TestBaseDirForInstallPath},

	{"TestBazelToolchainVars", pkg.TestBazelToolchainVars},

	{"TestBootJarsMarshaling", pkg.TestBootJarsMarshaling},

	{"TestBp2BuildAllowlist", pkg.TestBp2BuildAllowlist},

	{"TestCSuiteConfig", pkg.TestCSuiteConfig},

	{"TestConfiguredJarList", pkg.TestConfiguredJarList},

	{"TestConsistentNamespaceNames", pkg.TestConsistentNamespaceNames},

	{"TestConvertAllModulesInPackage", pkg.TestConvertAllModulesInPackage},

	{"TestDeclaringNamespaceInNonAndroidBpFile", pkg.TestDeclaringNamespaceInNonAndroidBpFile},

	{"TestDefaultProductVariableMarshaling", pkg.TestDefaultProductVariableMarshaling},

	{"TestDefaults", pkg.TestDefaults},

	{"TestDefaultsAllowMissingDependencies", pkg.TestDefaultsAllowMissingDependencies},

	{"TestDepSet", pkg.TestDepSet},

	{"TestDepSetInvalidOrder", pkg.TestDepSetInvalidOrder},

	{"TestDependingOnBlueprintModuleInRootNamespace", pkg.TestDependingOnBlueprintModuleInRootNamespace},

	{"TestDependingOnModuleByFullyQualifiedReference", pkg.TestDependingOnModuleByFullyQualifiedReference},

	{"TestDependingOnModuleInImportedNamespace", pkg.TestDependingOnModuleInImportedNamespace},

	{"TestDependingOnModuleInNonImportedNamespace", pkg.TestDependingOnModuleInNonImportedNamespace},

	{"TestDependingOnModuleInRootNamespace", pkg.TestDependingOnModuleInRootNamespace},

	{"TestDependingOnModuleInSameNamespace", pkg.TestDependingOnModuleInSameNamespace},

	{"TestDirectorySortedPaths", pkg.TestDirectorySortedPaths},

	{"TestDistErrorChecking", pkg.TestDistErrorChecking},

	{"TestDuplicateSingletonModule", pkg.TestDuplicateSingletonModule},

	{"TestDuplicateStringValueInSoongConfigStringVariable", pkg.TestDuplicateStringValueInSoongConfigStringVariable},

	{"TestErrorDependsOnDisabledModule", pkg.TestErrorDependsOnDisabledModule},

	{"TestExpand", pkg.TestExpand},

	{"TestExpandNinjaEscaped", pkg.TestExpandNinjaEscaped},

	{"TestExpandVars", pkg.TestExpandVars},

	{"TestFilterArchStruct", pkg.TestFilterArchStruct},

	{"TestFilterList", pkg.TestFilterList},

	{"TestFilterListPred", pkg.TestFilterListPred},

	{"TestFinalDepsPhase", pkg.TestFinalDepsPhase},

	{"TestFirstUniqueStrings", pkg.TestFirstUniqueStrings},

	{"TestFixtureDedup", pkg.TestFixtureDedup},

	{"TestFixtureValidateMockFS", pkg.TestFixtureValidateMockFS},

	{"TestGenerateDistContributionsForMake", pkg.TestGenerateDistContributionsForMake},

	{"TestGetDistContributions", pkg.TestGetDistContributions},

	{"TestGetDistForGoals", pkg.TestGetDistForGoals},

	{"TestImplicitlyImportRootNamespace", pkg.TestImplicitlyImportRootNamespace},

	{"TestImportingNonexistentNamespace", pkg.TestImportingNonexistentNamespace},

	{"TestInList", pkg.TestInList},

	{"TestIndexList", pkg.TestIndexList},

	{"TestInstall", pkg.TestInstall},

	{"TestInstallDependencyTag", pkg.TestInstallDependencyTag},

	{"TestInstallKatiEnabled", pkg.TestInstallKatiEnabled},

	{"TestInvokeBazelPopulatesBuildStatements", pkg.TestInvokeBazelPopulatesBuildStatements},

	{"TestInvokeBazelWritesBazelFiles", pkg.TestInvokeBazelWritesBazelFiles},

	{"TestJoinWithPrefix", pkg.TestJoinWithPrefix},

	{"TestLastUniqueStrings", pkg.TestLastUniqueStrings},

	{"TestLicense", pkg.TestLicense},

	{"TestLicenseKind", pkg.TestLicenseKind},

	{"TestLicenses", pkg.TestLicenses},

	{"TestMaybeRel", pkg.TestMaybeRel},

	{"TestMissingVendorConfig", pkg.TestMissingVendorConfig},

	{"TestModuleOptIn", pkg.TestModuleOptIn},

	{"TestModuleString", pkg.TestModuleString},

	{"TestModulesDoReceiveParentNamespace", pkg.TestModulesDoReceiveParentNamespace},

	{"TestMutatorAddMissingDependencies", pkg.TestMutatorAddMissingDependencies},

	{"TestNamespaceImportsNotTransitive", pkg.TestNamespaceImportsNotTransitive},

	{"TestNamespaceNotAtTopOfFile", pkg.TestNamespaceNotAtTopOfFile},

	{"TestNamespacesDontInheritParentNamespaces", pkg.TestNamespacesDontInheritParentNamespaces},

	{"TestNeverallow", pkg.TestNeverallow},

	{"TestNewCustomOnceKey", pkg.TestNewCustomOnceKey},

	{"TestNewOnceKey", pkg.TestNewOnceKey},

	{"TestNinjaDeps", pkg.TestNinjaDeps},

	{"TestNoCreateVariationsInFinalDeps", pkg.TestNoCreateVariationsInFinalDeps},

	{"TestNonExistentPropertyInSoongConfigModule", pkg.TestNonExistentPropertyInSoongConfigModule},

	{"TestOncePerPanic", pkg.TestOncePerPanic},

	{"TestOncePerReentrant", pkg.TestOncePerReentrant},

	{"TestOncePer_Get", pkg.TestOncePer_Get},

	{"TestOncePer_Get_panic", pkg.TestOncePer_Get_panic},

	{"TestOncePer_Get_wait", pkg.TestOncePer_Get_wait},

	{"TestOncePer_Once", pkg.TestOncePer_Once},

	{"TestOncePer_Once2StringSlice", pkg.TestOncePer_Once2StringSlice},

	{"TestOncePer_OnceStringSlice", pkg.TestOncePer_OnceStringSlice},

	{"TestOncePer_Once_wait", pkg.TestOncePer_Once_wait},

	{"TestOptionalPath", pkg.TestOptionalPath},

	{"TestPackage", pkg.TestPackage},

	{"TestPackagingBaseMultiTarget", pkg.TestPackagingBaseMultiTarget},

	{"TestPackagingBaseSingleTarget", pkg.TestPackagingBaseSingleTarget},

	{"TestPackagingWithSkipInstallDeps", pkg.TestPackagingWithSkipInstallDeps},

	{"TestPathDepsMutator", pkg.TestPathDepsMutator},

	{"TestPathForModuleInstall", pkg.TestPathForModuleInstall},

	{"TestPathForModuleInstallRecoveryAsBoot", pkg.TestPathForModuleInstallRecoveryAsBoot},

	{"TestPathForModuleSrc", pkg.TestPathForModuleSrc},

	{"TestPathForSource", pkg.TestPathForSource},

	{"TestPathRelativeToTop", pkg.TestPathRelativeToTop},

	{"TestPathsForModuleSrc", pkg.TestPathsForModuleSrc},

	{"TestPathsForModuleSrc_AllowMissingDependencies", pkg.TestPathsForModuleSrc_AllowMissingDependencies},

	{"TestPrebuilts", pkg.TestPrebuilts},

	{"TestPrefixInList", pkg.TestPrefixInList},

	{"TestPrintfIntoProperty", pkg.TestPrintfIntoProperty},

	{"TestProductConfigAnnotations", pkg.TestProductConfigAnnotations},

	{"TestProductVariables", pkg.TestProductVariables},

	{"TestProductVariablesDefaults", pkg.TestProductVariablesDefaults},

	{"TestRemoveFromList", pkg.TestRemoveFromList},

	{"TestRemoveListFromList", pkg.TestRemoveListFromList},

	{"TestRename", pkg.TestRename},

	{"TestRequestResultsAfterInvokeBazel", pkg.TestRequestResultsAfterInvokeBazel},

	{"TestRuleBuilder", pkg.TestRuleBuilder},

	{"TestRuleBuilderHashInputs", pkg.TestRuleBuilderHashInputs},

	{"TestRuleBuilder_Build", pkg.TestRuleBuilder_Build},

	{"TestSameNameInTwoNamespaces", pkg.TestSameNameInTwoNamespaces},

	{"TestSdkRegistry", pkg.TestSdkRegistry},

	{"TestSdkSpecFrom", pkg.TestSdkSpecFrom},

	{"TestSearchOrder", pkg.TestSearchOrder},

	{"TestSingletonModule", pkg.TestSingletonModule},

	{"TestSoongConfigModule", pkg.TestSoongConfigModule},

	{"TestSortedStringKeys", pkg.TestSortedStringKeys},

	{"TestSortedStringValues", pkg.TestSortedStringValues},

	{"TestSortedUniqueNamedPaths", pkg.TestSortedUniqueNamedPaths},

	{"TestSortedUniqueStringValues", pkg.TestSortedUniqueStringValues},

	{"TestSplitFileExt", pkg.TestSplitFileExt},

	{"TestSplitStringKeepingQuotedSubstring", pkg.TestSplitStringKeepingQuotedSubstring},

	{"TestSrcIsModule", pkg.TestSrcIsModule},

	{"TestSrcIsModuleWithTag", pkg.TestSrcIsModuleWithTag},

	{"TestTwoModulesWithSameNameInSameNamespace", pkg.TestTwoModulesWithSameNameInSameNamespace},

	{"TestTwoNamepacesInSameDir", pkg.TestTwoNamepacesInSameDir},

	{"TestTwoNamespacesCanImportEachOther", pkg.TestTwoNamespacesCanImportEachOther},

	{"TestUnusedSingletonModule", pkg.TestUnusedSingletonModule},

	{"TestUsedProperties", pkg.TestUsedProperties},

	{"TestValidateConfigAnnotations", pkg.TestValidateConfigAnnotations},

	{"TestValidateCorrectBuildParams", pkg.TestValidateCorrectBuildParams},

	{"TestValidateIncorrectBuildParams", pkg.TestValidateIncorrectBuildParams},

	{"TestValidatePath", pkg.TestValidatePath},

	{"TestValidateSafePath", pkg.TestValidateSafePath},

	{"TestVariantSingletonModule", pkg.TestVariantSingletonModule},

	{"TestVisibility", pkg.TestVisibility},

	{"TestVisibilityRuleSet", pkg.TestVisibilityRuleSet},

	{"Test_Shard", pkg.Test_Shard},

	{"Test_firstUnique", pkg.Test_firstUnique},

	{"Test_mergeApexVariations", pkg.Test_mergeApexVariations},

}

var e = []testing.InternalExample{

	
		{"DepSet_ToList_postordered", pkg.ExampleDepSet_ToList_postordered, "[a b c d]\n", false},
	

	
		{"DepSet_ToList_preordered", pkg.ExampleDepSet_ToList_preordered, "[d b a c]\n", false},
	

	
		{"DepSet_ToList_topological", pkg.ExampleDepSet_ToList_topological, "[d b c a]\n", false},
	

	
		{"DepSet_ToSortedList", pkg.ExampleDepSet_ToSortedList, "[a b c d]\n", false},
	

	
		{"OutputPath_InSameDir", pkg.ExampleOutputPath_InSameDir, "out/soong/system/framework/boot.art out/soong/system/framework/oat/arm/boot.vdex\nboot.art oat/arm/boot.vdex\n", false},
	

	
		{"OutputPath_ReplaceExtension", pkg.ExampleOutputPath_ReplaceExtension, "out/soong/system/framework/boot.art out/soong/system/framework/boot.oat\nboot.art boot.oat\n", false},
	

	
		{"RuleBuilder", pkg.ExampleRuleBuilder, "commands: \"ld a.o b.o -o out/soong/linked && echo success\"\ntools: [\"ld\"]\ninputs: [\"a.o\" \"b.o\"]\noutputs: [\"out/soong/linked\"]\n", false},
	

	

	
		{"RuleBuilderCommand_Flag", pkg.ExampleRuleBuilderCommand_Flag, "ls -l\n", false},
	

	
		{"RuleBuilderCommand_FlagForEachArg", pkg.ExampleRuleBuilderCommand_FlagForEachArg, "ls --sort=time --sort=size\n", false},
	

	
		{"RuleBuilderCommand_FlagForEachInput", pkg.ExampleRuleBuilderCommand_FlagForEachInput, "turbine --classpath a.jar --classpath b.jar\n", false},
	

	
		{"RuleBuilderCommand_FlagWithArg", pkg.ExampleRuleBuilderCommand_FlagWithArg, "ls --sort=time\n", false},
	

	
		{"RuleBuilderCommand_FlagWithInput", pkg.ExampleRuleBuilderCommand_FlagWithInput, "java -classpath=a\n", false},
	

	
		{"RuleBuilderCommand_FlagWithInputList", pkg.ExampleRuleBuilderCommand_FlagWithInputList, "java -classpath=a.jar:b.jar\n", false},
	

	
		{"RuleBuilderCommand_FlagWithList", pkg.ExampleRuleBuilderCommand_FlagWithList, "ls --sort=time,size\n", false},
	

	
		{"RuleBuilderCommand_FlagWithRspFileInputList", pkg.ExampleRuleBuilderCommand_FlagWithRspFileInputList, "javac @out/soong/foo.rsp\n", false},
	

	
		{"RuleBuilderCommand_Flags", pkg.ExampleRuleBuilderCommand_Flags, "ls -l -a\n", false},
	

	
		{"RuleBuilderCommand_String", pkg.ExampleRuleBuilderCommand_String, "FOO=foo echo $FOO\n", false},
	

	
		{"RuleBuilder_DeleteTemporaryFiles", pkg.ExampleRuleBuilder_DeleteTemporaryFiles, "commands: \"cp a out/soong/b && cp out/soong/b out/soong/c && rm -f out/soong/b\"\ntools: [\"cp\"]\ninputs: [\"a\"]\noutputs: [\"out/soong/c\"]\n", false},
	

	
		{"RuleBuilder_Installs", pkg.ExampleRuleBuilder_Installs, "rule.Installs().String() = \"out/soong/linked:/bin/linked out/soong/linked:/sbin/linked\"\n", false},
	

	
		{"RuleBuilder_SymlinkOutputs", pkg.ExampleRuleBuilder_SymlinkOutputs, "commands: \"ln -s a.o out/soong/a && cp out/soong/a out/soong/b\"\ntools: [\"ln\"]\ninputs: [\"a.o\"]\noutputs: [\"out/soong/a\" \"out/soong/b\"]\nsymlink_outputs: [\"out/soong/a\" \"out/soong/b\"]\n", false},
	

	
		{"RuleBuilder_Temporary", pkg.ExampleRuleBuilder_Temporary, "commands: \"cp a out/soong/b && cp out/soong/b out/soong/c\"\ntools: [\"cp\"]\ninputs: [\"a\"]\noutputs: [\"out/soong/c\"]\n", false},
	

	
		{"CopyOf", pkg.ExampleCopyOf, "a = [\"-1\" \"2\" \"3\"]\nb = [\"1\" \"2\" \"3\"]\n", false},
	

	
		{"CopyOf_append", pkg.ExampleCopyOf_append, "Without CopyOf:\na = [\"foo\"]\nb = [\"foo\" \"baz\"]\nc = [\"foo\" \"baz\"]\nWith CopyOf:\na = [\"foo\"]\nb = [\"foo\" \"bar\"]\nc = [\"foo\" \"baz\"]\n", false},
	

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
	return "android/soong/android"
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
