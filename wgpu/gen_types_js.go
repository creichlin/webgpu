//go:build js

package wgpu

import "syscall/js"

type Buffer struct {
	jsValue js.Value
}

func (g *Buffer) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *Buffer) toJS() any {
	return g.jsValue
}

type CommandEncoder struct {
	jsValue js.Value
}

func (g *CommandEncoder) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *CommandEncoder) toJS() any {
	return g.jsValue
}

type ComputePassEncoder struct {
	jsValue js.Value
}

func (g *ComputePassEncoder) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *ComputePassEncoder) toJS() any {
	return g.jsValue
}

type Queue struct {
	jsValue js.Value
}

func (g *Queue) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *Queue) toJS() any {
	return g.jsValue
}

type RenderPassEncoder struct {
	jsValue js.Value
}

func (g *RenderPassEncoder) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *RenderPassEncoder) toJS() any {
	return g.jsValue
}

type Surface struct {
	jsValue js.Value
}

func (g *Surface) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *Surface) toJS() any {
	return g.jsValue
}

type Texture struct {
	jsValue js.Value
}

func (g *Texture) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *Texture) toJS() any {
	return g.jsValue
}

type Adapter struct {
	jsValue js.Value
}

func (g *Adapter) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *Adapter) toJS() any {
	return g.jsValue
}

type BindGroup struct {
	jsValue js.Value
}

func (g *BindGroup) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *BindGroup) toJS() any {
	return g.jsValue
}

type BindGroupLayout struct {
	jsValue js.Value
}

func (g *BindGroupLayout) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *BindGroupLayout) toJS() any {
	return g.jsValue
}

type CommandBuffer struct {
	jsValue js.Value
}

func (g *CommandBuffer) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *CommandBuffer) toJS() any {
	return g.jsValue
}

type ComputePipeline struct {
	jsValue js.Value
}

func (g *ComputePipeline) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *ComputePipeline) toJS() any {
	return g.jsValue
}

type Device struct {
	jsValue js.Value
}

func (g *Device) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *Device) toJS() any {
	return g.jsValue
}

type Instance struct {
	jsValue js.Value
}

func (g *Instance) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *Instance) toJS() any {
	return g.jsValue
}

type PipelineLayout struct {
	jsValue js.Value
}

func (g *PipelineLayout) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *PipelineLayout) toJS() any {
	return g.jsValue
}

type QuerySet struct {
	jsValue js.Value
}

func (g *QuerySet) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *QuerySet) toJS() any {
	return g.jsValue
}

type RenderBundle struct {
	jsValue js.Value
}

func (g *RenderBundle) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *RenderBundle) toJS() any {
	return g.jsValue
}

type RenderBundleEncoder struct {
	jsValue js.Value
}

func (g *RenderBundleEncoder) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *RenderBundleEncoder) toJS() any {
	return g.jsValue
}

type RenderPipeline struct {
	jsValue js.Value
}

func (g *RenderPipeline) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *RenderPipeline) toJS() any {
	return g.jsValue
}

type Sampler struct {
	jsValue js.Value
}

func (g *Sampler) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *Sampler) toJS() any {
	return g.jsValue
}

type ShaderModule struct {
	jsValue js.Value
}

func (g *ShaderModule) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *ShaderModule) toJS() any {
	return g.jsValue
}

type TextureView struct {
	jsValue js.Value
}

func (g *TextureView) Release() {
	// no-op, just here to keep api compatibly with native version
}

func (g *TextureView) toJS() any {
	return g.jsValue
}
