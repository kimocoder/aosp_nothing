
package main

import (
	"io"

	"os"

	"reflect"
	"regexp"
	"testing"
	"time"

	pkg "github.com/google/blueprint/pathtools"
)

var t = []testing.InternalTest{

	{"TestFs_IsDir", pkg.TestFs_IsDir},

	{"TestFs_ListDirsRecursiveDontFollowSymlinks", pkg.TestFs_ListDirsRecursiveDontFollowSymlinks},

	{"TestFs_ListDirsRecursiveFollowSymlinks", pkg.TestFs_ListDirsRecursiveFollowSymlinks},

	{"TestFs_Lstat", pkg.TestFs_Lstat},

	{"TestFs_Readlink", pkg.TestFs_Readlink},

	{"TestFs_Stat", pkg.TestFs_Stat},

	{"TestGlob", pkg.TestGlob},

	{"TestGlobDontFollowDanglingSymlinks", pkg.TestGlobDontFollowDanglingSymlinks},

	{"TestGlobDontFollowSymlinks", pkg.TestGlobDontFollowSymlinks},

	{"TestGlobEscapes", pkg.TestGlobEscapes},

	{"TestGlobFollowDanglingSymlinks", pkg.TestGlobFollowDanglingSymlinks},

	{"TestGlobSymlinks", pkg.TestGlobSymlinks},

	{"TestLists_ReplaceExtension", pkg.TestLists_ReplaceExtension},

	{"TestMatch", pkg.TestMatch},

	{"TestMockFs_followSymlinks", pkg.TestMockFs_followSymlinks},

	{"TestMockFs_glob", pkg.TestMockFs_glob},

	{"TestMockGlob", pkg.TestMockGlob},

	{"TestMockGlobDontFollowDanglingSymlinks", pkg.TestMockGlobDontFollowDanglingSymlinks},

	{"TestMockGlobDontFollowSymlinks", pkg.TestMockGlobDontFollowSymlinks},

	{"TestMockGlobEscapes", pkg.TestMockGlobEscapes},

	{"TestMockGlobFollowDanglingSymlinks", pkg.TestMockGlobFollowDanglingSymlinks},

	{"TestMockGlobSymlinks", pkg.TestMockGlobSymlinks},

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
	return "github.com/google/blueprint/pathtools"
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
