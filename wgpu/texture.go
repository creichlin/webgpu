//go:build !js

package wgpu

/*

#include <stdlib.h>
#include <wgpu.h>

extern void gowebgpu_error_callback_c(enum WGPUPopErrorScopeStatus status, WGPUErrorType type, WGPUStringView message, void * userdata, void * userdata2);

static inline WGPUTextureView gowebgpu_texture_create_view(WGPUTexture texture, WGPUTextureViewDescriptor const * descriptor, WGPUDevice device, void * error_userdata) {
	WGPUTextureView ref = NULL;
	wgpuDevicePushErrorScope(device, WGPUErrorFilter_Validation);
	ref = wgpuTextureCreateView(texture, descriptor);

	WGPUPopErrorScopeCallbackInfo const err_cb = {
		.callback = gowebgpu_error_callback_c,
		.userdata1 = error_userdata,
	};

	wgpuDevicePopErrorScope(device, err_cb);

	return ref;
}

*/
import "C"
import (
	"sync/atomic"
	"unsafe"
)

type Texture struct {
	device   *Device
	ref      C.WGPUTexture
	released int32
}

func (p *Texture) TryCreateView(descriptor *TextureViewDescriptor) (*TextureView, error) {
	var desc *C.WGPUTextureViewDescriptor

	if descriptor != nil {
		desc = &C.WGPUTextureViewDescriptor{
			format:          C.WGPUTextureFormat(descriptor.Format),
			dimension:       C.WGPUTextureViewDimension(descriptor.Dimension),
			baseMipLevel:    C.uint32_t(descriptor.BaseMipLevel),
			mipLevelCount:   C.uint32_t(descriptor.MipLevelCount),
			baseArrayLayer:  C.uint32_t(descriptor.BaseArrayLayer),
			arrayLayerCount: C.uint32_t(descriptor.ArrayLayerCount),
			aspect:          C.WGPUTextureAspect(descriptor.Aspect),
		}

		if descriptor.Label != "" {
			label := C.CString(descriptor.Label)
			defer C.free(unsafe.Pointer(label))

			desc.label.data = label
			desc.label.length = C.WGPU_STRLEN
		}
	}

	var err error = nil

	errorCallbackHandle := makeErrorCallback(&err)
	defer errorCallbackHandle.Delete()

	ref := C.gowebgpu_texture_create_view(
		p.ref,
		desc,
		p.device.ref,
		errorCallbackHandle.ToPointer(),
	)
	if err != nil {
		C.wgpuTextureViewRelease(ref)
		return nil, err
	}

	return releaseOnGC(&TextureView{ref: ref}), nil
}

func (p *Texture) Destroy() {
	C.wgpuTextureDestroy(p.ref)
}

func (p *Texture) GetDepthOrArrayLayers() uint32 {
	return uint32(C.wgpuTextureGetDepthOrArrayLayers(p.ref))
}

func (p *Texture) GetDimension() TextureDimension {
	return TextureDimension(C.wgpuTextureGetDimension(p.ref))
}

func (p *Texture) GetFormat() TextureFormat {
	return TextureFormat(C.wgpuTextureGetFormat(p.ref))
}

func (p *Texture) GetHeight() uint32 {
	return uint32(C.wgpuTextureGetHeight(p.ref))
}

func (p *Texture) GetMipLevelCount() uint32 {
	return uint32(C.wgpuTextureGetMipLevelCount(p.ref))
}

func (p *Texture) GetSampleCount() uint32 {
	return uint32(C.wgpuTextureGetSampleCount(p.ref))
}

func (p *Texture) GetUsage() TextureUsage {
	return TextureUsage(C.wgpuTextureGetUsage(p.ref))
}

func (p *Texture) GetWidth() uint32 {
	return uint32(C.wgpuTextureGetWidth(p.ref))
}

func (p *Texture) Release() {
	if p.ref != nil && atomic.CompareAndSwapInt32(&p.released, 0, 1) {
		C.wgpuTextureRelease(p.ref)
	}
}
