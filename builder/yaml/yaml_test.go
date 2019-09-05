package yaml_test

import (
	"testing"

	"github.com/cjcodell1/tint/builder/yaml"
)

type buildTest struct {
	path     string
	machine string
	err error
}

var buildTests = []buildTest{
	//{"tm_examples/config1.yaml", "tm", nil},
	//{"tm_examples/config2.yaml", "tm", nil},
	//{"tm_examples/config3.yaml", "tm", nil},
	//{"tm_examples/config4.yaml", "tm", nil},

	{"dfa_examples/config1.yaml", "dfa", nil},
	{"dfa_examples/config2.yaml", "dfa", nil},
	{"dfa_examples/config3.yaml", "dfa", nil},
}

func TestBuild(t *testing.T) {
	for _, tc := range buildTests {
		_, err := yaml.Build(tc.path, tc.machine)
			if err != tc.err {
				t.Errorf("Build(%s, %s) == some_machine, %s != some_machine, %s", tc.path, tc.machine, err, tc.err)
		}
	}
}
