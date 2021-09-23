/* -*- mode:go; coding:utf-8; -*-
 * author: Eugene Zamriy <eugene@zamriy.info>
 * created: 2015-12-19
 * description: Mostly 1:1 bindings to the functions defined in rpmbuild.h.
 */

package rpm

/*
#cgo LDFLAGS: -lrpm -lrpmbuild
#include <rpm/rpmbuild.h>
#include <rpm/rpmspec.h>
*/
import "C"
import "unsafe"

type RpmBuildFlag uint32

const (
	RPMBUILD_NONE          = RpmBuildFlag(C.RPMBUILD_NONE)
	RPMBUILD_PREP          = RpmBuildFlag(C.RPMBUILD_PREP)                // execute %prep
	RPMBUILD_BUILD         = RpmBuildFlag(C.RPMBUILD_BUILD)               // execute %build
	RPMBUILD_INSTALL       = RpmBuildFlag(C.RPMBUILD_INSTALL)             // execute install
	RPMBUILD_CHECK         = RpmBuildFlag(C.RPMBUILD_CHECK)               // execute %check
	RPMBUILD_CLEAN         = RpmBuildFlag(C.RPMBUILD_CLEAN)               // execute %clean
	RPMBUILD_FILECHECK     = RpmBuildFlag(C.RPMBUILD_FILECHECK)           // check %files manifest
	RPMBUILD_PACKAGESOURCE = RpmBuildFlag(C.RPMBUILD_PACKAGESOURCE)       // create source package
	RPMBUILD_PACKAGEBINARY = RpmBuildFlag(C.RPMBUILD_PACKAGEBINARY)       // create binary package(s)
	RPMBUILD_RMSOURCE      = RpmBuildFlag(C.RPMBUILD_RMSOURCE)            // remove source(s) and patch(s)
	RPMBUILD_RMBUILD       = RpmBuildFlag(C.RPMBUILD_RMBUILD)             // remove build sub-tree
	RPMBUILD_STRINGBUF     = RpmBuildFlag(C.RPMBUILD_STRINGBUF)           // internal use only
	RPMBUILD_RMSPEC        = RpmBuildFlag(C.RPMBUILD_RMSPEC)              // remove spec file
	RPMBUILD_FILE_FILE     = RpmBuildFlag(C.RPMBUILD_FILE_FILE)           // rpmSpecPkgGetSection: %files -f
	RPMBUILD_FILE_LIST     = RpmBuildFlag(C.RPMBUILD_FILE_LIST)           // rpmSpecPkgGetSection: %files
	RPMBUILD_POLICY        = RpmBuildFlag(C.RPMBUILD_POLICY)              // rpmSpecPkgGetSection: %policy
	RPMBUILD_NOBUILD       = RpmBuildFlag(C.RPMBUILD_NOBUILD + (1 << 32)) // don't execute or package
)

type Spec struct {
	rpmSpec C.rpmSpec
}

// SpecParse (rpmSpecParse in RPM) parses spec file into Spec structure.
func SpecParse(specFile string, specFlags RpmSpecFlag) *Spec {
	cSpecFile := C.CString(specFile)
	defer C.free(unsafe.Pointer(cSpecFile))
	rpmSpec := C.rpmSpecParse(cSpecFile, C.rpmSpecFlags(specFlags), nil)
	return &Spec{rpmSpec: rpmSpec}
}

// Free (rpmSpecFree in RPM) destroys Spec structure.
func (spec *Spec) Free() {
	C.rpmSpecFree(spec.rpmSpec)
}

// SourceHeader (rpmSpecSourceHeader in RPM) returns the header of the src-RPM
// that would be built from the spec file.
func (spec *Spec) SourceHeader() *Header {
	cHeader := C.rpmSpecSourceHeader(spec.rpmSpec)
	return &Header{c_header: cHeader}
}
