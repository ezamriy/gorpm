/* -*- mode:go; coding:utf-8; -*-
 * author: Eugene Zamriy <eugene@zamriy.info>
 * created: 2015-12-19
 * description: Mostly 1:1 bindings to the functions defined in rpmspec.h.
 */

package rpm

/*
#cgo LDFLAGS: -lrpm
#include <rpm/rpmspec.h>
*/
import "C"

type RpmSpecFlag uint32

const (
	RPMSPEC_NONE    = RpmSpecFlag(C.RPMSPEC_NONE)
	RPMSPEC_ANYARCH = RpmSpecFlag(C.RPMSPEC_ANYARCH)
	RPMSPEC_FORCE   = RpmSpecFlag(C.RPMSPEC_FORCE)
	RPMSPEC_NOLANG  = RpmSpecFlag(C.RPMSPEC_NOLANG)
	RPMSPEC_NOUTF8  = RpmSpecFlag(C.RPMSPEC_NOUTF8)
)
