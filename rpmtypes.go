/* -*- mode:go; coding:utf-8; -*-
 * author: Eugene G. Zamriy <eugene@zamriy.info>
 * created: 26.11.2013 21:04
 * description: RPM data types (mostly 1:1 bindings to the rpmtypes.h)
 */


package rpm


// #cgo LDFLAGS: -lrpm
// #include <rpm/rpmtypes.h>
import "C"


type RpmRc int

const (
	RPMRC_OK         = RpmRc(C.RPMRC_OK)          // generic success code
	RPMRC_NOTFOUND   = RpmRc(C.RPMRC_NOTFOUND)    // generic not found code
	RPMRC_FAIL       = RpmRc(C.RPMRC_FAIL)        // generic failure code
	RPMRC_NOTTRUSTED = RpmRc(C.RPMRC_NOTTRUSTED)  // signature is OK, but key is not trusted
	RPMRC_NOKEY      = RpmRc(C.RPMRC_NOKEY)       // public key is unavailable
)
