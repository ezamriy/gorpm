/* -*- mode:go; coding:utf-8; -*-
 * author: Eugene Zamriy <eugene@zamriy.info>
 * created: 2015-12-18
 * description: rpmlib bindings tests.
 */

package rpm_test

import (
	"testing"

	"github.com/ezamriy/gorpm"
)

func TestVercmp(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{"1.5.1", "1.5.0", 1},
		{"2.4.99", "2.4.99-1", -1},
		{"0.rc1.7.fc23", "0.rc1.7.fc23", 0},
	}
	for _, c := range cases {
		got := rpm.Vercmp(c.a, c.b)
		if got != c.want {
			t.Errorf("Vercmp(%q, %q) == %d, expected %d", c.a, c.b, got, c.want)
		}
	}
}
