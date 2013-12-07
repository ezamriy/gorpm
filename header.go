/* -*- mode:go; coding:utf-8; -*-
 * author: Eugene G. Zamriy <eugene@zamriy.info>
 * created: 25.11.2013 21:43
 * description: Mostly 1:1 bindings to the functions defined in header.h
 */


package rpm


/*
#cgo LDFLAGS: -lrpm
#include <rpm/rpmlib.h>

const char * headerGetString_grapper(Header h, int tag) {
    return headerGetString(h, tag);
}

uint64_t headerGetNumber_grapper(Header h, int tag) {
    return headerGetNumber(h, tag);
}

int headerIsEntry_grapper(Header h, int tag) {
    return headerIsEntry(h, tag);
}
 */
import "C"


import (
	"errors"
)


type Header struct {
	c_header C.Header
}


// NewHeader (headerNew in RPM) creates new (empty) header instance.
func NewHeader() *Header {
	return &Header{c_header: C.headerNew()}
}


// Free (headerFree in RPM) dereferences a header instance.
func (hdr *Header) Free() {
	C.headerFree(hdr.c_header)
}


// IsSource (headerIsSource in RPM) checks if header is a source or binary
// package header.
// Returns true if source, false if binary.
func (hdr *Header) IsSource() bool {
	return int(C.headerIsSource(hdr.c_header)) == 1
}


// GetString (headerGetString in RPM) returns a simple string tag from header.
func (hdr *Header) GetString(tag RpmTag) (string, error) {
	c_str := C.headerGetString_grapper(hdr.c_header, C.int(tag))
	if c_str == nil {
		return "", errors.New("undefined tag value")
	}
	return C.GoString(c_str), nil
}


// GetNumber (headerGetNumber in RPM) returns a simple number tag from header.
func (hdr *Header) GetNumber(tag RpmTag) int64 {
	return int64(C.headerGetNumber_grapper(hdr.c_header, C.int(tag)))
}


// IsEntry (headerIsEntry in RPM) checks if tag exists in header.
// Returns true on success, false on failure.
func (hdr *Header) IsEntry(tag RpmTag) bool {
	return int(C.headerIsEntry_grapper(hdr.c_header, C.int(tag))) == 1
}
