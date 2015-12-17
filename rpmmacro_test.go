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
