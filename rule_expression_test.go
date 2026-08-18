package actionlint

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRuleExpressionGetActionOutputsTypeSkipsExpressionSelfRepositorySyntax(t *testing.T) {
	proj := &Project{filepath.Join("testdata", "action_metadata"), nil}
	rule := NewRuleExpression(NewLocalActionsCache(proj, nil), NewLocalReusableWorkflowCache(nil, "", nil))

	have := rule.getActionOutputsType(&String{Value: "$/${{ matrix.action }}", Pos: &Pos{}})
	want := NewMapObjectType(StringType{})

	if diff := cmp.Diff(want, have); diff != "" {
		t.Fatal(diff)
	}
	if errs := rule.Errs(); len(errs) > 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
}
