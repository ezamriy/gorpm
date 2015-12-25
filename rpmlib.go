/* -*- mode:go; coding:utf-8; -*-
 * author: Eugene G. Zamriy <eugene@zamriy.info>
 * created: 23.11.2013 17:18
 * description: Mostly 1:1 bindings to the functions defined in rpmlib.h.
 */


// Go bindings for RPM library.
package rpm

// #cgo LDFLAGS: -lrpm
// #include <rpm/rpmlib.h>
import "C"

import (
	"errors"
	"unsafe"
)


// EVR contains RPM package Epoch, Version and Release
type EVR struct {
	Epoch, Version, Release string
}

// ReadConfigFiles (rpmReadConfigFiles in RPM) reads macro configuration
// file(s) for a target.
// It returns 0 on success, -1 on error
func ReadConfigFiles(file *string, target *string) (int, error) {
	var file_ptr, target_ptr *C.char

	if file != nil {
		file_ptr = C.CString(*file)
		defer C.free(unsafe.Pointer(file_ptr))
	}

	if target != nil {
		target_ptr = C.CString(*target)
		defer C.free(unsafe.Pointer(target_ptr))
	}

	c_code, err := C.rpmReadConfigFiles(file_ptr, target_ptr)
	code := int(c_code)
	if code != 0 {
		return code, err
	}
	return code, nil
}


// ReadPackageFile (rpmReadPackageFile in RPM) fills package header with data
// from file handle, verifying digests/signatures. Note that file name argument
// is omitted - we are taking it from file descriptor structure.
// Returns integer status of operation (see RPMRC_* constants).
func ReadPackageFile(ts *RpmTs, fd *FD_t, hdr *Header) (RpmRc, error) {
	c_status := C.rpmReadPackageFile(ts.c_ts, fd.fd, fd.file_name, &hdr.c_header)
	status := RpmRc(c_status)
	if status == RPMRC_NOTFOUND {
		return status, errors.New("RPM header is not found")
	} else if status == RPMRC_FAIL {
		return status, errors.New("RPM header reading failed")
	}
	return status, nil
}


// LabelCompare compares two packages EVR's.
// Returns 1 if first package is "newer", 0 if equal, -1 if second package is newer.
func LabelCompare(evr1, evr2 *EVR) int {
	rc := Vercmp(evr1.Epoch, evr2.Epoch)
	if rc == 0 {
		rc = Vercmp(evr1.Version, evr2.Version)
		if rc == 0 {
			rc = Vercmp(evr1.Release, evr2.Release)
		}
	}
	return rc
}

// Vercmp (rpmvercmp in RPM) does segmented string comparison for version or
// release strings.
// Returns 1 if a is "newer", 0 if equal, -1 if b is "newer".
func Vercmp(a string, b string) int {
	c_a := C.CString(a)
	defer C.free(unsafe.Pointer(c_a))

	c_b := C.CString(b)
	defer C.free(unsafe.Pointer(c_b))

	return int(C.rpmvercmp(c_a, c_b))
}
