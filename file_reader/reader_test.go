package file_reader_test

import (
	"testing"

	"github.com/cjcodell1/tint/file_reader"
)

type readAllTest struct {
	path string
	expect string
	isErrNil bool
}

var readAllTests = []readAllTest {
	{"examples/file1", "", true},
	{"examples/file2", "abc", true},
	{"examples/file3", "a\nb\nc", true},
	{"examples/file4", "AaBbCc", true},
	// cannot run the next test because git will not add/commit an unreadable file
	// {"examples/file5", "", false}, // cannot read
	{"examples/file6", "", false}, // file does not exist
	{"examples/file7", "\n", true},
}

func TestReadAll(t *testing.T) {
	for _, tc := range readAllTests {
		got, gotErr := file_reader.ReadAll(tc.path)
		if (got != tc.expect) || (tc.isErrNil && (gotErr != nil)) {
			var expectErr string
			if tc.isErrNil {
				expectErr = "nil"
			} else {
				expectErr = "non-nil"
			}

			if gotErr == nil {
				t.Errorf("ReadAll(%s) == %s, %s != %s, %s", tc.path, got, "nil", tc.expect, expectErr)
			} else {
				t.Errorf("ReadAll(%s) == %s, %s != %s, %s", tc.path, got, gotErr.Error(), tc.expect, expectErr)
			}
		}
	}
}
