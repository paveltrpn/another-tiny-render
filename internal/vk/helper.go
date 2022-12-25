package vk

// #include <stdint.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"bytes"
	"reflect"
	"strings"
	"unsafe"
)

type integerConstrait interface {
	int | int32 | uint | uint32
}

// https://github.com/sparkkoori/go-vulkan/blob/master/v1.1/vk/common.go
//
// What to do with returned value of type *C.char
// after using? Call C.free() manually???
//
// C.CString() - good replacement for this!!!
func toCString(s string) *C.char {
	if s == "" {
		return nil
	}
	n := len(s)
	p := C.malloc(C.ulong(n + 1))

	slice := (*[1 << 31]C.char)(p)[0 : n+1]

	for i := 0; i < n; i++ {
		slice[i] = C.char(s[i])
	}
	slice[n] = 0

	_s := (*C.char)(p)
	return _s
}

// https://github.com/sparkkoori/go-vulkan/blob/master/v1.1/vk/common.go
//
// C.GoString - good replacement for this!!!
func toGoString(p *C.char) string {
	if p == nil {
		return ""
	}
	slice := (*[1 << 31]C.char)(unsafe.Pointer(p))
	var buffer bytes.Buffer

	for i := 0; ; i++ {
		if slice[i] == 0 {
			break
		}
		buffer.WriteByte(byte(slice[i]))
	}
	return buffer.String()
}

// https://github.com/go-gl/gl/blob/master/v4.5-core/gl/conversions.go
//
// Function takes a null-terminated Go string and returns its GL-compatible address.
// This function reaches into Go string storage in an unsafe way so the caller
// must ensure the string is not garbage collected.
func toCStringManaged(str string) *C.char {
	if !strings.HasSuffix(str, "\x00") {
		panic("str argument missing null terminator: " + str)
	}
	header := (*reflect.StringHeader)(unsafe.Pointer(&str))
	return (*C.char)(unsafe.Pointer(header.Data))
}

// Allocate C memory enought for store one or more entities
// with type T, where T assumed as vulkan structs.
// Parameter str - pointer to struct, because it passed
// by value, return value rt is needed - result pointer
// to allocated space
func allocVkStructs[T any, V integerConstrait](str *T, count V) (rt *T) {
	// debug - fmt.Println(unsafe.Sizeof(*str))
	rt = (*T)(C.malloc(C.size_t(unsafe.Sizeof(*str)) * C.size_t(count)))
	return
}

// split C string (*char) to array of strings (**char)
func splitCString(original *C.char, split *C.char) **C.char {
	goResult := strings.Split(C.GoString(original), C.GoString(split))
	cArray := C.malloc(C.size_t(len(goResult)) * C.size_t(unsafe.Sizeof(uintptr(0))))

	a := (*[1<<30 - 1]*C.char)(cArray)

	for idx, substring := range goResult {
		a[idx] = C.CString(substring)
	}

	return (**C.char)(cArray)
}

func strSliceToCArrray(slc []string) **C.char {
	cArray := C.malloc(C.size_t(len(slc)) * C.size_t(unsafe.Sizeof(uintptr(0))))

	// convert the C array to a Go Array so we can index it
	a := (*[1<<30 - 1]*C.char)(cArray)

	for idx, substring := range slc {
		a[idx] = C.CString(substring)
	}

	return (**C.char)(cArray)
}

func cFree[T any](p *T) {
	C.free(unsafe.Pointer(p))
}

// Convert pointer with C type, reference to memory allocated for
// one or more structs of C type (malloc(sizeof(type) * count)) to
// golang slice with C type
func cArrayToSlice[T any](arr *T, length int) (slice []T) {
	// Old way
	//
	// *[1 << 30]C.YourType doesn't do anything itself,
	//it's a type. Specifically, it's a pointer to an array
	//of size 1 << 30, of C.YourType values. The size is arbitrary,
	//and only represents an upper bound that needs to be valid on the host system.
	//
	// What you're doing in the third expression is a
	// type conversion. This converts the unsafe.Pointer
	// to a *[1 << 30]C.YourType.
	//
	// Then, you're taking that converted array value, and
	// turning it into a slice with a full slice expression
	// (Array values don't need to be dereferenced for a slice
	// expression, so there is no need to prefix
	// the value with a *, even though it is a pointer).
	//
	// slice = (*[1 << 30]ctVkExtensionProperties)(unsafe.Pointer(arr))[:length:length]

	// Modern way, since go 1.17
	slice = unsafe.Slice(arr, length)
	return
}
