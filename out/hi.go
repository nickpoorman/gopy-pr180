/*
cgo stubs for package hi.
File is generated by gopy. Do not edit.
gopy build -vm=python3 -output=/out github.com/go-python/gopy/_examples/hi
*/

package main

/*
#cgo pkg-config: /usr/lib/pkgconfig/python3.pc
// #define Py_LIMITED_API // need full API for PyRun*
#include <Python.h>
typedef uint8_t bool;
// static inline is trick for avoiding need for extra .c file
// the following are used for build value -- switch on reflect.Kind
// or the types equivalent
static inline PyObject* gopy_build_bool(uint8_t val) {
	return Py_BuildValue("b", val);
}
static inline PyObject* gopy_build_int64(int64_t val) {
	return Py_BuildValue("k", val);
}
static inline PyObject* gopy_build_uint64(uint64_t val) {
	return Py_BuildValue("K", val);
}
static inline PyObject* gopy_build_float64(double val) {
	return Py_BuildValue("d", val);
}
static inline PyObject* gopy_build_string(const char* val) {
	return Py_BuildValue("s", val);
}
static inline void gopy_decref(PyObject* obj) { // macro
	Py_XDECREF(obj);
}
static inline void gopy_incref(PyObject* obj) { // macro
	Py_XINCREF(obj);
}
static inline int gopy_method_check(PyObject* obj) { // macro
	return PyMethod_Check(obj);
}
static inline void gopy_err_handle() {
	if(PyErr_Occurred() != NULL) {
		PyErr_Print();
	}
}

*/
import "C"
import (
	"reflect"
	"unsafe"

	"github.com/go-python/gopy/gopyh" // handler

	"github.com/go-python/gopy/_examples/hi"
	"github.com/go-python/gopy/_examples/structs"
)

// main doesn't do anything in lib / pkg mode, but is essential for exe mode
func main() {

}

// initialization functions -- can be called from python after library is loaded
// GoPyInitRunFile runs a separate python file -- call in GoPyInit if it
// steals the main thread e.g., for GUI event loop, as in GoGi startup.

//export GoPyInit
func GoPyInit() {

}

// type for the handle -- int64 for speed (can switch to string)
type GoHandle int64
type CGoHandle C.longlong

// boolGoToPy converts a Go bool to python-compatible C.char
func boolGoToPy(b bool) C.char {
	if b {
		return 1
	}
	return 0
}

// boolPyToGo converts a python-compatible C.Char to Go bool
func boolPyToGo(b C.char) bool {
	if b != 0 {
		return true
	}
	return false
}

func complex64GoToPy(c complex64) *C.PyObject {
	return C.PyComplex_FromDoubles(C.double(real(c)), C.double(imag(c)))
}

func complex64PyToGo(o *C.PyObject) complex64 {
	v := C.PyComplex_AsCComplex(o)
	return complex(float32(v.real), float32(v.imag))
}

func complex128GoToPy(c complex128) *C.PyObject {
	return C.PyComplex_FromDoubles(C.double(real(c)), C.double(imag(c)))
}

func complex128PyToGo(o *C.PyObject) complex128 {
	v := C.PyComplex_AsCComplex(o)
	return complex(float64(v.real), float64(v.imag))
}

// --- generated code for package: hi below: ---

// ---- External Types Outside of Targeted Packages ---

// Converters for pointer handles for type: *structs.S2
func ptrFromHandle_Ptr_structs_S2(h CGoHandle) *structs.S2 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "*structs.S2")
	if p == nil {
		return nil
	}
	return gopyh.Embed(p, reflect.TypeOf(structs.S2{})).(*structs.S2)
}
func handleFromPtr_Ptr_structs_S2(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("*structs.S2", p))
}

// Converters for non-pointer handles for type: structs.S2
func ptrFromHandle_structs_S2(h CGoHandle) *structs.S2 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "structs.S2")
	if p == nil {
		return nil
	}
	return gopyh.Embed(p, reflect.TypeOf(structs.S2{})).(*structs.S2)
}
func handleFromPtr_structs_S2(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("structs.S2", p))
}

// ---- Package: go ---

// ---- Types ---

// Converters for non-pointer handles for type: []bool
func ptrFromHandle_Slice_bool(h CGoHandle) *[]bool {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]bool")
	if p == nil {
		return nil
	}
	return p.(*[]bool)
}
func handleFromPtr_Slice_bool(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]bool", p))
}

// --- wrapping slice: []bool ---
//export Slice_bool_CTor
func Slice_bool_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_bool(&[]bool{}))
}

//export Slice_bool_len
func Slice_bool_len(handle CGoHandle) int {
	return len(*ptrFromHandle_Slice_bool(handle))
}

//export Slice_bool_elem
func Slice_bool_elem(handle CGoHandle, _idx int) C.char {
	s := *ptrFromHandle_Slice_bool(handle)
	return boolGoToPy(s[_idx])
}

//export Slice_bool_set
func Slice_bool_set(handle CGoHandle, _idx int, _vl C.char) {
	s := *ptrFromHandle_Slice_bool(handle)
	s[_idx] = boolPyToGo(_vl)
}

//export Slice_bool_append
func Slice_bool_append(handle CGoHandle, _vl C.char) {
	s := ptrFromHandle_Slice_bool(handle)
	*s = append(*s, boolPyToGo(_vl))
}

