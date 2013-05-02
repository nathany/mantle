package gl

// #include <OpenGL/gl3.h>
// #include <stdlib.h>
import "C"

import "unsafe"

/*
  allocCString converts a Go string to a C string, ready to send to OpenGL
*/
func allocCString(str string) (ptr *C.GLchar, length C.GLint) {
    length = C.GLint(len(str))
    ptr = (*C.GLchar)(C.CString(str))
    return ptr, length
}

/*
  freeCString must be called for any C string created with allocCString
*/
func freeCString(ptr *C.GLchar) {
    C.free(unsafe.Pointer(ptr))
}

// stringFiller is a function passed to fillGoString
type stringFiller func(ptr *C.GLchar, size C.GLsizei)

/*
  fillGoString will call a C function with a buffer of a specified size.
*/
func fillGoString(size int, fun stringFiller) string {
    ptr := C.malloc(C.size_t(size))
    defer C.free(ptr)
    fun((*C.GLchar)(ptr), C.GLsizei(size))
    return C.GoString((*C.char)(ptr))
}
