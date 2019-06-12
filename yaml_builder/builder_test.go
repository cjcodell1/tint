package yaml_builder_test

import (
    "testing"

    "github.com/cjcodell1/tint/yaml_builder"
)

type buildTest struct {
    path string
    isErrNil bool
}

var buildTests = []buildTest{
    {"examples/config1.yaml", true},
    {"examples/config2.yaml", true},
    {"examples/config3.yaml", true},
    {"examples/config4.yaml", true},
}

func TestBuild(t *testing.T) {
    for _, tc := range buildTests {
        _, gotErr := yaml_builder.Build(tc.path)
        if (tc.isErrNil && (gotErr != nil)) {
            var expectErr string
            if tc.isErrNil {
                expectErr = "nil"
            } else {
                expectErr = "non-nil"
            }

            if gotErr == nil {
                t.Errorf("Build(%s) == some_tm, %s != some_tm, %s", tc.path, "nil", expectErr)
            } else {
                //t.Errorf("Build(%s) == some_tm, %s != some_tm, %s", tc.path, gotErr.Error(), expectErr)
                t.Error(gotErr.Error())
            }

        }
    }
}