// Converters for non-pointer handles for type: []byte
func ptrFromHandle_Slice_byte(h CGoHandle) *[]byte {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]byte")
	if p == nil {
		return nil
	}
	return p.(*[]byte)
}
func handleFromPtr_Slice_byte(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]byte", p))
}

// --- wrapping slice: []byte ---
//export Slice_byte_CTor
func Slice_byte_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_byte(&[]byte{}))
}

//export Slice_byte_len
func Slice_byte_len(handle CGoHandle) int {
	return len(*ptrFromHandle_Slice_byte(handle))
}

//export Slice_byte_elem
func Slice_byte_elem(handle CGoHandle, _idx int) C.char {
	s := *ptrFromHandle_Slice_byte(handle)
	return C.char(s[_idx])
}

//export Slice_byte_set
func Slice_byte_set(handle CGoHandle, _idx int, _vl C.char) {
	s := *ptrFromHandle_Slice_byte(handle)
	s[_idx] = byte(_vl)
}

//export Slice_byte_append
func Slice_byte_append(handle CGoHandle, _vl C.char) {
	s := ptrFromHandle_Slice_byte(handle)
	*s = append(*s, byte(_vl))
}

// Converters for non-pointer handles for type: []float32
func ptrFromHandle_Slice_float32(h CGoHandle) *[]float32 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]float32")
	if p == nil {
		return nil
	}
	return p.(*[]float32)
}
func handleFromPtr_Slice_float32(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]float32", p))
}

// --- wrapping slice: []float32 ---
//export Slice_float32_CTor
func Slice_float32_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_float32(&[]float32{}))
}

//export Slice_float32_len
func Slice_float32_len(handle CGoHandle) int {
	return len(*ptrFromHandle_Slice_float32(handle))
}

//export Slice_float32_elem
func Slice_float32_elem(handle CGoHandle, _idx int) C.float {
	s := *ptrFromHandle_Slice_float32(handle)
	return C.float(s[_idx])
}

//export Slice_float32_set
func Slice_float32_set(handle CGoHandle, _idx int, _vl C.float) {
	s := *ptrFromHandle_Slice_float32(handle)
	s[_idx] = float32(_vl)
}

//export Slice_float32_append
func Slice_float32_append(handle CGoHandle, _vl C.float) {
	s := ptrFromHandle_Slice_float32(handle)
	*s = append(*s, float32(_vl))
}

// Converters for non-pointer handles for type: []float64
func ptrFromHandle_Slice_float64(h CGoHandle) *[]float64 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]float64")
	if p == nil {
		return nil
	}
	return p.(*[]float64)
}
func handleFromPtr_Slice_float64(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]float64", p))
}

// --- wrapping slice: []float64 ---
//export Slice_float64_CTor
func Slice_float64_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_float64(&[]float64{}))
}

//export Slice_float64_len
func Slice_float64_len(handle CGoHandle) int {
	return len(*ptrFromHandle_Slice_float64(handle))
}

//export Slice_float64_elem
func Slice_float64_elem(handle CGoHandle, _idx int) C.double {
	s := *ptrFromHandle_Slice_float64(handle)
	return C.double(s[_idx])
}

//export Slice_float64_set
func Slice_float64_set(handle CGoHandle, _idx int, _vl C.double) {
	s := *ptrFromHandle_Slice_float64(handle)
	s[_idx] = float64(_vl)
}

//export Slice_float64_append
func Slice_float64_append(handle CGoHandle, _vl C.double) {
	s := ptrFromHandle_Slice_float64(handle)
	*s = append(*s, float64(_vl))
}

// Converters for non-pointer handles for type: []int
func ptrFromHandle_Slice_int(h CGoHandle) *[]int {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]int")
	if p == nil {
		return nil
	}
	return p.(*[]int)
}
func handleFromPtr_Slice_int(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]int", p))
}

// --- wrapping slice: []int ---
//export Slice_int_CTor
func Slice_int_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_int(&[]int{}))
}

//export Slice_int_len
func Slice_int_len(handle CGoHandle) int {
	return len(*ptrFromHandle_Slice_int(handle))
}

//export Slice_int_elem
func Slice_int_elem(handle CGoHandle, _idx int) C.longlong {
	s := *ptrFromHandle_Slice_int(handle)
	return C.longlong(s[_idx])
}

//export Slice_int_set
func Slice_int_set(handle CGoHandle, _idx int, _vl C.longlong) {
	s := *ptrFromHandle_Slice_int(handle)
	s[_idx] = int(_vl)
}

//export Slice_int_append
func Slice_int_append(handle CGoHandle, _vl C.longlong) {
	s := ptrFromHandle_Slice_int(handle)
	*s = append(*s, int(_vl))
}

// Converters for non-pointer handles for type: []int16
func ptrFromHandle_Slice_int16(h CGoHandle) *[]int16 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]int16")
	if p == nil {
		return nil
	}
	return p.(*[]int16)
}
func handleFromPtr_Slice_int16(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]int16", p))
}

// --- wrapping slice: []int16 ---
//export Slice_int16_CTor
func Slice_int16_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_int16(&[]int16{}))
}

//export Slice_int16_len
func Slice_int16_len(handle CGoHandle) int {
	return len(*ptrFromHandle_Slice_int16(handle))
}

