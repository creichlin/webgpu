//go:build !js

package wgpu

import _ "github.com/oliverbestmann/webgpu/linux/v27"

/*

// Android
#cgo android,amd64 LDFLAGS: -L${SRCDIR}/lib/android/amd64 -lwgpu_native
#cgo android,amd64 CFLAGS: -I${SRCDIR}/lib/android/amd64
#cgo android,386 LDFLAGS: -L${SRCDIR}/lib/android/386 -lwgpu_native
#cgo android,386 CFLAGS: -I${SRCDIR}/lib/android/386
#cgo android,arm64 LDFLAGS: -L${SRCDIR}/lib/android/arm64 -lwgpu_native
#cgo android,arm64 CFLAGS: -I${SRCDIR}/lib/android/arm64
#cgo android,arm LDFLAGS: -L${SRCDIR}/lib/android/arm -lwgpu_native
#cgo android,arm CFLAGS: -I${SRCDIR}/lib/android/arm

#cgo android LDFLAGS: -landroid -lm -llog

// Linux
#cgo linux,!android,amd64 LDFLAGS: -L${SRCDIR}/lib/linux/amd64 -lwgpu_native
#cgo linux,!android,amd64 CFLAGS: -I${SRCDIR}/lib/linux/amd64
#cgo linux,!android,arm64 LDFLAGS: -L${SRCDIR}/lib/linux/arm64 -lwgpu_native
#cgo linux,!android,arm64 CFLAGS: -I${SRCDIR}/lib/linux/arm64

#cgo linux,!android LDFLAGS: -lm -ldl

// iOS
#cgo ios,amd64 LDFLAGS: -L${SRCDIR}/lib/ios/amd64 -lwgpu_native
#cgo ios,amd64 CFLAGS: -I${SRCDIR}/lib/ios/amd64
#cgo ios,arm64 LDFLAGS: -L${SRCDIR}/lib/ios/arm64 -lwgpu_native
#cgo ios,arm64 CFLAGS: -I${SRCDIR}/lib/ios/arm64

// Darwin
#cgo darwin,!ios,amd64 LDFLAGS: -L${SRCDIR}/lib/darwin/amd64 -lwgpu_native
#cgo darwin,!ios,amd64 CFLAGS: -I${SRCDIR}/lib/darwin/amd64
#cgo darwin,!ios,arm64 LDFLAGS: -L${SRCDIR}/lib/darwin/arm64 -lwgpu_native
#cgo darwin,!ios,arm64 CFLAGS: -I${SRCDIR}/lib/darwin/arm64

#cgo darwin LDFLAGS: -framework Metal -framework QuartzCore

// Windows
#cgo windows,amd64 LDFLAGS: -L${SRCDIR}/lib/windows/amd64 -lwgpu_native
#cgo windows,amd64 CFLAGS: -I${SRCDIR}/lib/windows/amd64
#cgo windows,arm64 LDFLAGS: -L${SRCDIR}/lib/windows/arm64 -lwgpu_native
#cgo windows,arm64 CFLAGS: -I${SRCDIR}/lib/windows/arm64
#cgo windows,386 LDFLAGS: -L${SRCDIR}/lib/windows/386 -lwgpu_native
#cgo windows,386 CFLAGS: -I${SRCDIR}/lib/windows/386

#cgo windows LDFLAGS: -lopengl32 -lgdi32 -ld3dcompiler_47 -lws2_32 -luserenv -lbcrypt -lntdll

#include <stdio.h>
#include <wgpu.h>

#ifdef __ANDROID__
#include <android/log.h>
void logCallback_cgo(WGPULogLevel level, char const *msg) {
	switch (level) {
	case WGPULogLevel_Error:
		__android_log_write(ANDROID_LOG_ERROR, "GoLogWGPU", msg);
		break;
	case WGPULogLevel_Warn:
		__android_log_write(ANDROID_LOG_WARN, "GoLogWGPU", msg);
		break;
	default:
		__android_log_write(ANDROID_LOG_INFO, "GoLogWGPU", msg);
		break;
	}
}
#else
void logCallback_cgo(WGPULogLevel level, char const *msg) {
	char const *level_str;
	switch (level) {
	case WGPULogLevel_Error:
		level_str = "Error";
		break;
	case WGPULogLevel_Warn:
		level_str = "Warn";
		break;
	case WGPULogLevel_Info:
		level_str = "Info";
		break;
	case WGPULogLevel_Debug:
		level_str = "Debug";
		break;
	case WGPULogLevel_Trace:
		level_str = "Trace";
		break;
	default:
		level_str = "Unknown Level";
	}
	fprintf(stderr, "[wgpu] [%s] %s\n", level_str, msg);
}
#endif


*/
import "C"
import (
	"runtime"
	"sync/atomic"
)

func init() {
	C.wgpuSetLogCallback(C.WGPULogCallback(C.logCallback_cgo), nil)
}

func SetLogLevel(level LogLevel) {
	C.wgpuSetLogLevel(C.WGPULogLevel(level))
}

func GetVersion() Version {
	return Version(C.wgpuGetVersion())
}

