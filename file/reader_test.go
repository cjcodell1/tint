package file_test

import (
	"strings"
	"testing"

	"github.com/cjcodell1/tint/file"
)

type readAllTest struct {
	path     string
	expect   string
	isErrNil bool
}

var readAllTests = []readAllTest{
	{"examples/file1", "", true}, // contains the empty string
	{"examples/file2", "abc\n", true},
	{"examples/file3", "a\nb\nc\n", true},
	{"examples/file4", "AaBbCc\n", true},
	// cannot run the next test because git will not add/commit an unreadable file
	// {"examples/file5", "", false}, // cannot read
	{"examples/file6", "", false}, // file does not exist
	{"examples/file7", "\n\n", true},
	{"examples/file8", "hello\r\nworld\r\n", true}, // dos endings
	{"examples/file9", "\r\n", true},               // also dos endings
}

func TestReadAll(t *testing.T) {
	for _, tc := range readAllTests {
		got, gotErr := file.ReadAll(tc.path)
		if (got != tc.expect) || (tc.isErrNil && (gotErr != nil)) {
			var expectErr string
			if tc.isErrNil {
				expectErr = "nil"
			} else {
				expectErr = "non-nil"
			}

			if gotErr == nil {
				t.Errorf("ReadAll(%s) == %q, %s != %q, %s", tc.path, got, "nil", tc.expect, expectErr)
			} else {
				t.Errorf("ReadAll(%s) == %q, %s != %q, %s", tc.path, got, gotErr.Error(), tc.expect, expectErr)
			}
		}
	}
}

type readLinesTest struct {
	path     string
	expect   []string
	isErrNil bool
}

var readLinesTests = []readLinesTest{
	{"examples/file1", []string{""}, true}, // contains the empty string
	{"examples/file2", []string{"abc"}, true},
	{"examples/file3", []string{"a", "b", "c"}, true},
	{"examples/file4", []string{"AaBbCc"}, true},
	// cannot run the next test because git will not add/commit an unreadable file
	// {"examples/file5", "", false}, // cannot read
	{"examples/file6", []string{""}, false}, // file does not exist
	{"examples/file7", []string{"", ""}, true},
	{"examples/file8", []string{"hello", "world"}, true}, // dos endings
	{"examples/file9", []string{""}, true},               // also dos endings
}

func TestReadLines(t *testing.T) {
	for _, tc := range readLinesTests {
		got, gotErr := file.ReadLines(tc.path)
		if (strings.Join(got, "\n") != strings.Join(tc.expect, "\n")) || (tc.isErrNil && (gotErr != nil)) {
			var expectErr string
			if tc.isErrNil {
				expectErr = "nil"
			} else {
				expectErr = "non-nil"
			}

			if gotErr == nil {
				t.Errorf("Lines(%s) == %q, %s != %q, %s", tc.path, got, "nil", tc.expect, expectErr)
			} else {
				t.Errorf("Lines(%s) == %q, %s != %q, %s", tc.path, got, gotErr.Error(), tc.expect, expectErr)
			}
		}
	}
}