//export Slice_int16_elem
func Slice_int16_elem(handle CGoHandle, _idx int) C.short {
	s := *ptrFromHandle_Slice_int16(handle)
	return C.short(s[_idx])
}

//export Slice_int16_set
func Slice_int16_set(handle CGoHandle, _idx int, _vl C.short) {
	s := *ptrFromHandle_Slice_int16(handle)
	s[_idx] = int16(_vl)
}

//export Slice_int16_append
func Slice_int16_append(handle CGoHandle, _vl C.short) {
	s := ptrFromHandle_Slice_int16(handle)
	*s = append(*s, int16(_vl))
}

// Converters for non-pointer handles for type: []int32
func ptrFromHandle_Slice_int32(h CGoHandle) *[]int32 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]int32")
	if p == nil {
		return nil
	}
	return p.(*[]int32)
}
func handleFromPtr_Slice_int32(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]int32", p))
}

// --- wrapping slice: []int32 ---
//export Slice_int32_CTor
func Slice_int32_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_int32(&[]int32{}))
}

//export Slice_int32_len
func Slice_int32_len(handle CGoHandle) int {
	return len(*ptrFromHandle_Slice_int32(handle))
}

//export Slice_int32_elem
func Slice_int32_elem(handle CGoHandle, _idx int) C.long {
	s := *ptrFromHandle_Slice_int32(handle)
	return C.long(s[_idx])
}

//export Slice_int32_set
func Slice_int32_set(handle CGoHandle, _idx int, _vl C.long) {
	s := *ptrFromHandle_Slice_int32(handle)
	s[_idx] = int32(_vl)
}

//export Slice_int32_append
func Slice_int32_append(handle CGoHandle, _vl C.long) {
	s := ptrFromHandle_Slice_int32(handle)
	*s = append(*s, int32(_vl))
}

// Converters for non-pointer handles for type: []int64
func ptrFromHandle_Slice_int64(h CGoHandle) *[]int64 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]int64")
	if p == nil {
		return nil
	}
	return p.(*[]int64)
}
func handleFromPtr_Slice_int64(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]int64", p))
}

// --- wrapping slice: []int64 ---
//export Slice_int64_CTor
func Slice_int64_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_int64(&[]int64{}))
}

//export Slice_int64_len
func Slice_int64_len(handle CGoHandle) int {
	return len(*ptrFromHandle_Slice_int64(handle))
}

//export Slice_int64_elem
func Slice_int64_elem(handle CGoHandle, _idx int) C.longlong {
	s := *ptrFromHandle_Slice_int64(handle)
	return C.longlong(s[_idx])
}

//export Slice_int64_set
func Slice_int64_set(handle CGoHandle, _idx int, _vl C.longlong) {
	s := *ptrFromHandle_Slice_int64(handle)
	s[_idx] = int64(_vl)
}

//export Slice_int64_append
func Slice_int64_append(handle CGoHandle, _vl C.longlong) {
	s := ptrFromHandle_Slice_int64(handle)
	*s = append(*s, int64(_vl))
}

// Converters for non-pointer handles for type: []int8
func ptrFromHandle_Slice_int8(h CGoHandle) *[]int8 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]int8")
	if p == nil {
		return nil
	}
	return p.(*[]int8)
}
func handleFromPtr_Slice_int8(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]int8", p))
}

// --- wrapping slice: []int8 ---
//export Slice_int8_CTor
func Slice_int8_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_int8(&[]int8{}))
}

//export Slice_int8_len
func Slice_int8_len(handle CGoHandle) int {
	return len(*ptrFromHandle_Slice_int8(handle))
}

//export Slice_int8_elem
func Slice_int8_elem(handle CGoHandle, _idx int) C.char {
	s := *ptrFromHandle_Slice_int8(handle)
	return C.char(s[_idx])
}

//export Slice_int8_set
func Slice_int8_set(handle CGoHandle, _idx int, _vl C.char) {
	s := *ptrFromHandle_Slice_int8(handle)
	s[_idx] = int8(_vl)
}

//export Slice_int8_append
func Slice_int8_append(handle CGoHandle, _vl C.char) {
	s := ptrFromHandle_Slice_int8(handle)
	*s = append(*s, int8(_vl))
}

// Converters for non-pointer handles for type: []rune
func ptrFromHandle_Slice_rune(h CGoHandle) *[]rune {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]rune")
	if p == nil {
		return nil
	}
	return p.(*[]rune)
}
func handleFromPtr_Slice_rune(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]rune", p))
}

// --- wrapping slice: []rune ---
//export Slice_rune_CTor
func Slice_rune_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_rune(&[]rune{}))
}

//export Slice_rune_len
func Slice_rune_len(handle CGoHandle) int {
	return len(*ptrFromHandle_Slice_rune(handle))
}

//export Slice_rune_elem
func Slice_rune_elem(handle CGoHandle, _idx int) C.long {
	s := *ptrFromHandle_Slice_rune(handle)
	return C.long(s[_idx])
}

//export Slice_rune_set
func Slice_rune_set(handle CGoHandle, _idx int, _vl C.long) {
	s := *ptrFromHandle_Slice_rune(handle)
	s[_idx] = rune(_vl)
}

