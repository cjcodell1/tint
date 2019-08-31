package yaml_test

import (
	"testing"

	"github.com/cjcodell1/tint/builder/yaml"
)

type buildTest struct {
	path     string
	machine string
	isErrNil bool
}

var buildTests = []buildTest{
	{"examples/config1.yaml", "tm", true},
	{"examples/config2.yaml", "tm", true},
	{"examples/config3.yaml", "tm", true},
	{"examples/config4.yaml", "tm", true},
}

func TestBuild(t *testing.T) {
	for _, tc := range buildTests {
		_, gotErr := yaml.Build(tc.path, tc.machine)
		if tc.isErrNil && (gotErr != nil) {
			var expectErr string
			if tc.isErrNil {
				expectErr = "nil"
			} else {
				expectErr = "non-nil"
			}

			if gotErr == nil {
				t.Errorf("Build(%s) == some_tm, %s != some_tm, %s", tc.path, "nil", expectErr)
			} else {
				t.Errorf("Build(%s) == some_tm, %s != some_tm, %s", tc.path, gotErr.Error(), expectErr)
			}

		}
	}
}
