/* -*- mode:go; coding:utf-8; -*-
 * author: Eugene G. Zamriy <eugene@zamriy.info>
 * created: 23.11.2013 17:18
 * description: Mostly 1:1 bindings to the functions defined in rpmts.h.
 */


package rpm


// #cgo LDFLAGS: -lrpm
// #include <rpm/rpmlib.h>
// #include <rpm/rpmts.h>
import "C"


type RpmTs struct {
	c_ts C.rpmts
}


// RpmTsCreate (rpmtsCreate in RPM) creates an empty transaction set.
func RpmTsCreate() *RpmTs {
	return &RpmTs{c_ts: C.rpmtsCreate()}
}


// Free (rpmtsFree in RPM) destroys transaction set and closes the database.
func (ts *RpmTs) Free() {
	C.rpmtsFree(ts.c_ts)
}
