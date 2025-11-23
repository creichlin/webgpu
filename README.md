# WebGPU

Current upstream version: v27.0.2.0

Go bindings for WebGPU, a cross-platform, safe graphics API.

It runs natively using [wgpu-native](https://github.com/gfx-rs/wgpu-native) on Vulkan, Metal, D3D12 and OpenGL ES.
As there is still no cgo for wasm, wasm/web builds with `GOOS=js` are using the browser WebGPU interface
directly.

This fork adds garbage collection to types returned by webgpu.

For more information, see:

- [WebGPU](https://gpuweb.github.io/gpuweb/)
- [WGSL](https://gpuweb.github.io/gpuweb/wgsl/)
- [webgpu-native](https://github.com/webgpu-native/webgpu-headers)

The included static libraries downloaded from the wgpu-native project.

## Examples

You can find some examples in the examples directory:
https://github.com/oliverbestmann/webgpu/tree/main/examples

## Special thanks

This is a fork of [cogentcore/webgpu/](https://github.com/cogentcore/webgpu/).
Thanks to them for the work they put into WebGPU for Go.
