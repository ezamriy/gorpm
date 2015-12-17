/* -*- mode:go; coding:utf-8; -*-
 * author: Christian Brunner <chb@muc.de>
 * created: 25.11.2013 23:58
 * description: Mostly 1:1 bindings to the functions defined in rpmdb.h
 */


package rpm

/*
#cgo LDFLAGS: -lrpm
#include <rpm/rpmdb.h>
*/
import "C"
import "syscall"

type RpmDb struct {
	c_db C.rpmdb
}

type RpmdbMatchIterator struct {
	c_rpmdbMatchIterator C.rpmdbMatchIterator
}

// DbOpen (rpmdbOpen in RPM) open rpm database.
func DbOpen() *RpmDb {
	var c_db C.rpmdb
	C.rpmdbOpen(C.CString(""), &c_db, syscall.O_RDONLY, 0644)
	return &RpmDb{c_db: c_db}
}

// InitIterator (rpmdbInitIterator in RPM) return database iterator. 
func (db *RpmDb) InitIterator(tag RpmTag) *RpmdbMatchIterator {
	c_rpmdbMatchIterator := C.rpmdbInitIterator(db.c_db, C.rpmTag(tag), nil, 0)
	return &RpmdbMatchIterator{c_rpmdbMatchIterator: c_rpmdbMatchIterator}
}

// NextIterator(rpmdbNextIterator in RPM) return next package header from iteration. 
func (db *RpmDb) NextIterator(iterator *RpmdbMatchIterator) *Header {
	c_header := C.rpmdbNextIterator(iterator.c_rpmdbMatchIterator)
	if c_header == nil {
		return nil
	} else {
		return &Header{c_header: c_header}
	}
}

// FreeIterator (rpmdbFreeIterator in RPM) destroy rpm database iterator.
func (db *RpmDb) FreeIterator(iterator *RpmdbMatchIterator) {
	C.rpmdbFreeIterator(iterator.c_rpmdbMatchIterator)
}

// Close (rpmdbClose in RPM) close all database indices and free rpmdb.
func (db *RpmDb) Close() {
	C.rpmdbClose(db.c_db)
}