//export Slice_rune_append
func Slice_rune_append(handle CGoHandle, _vl C.long) {
	s := ptrFromHandle_Slice_rune(handle)
	*s = append(*s, rune(_vl))
}

// Converters for non-pointer handles for type: []string
func ptrFromHandle_Slice_string(h CGoHandle) *[]string {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]string")
	if p == nil {
		return nil
	}
	return p.(*[]string)
}
func handleFromPtr_Slice_string(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]string", p))
}

// --- wrapping slice: []string ---
//export Slice_string_CTor
func Slice_string_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_string(&[]string{}))
}

//export Slice_string_len
func Slice_string_len(handle CGoHandle) int {
	return len(*ptrFromHandle_Slice_string(handle))
}

//export Slice_string_elem
func Slice_string_elem(handle CGoHandle, _idx int) *C.char {
	s := *ptrFromHandle_Slice_string(handle)
	return C.CString(s[_idx])
}

//export Slice_string_set
func Slice_string_set(handle CGoHandle, _idx int, _vl *C.char) {
	s := *ptrFromHandle_Slice_string(handle)
	s[_idx] = C.GoString(_vl)
}

//export Slice_string_append
func Slice_string_append(handle CGoHandle, _vl *C.char) {
	s := ptrFromHandle_Slice_string(handle)
	*s = append(*s, C.GoString(_vl))
}

// Converters for non-pointer handles for type: []uint
func ptrFromHandle_Slice_uint(h CGoHandle) *[]uint {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]uint")
	if p == nil {
		return nil
	}
	return p.(*[]uint)
}
func handleFromPtr_Slice_uint(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]uint", p))
}

// --- wrapping slice: []uint ---
//export Slice_uint_CTor
func Slice_uint_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_uint(&[]uint{}))
}

//export Slice_uint_len
func Slice_uint_len(handle CGoHandle) int {
	return len(*ptrFromHandle_Slice_uint(handle))
}

//export Slice_uint_elem
func Slice_uint_elem(handle CGoHandle, _idx int) C.ulonglong {
	s := *ptrFromHandle_Slice_uint(handle)
	return C.ulonglong(s[_idx])
}

//export Slice_uint_set
func Slice_uint_set(handle CGoHandle, _idx int, _vl C.ulonglong) {
	s := *ptrFromHandle_Slice_uint(handle)
	s[_idx] = uint(_vl)
}

//export Slice_uint_append
func Slice_uint_append(handle CGoHandle, _vl C.ulonglong) {
	s := ptrFromHandle_Slice_uint(handle)
	*s = append(*s, uint(_vl))
}

// Converters for non-pointer handles for type: []uint16
func ptrFromHandle_Slice_uint16(h CGoHandle) *[]uint16 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]uint16")
	if p == nil {
		return nil
	}
	return p.(*[]uint16)
}
func handleFromPtr_Slice_uint16(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]uint16", p))
}

// --- wrapping slice: []uint16 ---
//export Slice_uint16_CTor
func Slice_uint16_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_uint16(&[]uint16{}))
}

//export Slice_uint16_len
func Slice_uint16_len(handle CGoHandle) int {
	return len(*ptrFromHandle_Slice_uint16(handle))
}

//export Slice_uint16_elem
func Slice_uint16_elem(handle CGoHandle, _idx int) C.ushort {
	s := *ptrFromHandle_Slice_uint16(handle)
	return C.ushort(s[_idx])
}

//export Slice_uint16_set
func Slice_uint16_set(handle CGoHandle, _idx int, _vl C.ushort) {
	s := *ptrFromHandle_Slice_uint16(handle)
	s[_idx] = uint16(_vl)
}

//export Slice_uint16_append
func Slice_uint16_append(handle CGoHandle, _vl C.ushort) {
	s := ptrFromHandle_Slice_uint16(handle)
	*s = append(*s, uint16(_vl))
}

// Converters for non-pointer handles for type: []uint32
func ptrFromHandle_Slice_uint32(h CGoHandle) *[]uint32 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]uint32")
	if p == nil {
		return nil
	}
	return p.(*[]uint32)
}
func handleFromPtr_Slice_uint32(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]uint32", p))
}

// --- wrapping slice: []uint32 ---
//export Slice_uint32_CTor
func Slice_uint32_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_uint32(&[]uint32{}))
}

//export Slice_uint32_len
func Slice_uint32_len(handle CGoHandle) int {
	return len(*ptrFromHandle_Slice_uint32(handle))
}

//export Slice_uint32_elem
func Slice_uint32_elem(handle CGoHandle, _idx int) C.ulong {
	s := *ptrFromHandle_Slice_uint32(handle)
	return C.ulong(s[_idx])
}

//export Slice_uint32_set
func Slice_uint32_set(handle CGoHandle, _idx int, _vl C.ulong) {
	s := *ptrFromHandle_Slice_uint32(handle)
	s[_idx] = uint32(_vl)
}

//export Slice_uint32_append
func Slice_uint32_append(handle CGoHandle, _vl C.ulong) {
	s := ptrFromHandle_Slice_uint32(handle)
	*s = append(*s, uint32(_vl))
}

// Converters for non-pointer handles for type: []uint64
func ptrFromHandle_Slice_uint64(h CGoHandle) *[]uint64 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]uint64")
	if p == nil {
		return nil
	}
	return p.(*[]uint64)
}
func handleFromPtr_Slice_uint64(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]uint64", p))
}