type (
	Adapter struct {
		ref      C.WGPUAdapter
		released int32
	}
	BindGroup struct {
		ref      C.WGPUBindGroup
		released int32
	}
	BindGroupLayout struct {
		ref      C.WGPUBindGroupLayout
		released int32
	}
	CommandBuffer struct {
		ref      C.WGPUCommandBuffer
		released int32
	}
	ComputePipeline struct {
		ref      C.WGPUComputePipeline
		released int32
	}
	Device struct {
		ref      C.WGPUDevice
		released int32
	}
	Instance struct {
		ref      C.WGPUInstance
		released int32
	}
	PipelineLayout struct {
		ref      C.WGPUPipelineLayout
		released int32
	}
	QuerySet struct {
		ref      C.WGPUQuerySet
		released int32
	}
	RenderBundle struct {
		ref      C.WGPURenderBundle
		released int32
	}
	RenderBundleEncoder struct {
		ref      C.WGPURenderBundleEncoder
		released int32
	}
	RenderPipeline struct {
		ref      C.WGPURenderPipeline
		released int32
	}
	Sampler struct {
		ref      C.WGPUSampler
		released int32
	}
	ShaderModule struct {
		ref      C.WGPUShaderModule
		released int32
	}
	TextureView struct {
		ref      C.WGPUTextureView
		released int32
	}
)

func (p *Adapter) Release() {
	if p.ref != nil && atomic.CompareAndSwapInt32(&p.released, 0, 1) {
		C.wgpuAdapterRelease(p.ref)
	}
}
func (p *BindGroup) Release() {
	if p.ref != nil && atomic.CompareAndSwapInt32(&p.released, 0, 1) {
		C.wgpuBindGroupRelease(p.ref)
	}
}
func (p *BindGroupLayout) Release() {
	if p.ref != nil && atomic.CompareAndSwapInt32(&p.released, 0, 1) {
		C.wgpuBindGroupLayoutRelease(p.ref)
	}
}
func (p *CommandBuffer) Release() {
	if p.ref != nil && atomic.CompareAndSwapInt32(&p.released, 0, 1) {
		C.wgpuCommandBufferRelease(p.ref)
	}
}
func (p *ComputePipeline) Release() {
	if p.ref != nil && atomic.CompareAndSwapInt32(&p.released, 0, 1) {
		C.wgpuComputePipelineRelease(p.ref)
	}
}
func (p *Device) Release() {
	if p.ref != nil && atomic.CompareAndSwapInt32(&p.released, 0, 1) {
		C.wgpuDeviceRelease(p.ref)
	}
}
func (p *Instance) Release() {
	if p.ref != nil && atomic.CompareAndSwapInt32(&p.released, 0, 1) {
		C.wgpuInstanceRelease(p.ref)
	}
}
func (p *PipelineLayout) Release() {
	if p.ref != nil && atomic.CompareAndSwapInt32(&p.released, 0, 1) {
		C.wgpuPipelineLayoutRelease(p.ref)
	}
}
func (p *QuerySet) Release() {
	if p.ref != nil && atomic.CompareAndSwapInt32(&p.released, 0, 1) {
		C.wgpuQuerySetRelease(p.ref)
	}
}
func (p *RenderBundle) Release() {
	if p.ref != nil && atomic.CompareAndSwapInt32(&p.released, 0, 1) {
		C.wgpuRenderBundleRelease(p.ref)
	}
}
func (p *RenderBundleEncoder) Release() {
	if p.ref != nil && atomic.CompareAndSwapInt32(&p.released, 0, 1) {
		C.wgpuRenderBundleEncoderRelease(p.ref)
	}
}
func (p *RenderPipeline) Release() {
	if p.ref != nil && atomic.CompareAndSwapInt32(&p.released, 0, 1) {
		C.wgpuRenderPipelineRelease(p.ref)
	}
}
func (p *Sampler) Release() {
	if p.ref != nil && atomic.CompareAndSwapInt32(&p.released, 0, 1) {
		C.wgpuSamplerRelease(p.ref)
	}
}
func (p *ShaderModule) Release() {
	if p.ref != nil && atomic.CompareAndSwapInt32(&p.released, 0, 1) {
		C.wgpuShaderModuleRelease(p.ref)
	}
}
func (p *TextureView) Release() {
	if p.ref != nil && atomic.CompareAndSwapInt32(&p.released, 0, 1) {
		C.wgpuTextureViewRelease(p.ref)
	}
}

func (p *Device) addRef() *Device {
	if atomic.LoadInt32(&p.released) != 0 {
		panic("addRef called on a device that was already released")
	}

	C.wgpuDeviceAddRef(p.ref)

	// return a new object that can be individually garbage collected
	return releaseOnGC(&Device{ref: p.ref})
}

type releaser interface{ Release() }

func releaseNow[T releaser](value T) {
	value.Release()
}

func releaseOnGC[T releaser](value T) T {
	runtime.SetFinalizer(value, releaseNow[T])
	return value
}

// cBool converts the given Go bool to a C.WGPUBool.
func cBool(b bool) C.WGPUBool {
	if b {
		return 1
	}
	return 0
}

// goBool converts the given C.WGPUBool to a Go bool.
func goBool(b C.WGPUBool) bool {
	return b != 0
}
