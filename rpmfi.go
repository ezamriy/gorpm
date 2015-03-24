/* -*- mode:go; coding:utf-8; -*-
 * author: Christian Brunner <chb@muc.de>
 * created: 24.03.2015 20:45
 * description: Mostly 1:1 bindings to the functions defined in rpmfi.h
 */


package rpm


/*
#cgo LDFLAGS: -lrpm
#include <rpm/rpmfi.h>
*/
import "C"

type RpmFi struct {
	c_fi C.rpmfi
}

// FiNew (rpmfiNew in RPM) create and load a file info set.
func FiNew(hdr *Header) *RpmFi {
	c_fi := C.rpmfiNew(nil, hdr.c_header, C.rpmTag(RPMTAG_BASENAMES), C.RPMFI_KEEPHEADER)
	return &RpmFi{c_fi: c_fi}
}

// Free (rpmfiFree in RPM) destroy a file info set.
func (fi *RpmFi) Free() {
	C.rpmfiFree(fi.c_fi)
}

// Next (rpmfiNext in RPM) return next file iterator index. 
func (fi *RpmFi) Next() bool {
	r := C.rpmfiNext(fi.c_fi)
	if r < 0 {
		return false
	} else {
		return true
	}
}

// FN (rpmfiFN in RPM) return current file name from file info set. 
func (fi *RpmFi) FN() string {
	c_fn := C.rpmfiFN(fi.c_fi)
	return C.GoString(c_fn)
}