// --- wrapping slice: []uint64 ---
//export Slice_uint64_CTor
func Slice_uint64_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_uint64(&[]uint64{}))
}

//export Slice_uint64_len
func Slice_uint64_len(handle CGoHandle) int {
	return len(*ptrFromHandle_Slice_uint64(handle))
}

//export Slice_uint64_elem
func Slice_uint64_elem(handle CGoHandle, _idx int) C.ulonglong {
	s := *ptrFromHandle_Slice_uint64(handle)
	return C.ulonglong(s[_idx])
}

//export Slice_uint64_set
func Slice_uint64_set(handle CGoHandle, _idx int, _vl C.ulonglong) {
	s := *ptrFromHandle_Slice_uint64(handle)
	s[_idx] = uint64(_vl)
}

//export Slice_uint64_append
func Slice_uint64_append(handle CGoHandle, _vl C.ulonglong) {
	s := ptrFromHandle_Slice_uint64(handle)
	*s = append(*s, uint64(_vl))
}

// Converters for non-pointer handles for type: []uint8
func ptrFromHandle_Slice_uint8(h CGoHandle) *[]uint8 {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[]uint8")
	if p == nil {
		return nil
	}
	return p.(*[]uint8)
}
func handleFromPtr_Slice_uint8(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[]uint8", p))
}

// --- wrapping slice: []uint8 ---
//export Slice_uint8_CTor
func Slice_uint8_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Slice_uint8(&[]uint8{}))
}

//export Slice_uint8_len
func Slice_uint8_len(handle CGoHandle) int {
	return len(*ptrFromHandle_Slice_uint8(handle))
}

//export Slice_uint8_elem
func Slice_uint8_elem(handle CGoHandle, _idx int) C.uchar {
	s := *ptrFromHandle_Slice_uint8(handle)
	return C.uchar(s[_idx])
}

//export Slice_uint8_set
func Slice_uint8_set(handle CGoHandle, _idx int, _vl C.uchar) {
	s := *ptrFromHandle_Slice_uint8(handle)
	s[_idx] = uint8(_vl)
}

//export Slice_uint8_append
func Slice_uint8_append(handle CGoHandle, _vl C.uchar) {
	s := ptrFromHandle_Slice_uint8(handle)
	*s = append(*s, uint8(_vl))
}

// ---- Package: hi ---

// ---- Types ---

// Converters for pointer handles for type: *hi.Couple
func ptrFromHandle_Ptr_hi_Couple(h CGoHandle) *hi.Couple {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "*hi.Couple")
	if p == nil {
		return nil
	}
	return gopyh.Embed(p, reflect.TypeOf(hi.Couple{})).(*hi.Couple)
}
func handleFromPtr_Ptr_hi_Couple(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("*hi.Couple", p))
}

// Converters for pointer handles for type: *hi.Person
func ptrFromHandle_Ptr_hi_Person(h CGoHandle) *hi.Person {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "*hi.Person")
	if p == nil {
		return nil
	}
	return gopyh.Embed(p, reflect.TypeOf(hi.Person{})).(*hi.Person)
}
func handleFromPtr_Ptr_hi_Person(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("*hi.Person", p))
}

// Converters for non-pointer handles for type: [2]int
func ptrFromHandle_Array_2_int(h CGoHandle) *[2]int {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "[2]int")
	if p == nil {
		return nil
	}
	return p.(*[2]int)
}
func handleFromPtr_Array_2_int(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("[2]int", p))
}

// --- wrapping slice: [2]int ---
//export Array_2_int_CTor
func Array_2_int_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_Array_2_int(&[2]int{}))
}

//export Array_2_int_len
func Array_2_int_len(handle CGoHandle) int {
	return len(*ptrFromHandle_Array_2_int(handle))
}

//export Array_2_int_elem
func Array_2_int_elem(handle CGoHandle, _idx int) C.longlong {
	s := *ptrFromHandle_Array_2_int(handle)
	return C.longlong(s[_idx])
}

//export Array_2_int_set
func Array_2_int_set(handle CGoHandle, _idx int, _vl C.longlong) {
	s := *ptrFromHandle_Array_2_int(handle)
	s[_idx] = int(_vl)
}

// Converters for non-pointer handles for type: hi.Couple
func ptrFromHandle_hi_Couple(h CGoHandle) *hi.Couple {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "hi.Couple")
	if p == nil {
		return nil
	}
	return gopyh.Embed(p, reflect.TypeOf(hi.Couple{})).(*hi.Couple)
}
func handleFromPtr_hi_Couple(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("hi.Couple", p))
}

// Converters for non-pointer handles for type: hi.Floats
func ptrFromHandle_hi_Floats(h CGoHandle) *hi.Floats {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "hi.Floats")
	if p == nil {
		return nil
	}
	return p.(*hi.Floats)
}
func handleFromPtr_hi_Floats(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("hi.Floats", p))
}

// Converters for pointer handles for type: hi.PersIface
func ptrFromHandle_hi_PersIface(h CGoHandle) hi.PersIface {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "hi.PersIface")
	if p == nil {
		return nil
	}
	return p.(hi.PersIface)
}
func handleFromPtr_hi_PersIface(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("hi.PersIface", p))
}

