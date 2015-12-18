/* -*- mode:go; coding:utf-8; -*-
 * author: Eugene Zamriy <eugene@zamriy.info>
 * created: 2015-12-18
 * description: Mostly 1:1 bindings to the functions defined in rpmmacro.h.
 */

package rpm

/*
#cgo LDFLAGS: -lrpm
#include <rpm/rpmmacro.h>

const char * rpmExpand_grapper(const char * macro) {
    return rpmExpand(macro, (const char*) NULL);
}
*/
import "C"

import (
	"unsafe"
)

// Expand (rpmExpand in RPM) returns macro expansion.
func Expand(macro string) string {
	cMacro := C.CString(macro)
	defer C.free(unsafe.Pointer(cMacro))
	cExpansion := C.rpmExpand_grapper(cMacro)
	defer C.free(unsafe.Pointer(cExpansion))
	return C.GoString(cExpansion)
}

// ExpandNumeric (rpmExpandNumeric) returns macro expansion as a numeric value.
func ExpandNumeric(macro string) int {
	cMacro := C.CString(macro)
	defer C.free(unsafe.Pointer(cMacro))
	return int(C.rpmExpandNumeric(cMacro))
}
