package wgpu

// #include "wgpu_go_wrappers.h"
import "C"
import "unsafe"

type stringView struct {
	view C.WGPUStringView
}

func stringViewOf(value string) stringView {
	return stringView{
		view: C.WGPUStringView{
			data:   C.CString(value),
			length: C.size_t(len(value)),
		},
	}
}

func (v stringView) ToC() C.WGPUStringView {
	return v.view
}

func (v stringView) Release() {
	C.free(unsafe.Pointer(v.view.data))
}