// Converters for non-pointer handles for type: hi.Person
func ptrFromHandle_hi_Person(h CGoHandle) *hi.Person {
	p := gopyh.VarFromHandle((gopyh.CGoHandle)(h), "hi.Person")
	if p == nil {
		return nil
	}
	return gopyh.Embed(p, reflect.TypeOf(hi.Person{})).(*hi.Person)
}
func handleFromPtr_hi_Person(p interface{}) CGoHandle {
	return CGoHandle(gopyh.Register("hi.Person", p))
}

// ---- Global Variables: can only use functions to access ---
//export hi_Anon
func hi_Anon() CGoHandle {
	return handleFromPtr_hi_Person(&hi.Anon)
}

//export hi_Set_Anon
func hi_Set_Anon(val CGoHandle) {
	hi.Anon = *ptrFromHandle_hi_Person(val)
}

//export hi_Debug
func hi_Debug() C.char {
	return boolGoToPy(hi.Debug)
}

//export hi_Set_Debug
func hi_Set_Debug(val C.char) {
	hi.Debug = boolPyToGo(val)
}

//export hi_IntArray
func hi_IntArray() CGoHandle {
	return handleFromPtr_Array_2_int(&hi.IntArray)
}

//export hi_IntSlice
func hi_IntSlice() CGoHandle {
	return handleFromPtr_Slice_int(&hi.IntSlice)
}

//export hi_Set_IntSlice
func hi_Set_IntSlice(val CGoHandle) {
	hi.IntSlice = *ptrFromHandle_Slice_int(val)
}

// ---- Interfaces ---

//export hi_PersIface_GetAge
func hi_PersIface_GetAge(_handle CGoHandle) C.longlong {
	vifc, __err := gopyh.VarFromHandleTry((gopyh.CGoHandle)(_handle), "hi.PersIface")
	if __err != nil {
		return C.longlong(0)
	}
	return C.longlong(vifc.(hi.PersIface).GetAge())

}

//export hi_PersIface_GetName
func hi_PersIface_GetName(_handle CGoHandle) *C.char {
	vifc, __err := gopyh.VarFromHandleTry((gopyh.CGoHandle)(_handle), "hi.PersIface")
	if __err != nil {
		return C.CString("")
	}
	return C.CString(vifc.(hi.PersIface).GetName())

}

//export hi_PersIface_Greet
func hi_PersIface_Greet(_handle CGoHandle) *C.char {
	vifc, __err := gopyh.VarFromHandleTry((gopyh.CGoHandle)(_handle), "hi.PersIface")
	if __err != nil {
		return C.CString("")
	}
	return C.CString(vifc.(hi.PersIface).Greet())

}

//export hi_PersIface_SetAge
func hi_PersIface_SetAge(_handle CGoHandle, age C.longlong, goRun C.char) {
	vifc, __err := gopyh.VarFromHandleTry((gopyh.CGoHandle)(_handle), "hi.PersIface")
	if __err != nil {
		return
	}
	if boolPyToGo(goRun) {
		go vifc.(hi.PersIface).SetAge(int(age))
	} else {
		vifc.(hi.PersIface).SetAge(int(age))
	}
}

//export hi_PersIface_SetName
func hi_PersIface_SetName(_handle CGoHandle, n *C.char, goRun C.char) {
	vifc, __err := gopyh.VarFromHandleTry((gopyh.CGoHandle)(_handle), "hi.PersIface")
	if __err != nil {
		return
	}
	if boolPyToGo(goRun) {
		go vifc.(hi.PersIface).SetName(C.GoString(n))
	} else {
		vifc.(hi.PersIface).SetName(C.GoString(n))
	}
}

// ---- Structs ---

// --- wrapping struct: hi.Couple ---
//export hi_Couple_CTor
func hi_Couple_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_hi_Couple(&hi.Couple{}))
}

//export hi_Couple_P1_Get
func hi_Couple_P1_Get(handle CGoHandle) CGoHandle {
	op := ptrFromHandle_hi_Couple(handle)
	return handleFromPtr_hi_Person(&op.P1)
}

//export hi_Couple_P1_Set
func hi_Couple_P1_Set(handle CGoHandle, val CGoHandle) {
	op := ptrFromHandle_hi_Couple(handle)
	op.P1 = *ptrFromHandle_hi_Person(val)
}

//export hi_Couple_P2_Get
func hi_Couple_P2_Get(handle CGoHandle) CGoHandle {
	op := ptrFromHandle_hi_Couple(handle)
	return handleFromPtr_hi_Person(&op.P2)
}

//export hi_Couple_P2_Set
func hi_Couple_P2_Set(handle CGoHandle, val CGoHandle) {
	op := ptrFromHandle_hi_Couple(handle)
	op.P2 = *ptrFromHandle_hi_Person(val)
}

//export hi_Couple_String
func hi_Couple_String(_handle CGoHandle) *C.char {
	vifc, __err := gopyh.VarFromHandleTry((gopyh.CGoHandle)(_handle), "*hi.Couple")
	if __err != nil {
		return C.CString("")
	}
	return C.CString(gopyh.Embed(vifc, reflect.TypeOf(hi.Couple{})).(*hi.Couple).String())

}

// --- wrapping struct: hi.Person ---
//export hi_Person_CTor
func hi_Person_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_hi_Person(&hi.Person{}))
}

