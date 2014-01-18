/* -*- mode:go; coding:utf-8; -*-
 * author: Eugene G. Zamriy <eugene@zamriy.info>
 * created: 18.01.2014 23:28
 * description: gorpm library usage example. Returns yum info-like information
 *              for specified RPM file.
 */


package main

// #cgo LDFLAGS: -lrpm -lrpmio
// #include <rpm/rpmlib.h>
import "C"
import (
	"fmt"
	"github.com/ezamriy/gorpm"
	"os"
)


func print_info(hdr *rpm.Header) {
	name, _ := hdr.GetString(rpm.RPMTAG_NAME)
	arch, _ := hdr.GetString(rpm.RPMTAG_ARCH)
	version, _ := hdr.GetString(rpm.RPMTAG_VERSION)
	release, _ := hdr.GetString(rpm.RPMTAG_RELEASE)
	summary, _ := hdr.GetString(rpm.RPMTAG_SUMMARY)
	url, _ := hdr.GetString(rpm.RPMTAG_URL)
	license, _ := hdr.GetString(rpm.RPMTAG_LICENSE)
	description, _ := hdr.GetString(rpm.RPMTAG_DESCRIPTION)
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Arch: %s\n", arch)
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Release: %s\n", release)
	fmt.Printf("Size: %d\n", hdr.GetNumber(rpm.RPMTAG_SIZE))
	fmt.Printf("Summary: %s\n", summary)
	fmt.Printf("URL: %s\n", url)
	fmt.Printf("License: %s\n", license)
	fmt.Printf("Description: %s\n\n", description)
}


func format_flags(flags rpm.RpmFlag) string {
	s := ""
	if flags & rpm.RPMSENSE_LESS > 0 {
		s = s + "<"
	} else if flags & rpm.RPMSENSE_GREATER > 0 {
		s = s + ">"
	}
	if flags & rpm.RPMSENSE_EQUAL > 0 {
		s = s + "="
	}
	return s
}


func print_ds(hdr *rpm.Header, tag rpm.RpmTag) {
	ds := rpm.NewDependencySet(hdr, tag)
	defer ds.Free()

	if ds.Count() == 0 {
		return
	}
	switch tag {
	case rpm.RPMTAG_PROVIDENAME: fmt.Println("Provides:")
	case rpm.RPMTAG_REQUIRENAME: fmt.Println("Requires:")
	case rpm.RPMTAG_CONFLICTNAME: fmt.Println("Conflicts:")
	case rpm.RPMTAG_OBSOLETENAME: fmt.Println("Obsoletes:")
	default: fmt.Println("Unknown DS:")
	}
	for ds.Next() >= 0 {
		name, _ := ds.Name()
		flags := ds.Flags()
		evr, _ := ds.EVR()
		fmt.Printf("\t%+v\t%+v\t%+v\n", name, format_flags(flags), evr)
	}
}


func main() {
	file := os.Args[1]

	_, err := rpm.ReadConfigFiles(nil, nil)
	if err != nil {
		panic(fmt.Sprintf("cannot read RPM config files: %+v", err))
	}

	fd, err := rpm.Fopen(file, "r")
	if err != nil {
		panic(fmt.Sprintf("cannot read open file %+v: %+v", file, err))
	}
	defer fd.Fclose()

	ts := rpm.RpmTsCreate()
	defer ts.Free()

	header := rpm.NewHeader()
	defer header.Free()

	rc, err := rpm.ReadPackageFile(ts, fd, header)
	if err != nil {
		panic(fmt.Sprintf("cannot read RPM package file %+v: %+v", file, err))
	} else if rc == rpm.RPMRC_NOTTRUSTED {
		fmt.Println("signature verified, but key is not trusted")
	} else if rc == rpm.RPMRC_NOKEY {
		fmt.Println("public key is not available")
	}
	
	print_info(header)
	print_ds(header, rpm.RPMTAG_PROVIDENAME)
	print_ds(header, rpm.RPMTAG_REQUIRENAME)
	print_ds(header, rpm.RPMTAG_CONFLICTNAME)
	print_ds(header, rpm.RPMTAG_OBSOLETENAME)
}
