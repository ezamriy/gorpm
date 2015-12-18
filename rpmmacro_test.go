/* -*- mode:go; coding:utf-8; -*-
 * author: Eugene Zamriy <eugene@zamriy.info>
 * created: 2015-12-18
 * description: rpmmacro bindings tests.
 */

package rpm_test

import (
	"testing"

	"github.com/ezamriy/gorpm"
)

func TestExpand(t *testing.T) {
	macro := "%{!?unknown_macro:it_works}"
	want := "it_works"
	got := rpm.Expand(macro)
	if got != want {
		t.Errorf("Expand(%q) == %q, expected %q", macro, got, want)
	}
}

func TestExpandNumeric(t *testing.T) {
	cases := []struct {
		macro string
		want  int
	}{
		{"%{!?unknown_macro:Y}", 1},
		{"%{!?unknown_macro:N}", 0},
		{"%{!?unknown_macro:12}", 12},
	}
	for _, c := range cases {
		got := rpm.ExpandNumeric(c.macro)
		if got != c.want {
			t.Errorf("ExpandNumeric(%q) == %d, expected %d", c.macro, got, c.want)
		}
	}
}
