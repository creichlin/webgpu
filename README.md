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

## Prebuild libraries

This repository uses prebuild libraries provided by `wgpu-native`.
All libraries combined are more than 512mb in size, which is more than `go get` allows
in a single library. THis is an "opinionated limited" which is not configurable.

To work around that, libraries for the different systems are split into branches.
An example is here: https://github.com/oliverbestmann/webgpu/tree/libs-linux/libs-linux

The `update-wgpu.sh` script updates all those branches and updates the go.mod file to
pull the prebuild libraries as dependencies in.

## Examples

You can find some examples in the examples directory:
https://github.com/oliverbestmann/webgpu/tree/main/examples

## Special thanks

This is a fork of [cogentcore/webgpu/](https://github.com/cogentcore/webgpu/).
Thanks to them for the work they put into WebGPU for Go.
