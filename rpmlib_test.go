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


func TestLabelCompare(t *testing.T) {
	cases := []struct {
		evr1, evr2 rpm.EVR
		want       int
	}{
		{rpm.EVR{"0", "2.7.11", "4.el7"}, rpm.EVR{"0", "2.7.11", "4.el7"}, 0},
		{rpm.EVR{"1", "2.4.4", "17.el7"}, rpm.EVR{"", "2.4.4", "17.el7"}, 1},
		{rpm.EVR{"1", "2.0a1", "5.20110709git097ed8.el7"},
			rpm.EVR{"1", "2.0a1", "5.20110710git0c7ef8.el7"}, -1},
	}
	for _, c := range cases {
		got := rpm.LabelCompare(&c.evr1, &c.evr2)
		if got != c.want {
			t.Errorf("LabelCompare(%q, %q) == %d, expected %d", c.evr1, c.evr2, got, c.want)
		}
	}
}

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