//export hi_Person_Name_Get
func hi_Person_Name_Get(handle CGoHandle) *C.char {
	op := ptrFromHandle_hi_Person(handle)
	return C.CString(op.Name)
}

//export hi_Person_Name_Set
func hi_Person_Name_Set(handle CGoHandle, val *C.char) {
	op := ptrFromHandle_hi_Person(handle)
	op.Name = C.GoString(val)
}

//export hi_Person_Age_Get
func hi_Person_Age_Get(handle CGoHandle) C.longlong {
	op := ptrFromHandle_hi_Person(handle)
	return C.longlong(op.Age)
}

//export hi_Person_Age_Set
func hi_Person_Age_Set(handle CGoHandle, val C.longlong) {
	op := ptrFromHandle_hi_Person(handle)
	op.Age = int(val)
}

//export hi_Person_String
func hi_Person_String(_handle CGoHandle) *C.char {
	vifc, __err := gopyh.VarFromHandleTry((gopyh.CGoHandle)(_handle), "*hi.Person")
	if __err != nil {
		return C.CString("")
	}
	return C.CString(gopyh.Embed(vifc, reflect.TypeOf(hi.Person{})).(*hi.Person).String())

}

//export hi_Person_Greet
func hi_Person_Greet(_handle CGoHandle) *C.char {
	vifc, __err := gopyh.VarFromHandleTry((gopyh.CGoHandle)(_handle), "*hi.Person")
	if __err != nil {
		return C.CString("")
	}
	return C.CString(gopyh.Embed(vifc, reflect.TypeOf(hi.Person{})).(*hi.Person).Greet())

}

//export hi_Person_Work
func hi_Person_Work(_handle CGoHandle, h C.longlong) *C.char {
	vifc, __err := gopyh.VarFromHandleTry((gopyh.CGoHandle)(_handle), "*hi.Person")
	if __err != nil {
		return C.CString("")
	}
	__err = gopyh.Embed(vifc, reflect.TypeOf(hi.Person{})).(*hi.Person).Work(int(h))

	if __err != nil {
		estr := C.CString(__err.Error())
		C.PyErr_SetString(C.PyExc_RuntimeError, estr)
		return estr
	}
	return C.CString("")
}

//export hi_Person_Salary
func hi_Person_Salary(_handle CGoHandle, h C.longlong) C.longlong {
	vifc, __err := gopyh.VarFromHandleTry((gopyh.CGoHandle)(_handle), "*hi.Person")
	if __err != nil {
		return C.longlong(0)
	}
	cret, __err := gopyh.Embed(vifc, reflect.TypeOf(hi.Person{})).(*hi.Person).Salary(int(h))

	if __err != nil {
		estr := C.CString(__err.Error())
		C.PyErr_SetString(C.PyExc_RuntimeError, estr)
		C.free(unsafe.Pointer(estr))
	}
	return C.longlong(cret)
}

//export hi_Person_GetName
func hi_Person_GetName(_handle CGoHandle) *C.char {
	vifc, __err := gopyh.VarFromHandleTry((gopyh.CGoHandle)(_handle), "*hi.Person")
	if __err != nil {
		return C.CString("")
	}
	return C.CString(gopyh.Embed(vifc, reflect.TypeOf(hi.Person{})).(*hi.Person).GetName())

}

//export hi_Person_GetAge
func hi_Person_GetAge(_handle CGoHandle) C.longlong {
	vifc, __err := gopyh.VarFromHandleTry((gopyh.CGoHandle)(_handle), "*hi.Person")
	if __err != nil {
		return C.longlong(0)
	}
	return C.longlong(gopyh.Embed(vifc, reflect.TypeOf(hi.Person{})).(*hi.Person).GetAge())

}

//export hi_Person_SetName
func hi_Person_SetName(_handle CGoHandle, n *C.char, goRun C.char) {
	vifc, __err := gopyh.VarFromHandleTry((gopyh.CGoHandle)(_handle), "*hi.Person")
	if __err != nil {
		return
	}
	if boolPyToGo(goRun) {
		go gopyh.Embed(vifc, reflect.TypeOf(hi.Person{})).(*hi.Person).SetName(C.GoString(n))
	} else {
		gopyh.Embed(vifc, reflect.TypeOf(hi.Person{})).(*hi.Person).SetName(C.GoString(n))
	}
}

//export hi_Person_SetAge
func hi_Person_SetAge(_handle CGoHandle, age C.longlong, goRun C.char) {
	vifc, __err := gopyh.VarFromHandleTry((gopyh.CGoHandle)(_handle), "*hi.Person")
	if __err != nil {
		return
	}
	if boolPyToGo(goRun) {
		go gopyh.Embed(vifc, reflect.TypeOf(hi.Person{})).(*hi.Person).SetAge(int(age))
	} else {
		gopyh.Embed(vifc, reflect.TypeOf(hi.Person{})).(*hi.Person).SetAge(int(age))
	}
}

//export hi_Person_SetFmS2
func hi_Person_SetFmS2(_handle CGoHandle, s2 CGoHandle, goRun C.char) {
	vifc, __err := gopyh.VarFromHandleTry((gopyh.CGoHandle)(_handle), "*hi.Person")
	if __err != nil {
		return
	}
	if boolPyToGo(goRun) {
		go gopyh.Embed(vifc, reflect.TypeOf(hi.Person{})).(*hi.Person).SetFmS2(*ptrFromHandle_structs_S2(s2))
	} else {
		gopyh.Embed(vifc, reflect.TypeOf(hi.Person{})).(*hi.Person).SetFmS2(*ptrFromHandle_structs_S2(s2))
	}
}

