/* -*- mode:go; coding:utf-8; -*-
 * author: Eugene G. Zamriy <eugene@zamriy.info>
 * created: 23.11.2013 17:29
 * description: Mostly 1:1 bindings to the functions defined in rpmio.h.
 */


package rpm


// #cgo LDFLAGS: -lrpm -lrpmio
// #include <rpm/rpmlib.h>
import "C"

import (
	"path"
	"unsafe"
)


type FD_t struct {
	path, mode, file_name *C.char
	fd C.FD_t
}


// Fopen (Fopen in RPM) opens the file and associates a stream with it.
// It returns opened file descriptor on success, nil on error.
func Fopen(file_path string, mode string) (*FD_t, error) {
	c_path := C.CString(file_path)
	c_mode := C.CString(mode)
	_, file_name := path.Split(file_path)
	c_file_name := C.CString(file_name)

	fd, err := C.Fopen(c_path, c_mode)

	if fd == nil {
		C.free(unsafe.Pointer(c_path))
		C.free(unsafe.Pointer(c_mode))
		C.free(unsafe.Pointer(c_file_name))
		return nil, err
	}
	return &FD_t{path: c_path, mode: c_mode, file_name: c_file_name, fd: fd}, nil
}


// Fclose (Fclose in RPM) closes the stream pointed to by file descriptor FD_t.
// It returns 0 on success.
func (fd *FD_t) Fclose() (int, error) {
	defer C.free(unsafe.Pointer(fd.path))
	defer C.free(unsafe.Pointer(fd.mode))
	defer C.free(unsafe.Pointer(fd.file_name))

	c_code, err := C.Fclose(fd.fd)
	code := int(c_code)
	if code != 0 {
		return code, err
	}
	return code, nil
}
