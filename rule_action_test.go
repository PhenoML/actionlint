package actionlint

import "testing"

func TestRuleActionCheckActionUsesFormat(t *testing.T) {
	tests := []struct {
		uses string
		ok   bool
	}{
		{"actions/checkout@v4", true},
		{"./path/to/action", true},
		{"$/path/to/action", true},
		{"docker://alpine:3.20", true},
		{"actions/checkout", false},
		{"checkout@v4", false},
		{"./path/to/action@v1", true},
		{"$/path/to/action@v1", false},
	}

	for _, tc := range tests {
		t.Run(tc.uses, func(t *testing.T) {
			r := NewRuleAction(newNullLocalActionsCache(nil))
			err := r.VisitStep(&Step{
				Exec: &ExecAction{
					Uses: &String{Value: tc.uses, Pos: &Pos{}},
				},
			})
			if err != nil {
				t.Fatal(err)
			}
			errs := r.Errs()
			if tc.ok && len(errs) > 0 {
				t.Fatalf("unexpected errors: %v", errs)
			}
			if !tc.ok && len(errs) == 0 {
				t.Fatal("expected an error")
			}
		})
	}
}
