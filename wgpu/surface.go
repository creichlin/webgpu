//go:build !js

package wgpu

/*

#include <stdlib.h>
#include <wgpu.h>

extern void gowebgpu_error_callback_c(enum WGPUPopErrorScopeStatus status, WGPUErrorType type, WGPUStringView message, void * userdata, void * userdata2);

static inline WGPUTexture gowebgpu_surface_get_current_texture(WGPUSurface surface, WGPUDevice device, void * error_userdata) {
	WGPUSurfaceTexture ref;
	wgpuDevicePushErrorScope(device, WGPUErrorFilter_Validation);
	wgpuSurfaceGetCurrentTexture(surface, &ref);

	WGPUPopErrorScopeCallbackInfo const err_cb = {
		.callback = gowebgpu_error_callback_c,
		.userdata1 = error_userdata,
	};

	wgpuDevicePopErrorScope(device, err_cb);

	return ref.texture;
}

*/
import "C"
import (
	"errors"
	"runtime"
	"sync/atomic"
	"unsafe"
)

type Surface struct {
	device   *Device
	ref      C.WGPUSurface
	released int32
}

func (p *Surface) GetCapabilities(adapter *Adapter) (ret SurfaceCapabilities) {
	var caps C.WGPUSurfaceCapabilities
	C.wgpuSurfaceGetCapabilities(p.ref, adapter.ref, &caps)

	if caps.alphaModeCount == 0 && caps.formatCount == 0 && caps.presentModeCount == 0 {
		return
	}
	if caps.formatCount > 0 {
		caps.formats = (*C.WGPUTextureFormat)(C.malloc(C.size_t(unsafe.Sizeof(C.WGPUTextureFormat(0))) * caps.formatCount))
		defer C.free(unsafe.Pointer(caps.formats))
	}
	if caps.presentModeCount > 0 {
		caps.presentModes = (*C.WGPUPresentMode)(C.malloc(C.size_t(unsafe.Sizeof(C.WGPUPresentMode(0))) * caps.presentModeCount))
		defer C.free(unsafe.Pointer(caps.presentModes))
	}
	if caps.alphaModeCount > 0 {
		caps.alphaModes = (*C.WGPUCompositeAlphaMode)(C.malloc(C.size_t(unsafe.Sizeof(C.WGPUCompositeAlphaMode(0))) * caps.alphaModeCount))
		defer C.free(unsafe.Pointer(caps.alphaModes))
	}

	C.wgpuSurfaceGetCapabilities(p.ref, adapter.ref, &caps)

	if caps.formatCount > 0 {
		formatsTmp := unsafe.Slice((*TextureFormat)(caps.formats), caps.formatCount)
		ret.Formats = make([]TextureFormat, caps.formatCount)
		copy(ret.Formats, formatsTmp)
	}
	if caps.presentModeCount > 0 {
		presentModesTmp := unsafe.Slice((*PresentMode)(caps.presentModes), caps.presentModeCount)
		ret.PresentModes = make([]PresentMode, caps.presentModeCount)
		copy(ret.PresentModes, presentModesTmp)
	}
	if caps.alphaModeCount > 0 {
		alphaModesTmp := unsafe.Slice((*CompositeAlphaMode)(caps.alphaModes), caps.alphaModeCount)
		ret.AlphaModes = make([]CompositeAlphaMode, caps.alphaModeCount)
		copy(ret.AlphaModes, alphaModesTmp)
	}

	return
}

func (p *Surface) Configure(device *Device, config *SurfaceConfiguration) {
	p.device = device.addRef()

	var pinner runtime.Pinner
	defer pinner.Unpin()

	var cfg *C.WGPUSurfaceConfiguration
	if config != nil {
		var nextInChain *C.WGPUSurfaceConfigurationExtras

		if config.DesiredMaximumFrameLatency > 0 {
			nextInChain = &C.WGPUSurfaceConfigurationExtras{
				chain: C.WGPUChainedStruct{
					sType: C.WGPUSType_SurfaceConfigurationExtras,
				},
				desiredMaximumFrameLatency: 1,
			}

			pinner.Pin(nextInChain)
		}

		cfg = &C.WGPUSurfaceConfiguration{
			device:      p.device.ref,
			format:      C.WGPUTextureFormat(config.Format),
			usage:       C.WGPUTextureUsage(config.Usage),
			alphaMode:   C.WGPUCompositeAlphaMode(config.AlphaMode),
			width:       C.uint32_t(config.Width),
			height:      C.uint32_t(config.Height),
			presentMode: C.WGPUPresentMode(config.PresentMode),
			nextInChain: (*C.WGPUChainedStruct)(unsafe.Pointer(nextInChain)),
		}

		if len(config.ViewFormats) > 0 {
			pinner.Pin(&config.ViewFormats[0])

			cfg.viewFormatCount = C.size_t(len(config.ViewFormats))
			cfg.viewFormats = (*C.WGPUTextureFormat)(&config.ViewFormats[0])
		}
	}

	C.wgpuSurfaceConfigure(p.ref, cfg)
}

// NOTE: you should typically not call [Texture.Release] on the returned texture.
// Instead, you should call [TextureView.Release] on any [TextureView] you create from it.
func (p *Surface) TryGetCurrentTexture() (*Texture, error) {
	if p.device == nil {
		return nil, errors.New("surface not configured")
	}

	var err error = nil

	errorCallbackHandle := makeErrorCallback(&err)
	defer errorCallbackHandle.Delete()

	ref := C.gowebgpu_surface_get_current_texture(
		p.ref,
		p.device.ref,
		errorCallbackHandle.ToPointer(),
	)
	if err != nil {
		if ref != nil {
			C.wgpuTextureRelease(ref)
		}
		return nil, err
	}

	// setting released to one forces Texture.Release() to be a noop
	return &Texture{device: p.device.addRef(), ref: ref, released: 0}, nil
}

func (p *Surface) Present() {
	C.wgpuSurfacePresent(p.ref)
}

func (p *Surface) Release() {
	if p.ref != nil && atomic.CompareAndSwapInt32(&p.released, 0, 1) {
		C.wgpuSurfaceRelease(p.ref)
	}
}
