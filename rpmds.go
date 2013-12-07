/* -*- mode:go; coding:utf-8; -*-
 * author: Eugene G. Zamriy <eugene@zamriy.info>
 * created: 25.11.2013 23:58
 * description: Mostly 1:1 bindings to the functions defined in rpmds.h
 */


package rpm


/*
#cgo LDFLAGS: -lrpm
#include <rpm/rpmlib.h>

rpmds rpmdsNew_grapper(Header h, int tag, int flags) {
    return rpmdsNew(h, tag, flags);
}
 */
import "C"

import (
	"errors"
)


type RpmFlag uint32

// NOTE: hardcoded flags are missing in 4.8 version of RPM (RHEL 6)
const (
	RPMSENSE_ANY           = RpmFlag(C.RPMSENSE_ANY)
	RPMSENSE_LESS          = RpmFlag(C.RPMSENSE_LESS)
	RPMSENSE_GREATER       = RpmFlag(C.RPMSENSE_GREATER)
	RPMSENSE_EQUAL         = RpmFlag(C.RPMSENSE_EQUAL)
	RPMSENSE_POSTTRANS     = RpmFlag(1 << 5)              // %posttrans dependency
	RPMSENSE_PREREQ        = RpmFlag(C.RPMSENSE_PREREQ)   // legacy prereq dependency
	RPMSENSE_PRETRANS      = RpmFlag(1 << 7)              // pre-transaction dependency
	RPMSENSE_INTERP        = RpmFlag(C.RPMSENSE_INTERP)   // interpreter used by scriptlet
	RPMSENSE_SCRIPT_PRE    = RpmFlag(1 << 9)              // %pre dependency
	RPMSENSE_SCRIPT_POST   = RpmFlag(1 << 10)             // %post dependency
	RPMSENSE_SCRIPT_PREUN  = RpmFlag(1 << 11)             // %preun dependency
	RPMSENSE_SCRIPT_POSTUN = RpmFlag(1 << 12)             // %postun dependency
	RPMSENSE_SCRIPT_VERIFY = RpmFlag(1 << 13)             // %verify dependency
	RPMSENSE_FIND_REQUIRES = RpmFlag(1 << 14)             // find-requires generated dependency
	RPMSENSE_FIND_PROVIDES = RpmFlag(1 << 15)             // find-provides generated dependency
	RPMSENSE_TRIGGERIN     = RpmFlag(1 << 16)             // %triggerin dependency
	RPMSENSE_TRIGGERUN     = RpmFlag(1 << 17)             // %triggerun dependency
	RPMSENSE_TRIGGERPOSTUN = RpmFlag(1 << 18)             // %triggerpostun dependency
	RPMSENSE_MISSINGOK     = RpmFlag(1 << 19)             // suggests/enhances hint
	RPMSENSE_RPMLIB        = RpmFlag(1 << 24)             // rpmlib(feature) dependency
	RPMSENSE_TRIGGERPREIN  = RpmFlag(1 << 25)             // %triggerprein dependency
	RPMSENSE_KEYRING       = RpmFlag(1 << 26)
	RPMSENSE_CONFIG        = RpmFlag(1 << 28)
)


type DependencySet struct {
	c_rpmds C.rpmds
}


// NewDependencySet (rpmdsNew in RPM) creates and loads a dependency set.
func NewDependencySet(hdr *Header, tag RpmTag) *DependencySet {
	c_rpmds := C.rpmdsNew_grapper(hdr.c_header, C.int(tag), C.int(RPMSENSE_ANY))
	return &DependencySet{c_rpmds: c_rpmds}
}


// Free (rpmdsFree in RPM) destroys a dependency set.
func (ds *DependencySet) Free() {
	C.rpmdsFree(ds.c_rpmds)
}


// Count (rpmdsCount in RPM) returns dependency set count.
func (ds *DependencySet) Count() int {
	return int(C.rpmdsCount(ds.c_rpmds))
}


// Next (rpmdsNext in RPM) returns next dependency set iterator index or -1
// on termination.
func (ds *DependencySet) Next() int {
	return int(C.rpmdsNext(ds.c_rpmds))
}


// Name (rpmdsN in RPM) returns current dependency name.
func (ds *DependencySet) Name() (string, error) {
	c_name := C.rpmdsN(ds.c_rpmds)
	if c_name == nil {
		return "", errors.New("invalid data")
	}
	return C.GoString(c_name), nil
}


// EVR (rpmdsEVR in RPM) returns current dependency epoch-version-release.
func (ds *DependencySet) EVR() (string, error) {
	c_EVR := C.rpmdsEVR(ds.c_rpmds)
	if c_EVR == nil {
		return "", errors.New("invalid data")
	}
	return C.GoString(c_EVR), nil
}


// DNEVR (rpmdsDNEVR in RPM) returns current formatted dependency string.
func (ds *DependencySet) DNEVR() (string, error) {
	c_DNEVR := C.rpmdsDNEVR(ds.c_rpmds)
	if c_DNEVR == nil {
		return "", errors.New("invalid data")
	}
	return C.GoString(c_DNEVR), nil
}


// Flags (rpmdsFlags in RPM) returns current dependency flags.
func (ds *DependencySet) Flags() RpmFlag {
	return RpmFlag(C.rpmdsFlags(ds.c_rpmds))
}
