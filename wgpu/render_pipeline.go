//go:build !js

package wgpu

/*

#include <stdlib.h>
#include <wgpu.h>

*/
import "C"

func (p *RenderPipeline) GetBindGroupLayout(groupIndex uint32) *BindGroupLayout {
	ref := C.wgpuRenderPipelineGetBindGroupLayout(p.ref, C.uint32_t(groupIndex))
	if ref == nil {
		panic("Failed to acquire BindGroupLayout")
	}

	return releaseOnGC(&BindGroupLayout{ref: ref})
}