//export hi_Person_SetFmS2Ptr
func hi_Person_SetFmS2Ptr(_handle CGoHandle, s2 CGoHandle, goRun C.char) {
	vifc, __err := gopyh.VarFromHandleTry((gopyh.CGoHandle)(_handle), "*hi.Person")
	if __err != nil {
		return
	}
	if boolPyToGo(goRun) {
		go gopyh.Embed(vifc, reflect.TypeOf(hi.Person{})).(*hi.Person).SetFmS2Ptr(ptrFromHandle_Ptr_structs_S2(s2))
	} else {
		gopyh.Embed(vifc, reflect.TypeOf(hi.Person{})).(*hi.Person).SetFmS2Ptr(ptrFromHandle_Ptr_structs_S2(s2))
	}
}

//export hi_Person_ReturnS2Ptr
func hi_Person_ReturnS2Ptr(_handle CGoHandle) CGoHandle {
	vifc, __err := gopyh.VarFromHandleTry((gopyh.CGoHandle)(_handle), "*hi.Person")
	if __err != nil {
		return handleFromPtr_Ptr_structs_S2(nil)
	}
	return handleFromPtr_Ptr_structs_S2(gopyh.Embed(vifc, reflect.TypeOf(hi.Person{})).(*hi.Person).ReturnS2Ptr())

}

// ---- Slices ---

// --- wrapping slice: hi.Floats ---
//export hi_Floats_CTor
func hi_Floats_CTor() CGoHandle {
	return CGoHandle(handleFromPtr_hi_Floats(&hi.Floats{}))
}

//export hi_Floats_len
func hi_Floats_len(handle CGoHandle) int {
	return len(*ptrFromHandle_hi_Floats(handle))
}

//export hi_Floats_elem
func hi_Floats_elem(handle CGoHandle, _idx int) C.float {
	s := *ptrFromHandle_hi_Floats(handle)
	return C.float(float32(s[_idx]))
}

//export hi_Floats_set
func hi_Floats_set(handle CGoHandle, _idx int, _vl C.float) {
	s := *ptrFromHandle_hi_Floats(handle)
	s[_idx] = hi.Float(float32(_vl))
}

//export hi_Floats_append
func hi_Floats_append(handle CGoHandle, _vl C.float) {
	s := ptrFromHandle_hi_Floats(handle)
	*s = append(*s, hi.Float(float32(_vl)))
}

// ---- Maps ---

// ---- Constructors ---

//export hi_NewCouple
func hi_NewCouple(p1 CGoHandle, p2 CGoHandle) CGoHandle {
	cret := hi.NewCouple(*ptrFromHandle_hi_Person(p1), *ptrFromHandle_hi_Person(p2))

	return handleFromPtr_hi_Couple(&cret)
}

//export hi_NewPerson
func hi_NewPerson(name *C.char, age C.longlong) CGoHandle {
	return handleFromPtr_Ptr_hi_Person(hi.NewPerson(C.GoString(name), int(age)))

}

//export hi_NewActivePerson
func hi_NewActivePerson(h C.longlong) CGoHandle {
	cret, __err := hi.NewActivePerson(int(h))

	if __err != nil {
		estr := C.CString(__err.Error())
		C.PyErr_SetString(C.PyExc_RuntimeError, estr)
		C.free(unsafe.Pointer(estr))
	}
	return handleFromPtr_Ptr_hi_Person(cret)
}

//export hi_NewPersonWithAge
func hi_NewPersonWithAge(age C.longlong) CGoHandle {
	return handleFromPtr_Ptr_hi_Person(hi.NewPersonWithAge(int(age)))

}

// ---- Functions ---

//export hi_PersonAsIface
func hi_PersonAsIface(name *C.char, age C.longlong) CGoHandle {
	return handleFromPtr_hi_PersIface(hi.PersonAsIface(C.GoString(name), int(age)))

}

//export hi_Add
func hi_Add(i C.longlong, j C.longlong) C.longlong {
	return C.longlong(hi.Add(int(i), int(j)))

}

//export hi_Concat
func hi_Concat(s1 *C.char, s2 *C.char) *C.char {
	return C.CString(hi.Concat(C.GoString(s1), C.GoString(s2)))

}

//export hi_Hello
func hi_Hello(s *C.char, goRun C.char) {
	if boolPyToGo(goRun) {
		go hi.Hello(C.GoString(s))
	} else {
		hi.Hello(C.GoString(s))
	}
}

//export hi_LookupQuestion
func hi_LookupQuestion(n C.longlong) *C.char {
	cret, __err := hi.LookupQuestion(int(n))

	if __err != nil {
		estr := C.CString(__err.Error())
		C.PyErr_SetString(C.PyExc_RuntimeError, estr)
		C.free(unsafe.Pointer(estr))
	}
	return C.CString(cret)
}

//export hi_Hi
func hi_Hi(goRun C.char) {
	if boolPyToGo(goRun) {
		go hi.Hi()
	} else {
		hi.Hi()
	}
}
