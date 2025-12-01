package wgpu

func (p *Buffer) MapAsync(mode MapMode, offset uint64, size uint64, callback BufferMapCallback) {
	err := p.TryMapAsync(mode, offset, size, callback)
	panicIf(err, "Buffer.MapAsync failed")
}

func (p *Buffer) Unmap() { err := p.TryUnmap(); panicIf(err, "Buffer.Unmap failed") }

func (p *CommandEncoder) ClearBuffer(buffer *Buffer, offset uint64, size uint64) {
	err := p.TryClearBuffer(buffer, offset, size)
	panicIf(err, "CommandEncoder.ClearBuffer failed")
}

func (p *CommandEncoder) CopyBufferToBuffer(source *Buffer, sourceOffset uint64, destination *Buffer, destinationOffset uint64, size uint64) {
	err := p.TryCopyBufferToBuffer(source, sourceOffset, destination, destinationOffset, size)
	panicIf(err, "CommandEncoder.CopyBufferToBuffer failed")
}

func (p *CommandEncoder) CopyBufferToTexture(source *TexelCopyBufferInfo, destination *TexelCopyTextureInfo, copySize *Extent3D) {
	err := p.TryCopyBufferToTexture(source, destination, copySize)
	panicIf(err, "CommandEncoder.CopyBufferToTexture failed")
}

func (p *CommandEncoder) CopyTextureToBuffer(source *TexelCopyTextureInfo, destination *TexelCopyBufferInfo, copySize *Extent3D) {
	err := p.TryCopyTextureToBuffer(source, destination, copySize)
	panicIf(err, "CommandEncoder.CopyTextureToBuffer failed")
}

func (p *CommandEncoder) CopyTextureToTexture(source *TexelCopyTextureInfo, destination *TexelCopyTextureInfo, copySize *Extent3D) {
	err := p.TryCopyTextureToTexture(source, destination, copySize)
	panicIf(err, "CommandEncoder.CopyTextureToTexture failed")
}

func (p *CommandEncoder) Finish(descriptor *CommandBufferDescriptor) *CommandBuffer {
	r0, err := p.TryFinish(descriptor)
	panicIf(err, "CommandEncoder.Finish failed")
	return r0
}

func (p *CommandEncoder) InsertDebugMarker(markerLabel string) {
	err := p.TryInsertDebugMarker(markerLabel)
	panicIf(err, "CommandEncoder.InsertDebugMarker failed")
}

func (p *CommandEncoder) PopDebugGroup() {
	err := p.TryPopDebugGroup()
	panicIf(err, "CommandEncoder.PopDebugGroup failed")
}

func (p *CommandEncoder) PushDebugGroup(groupLabel string) {
	err := p.TryPushDebugGroup(groupLabel)
	panicIf(err, "CommandEncoder.PushDebugGroup failed")
}

func (p *CommandEncoder) ResolveQuerySet(querySet *QuerySet, firstQuery uint32, queryCount uint32, destination *Buffer, destinationOffset uint64) {
	err := p.TryResolveQuerySet(querySet, firstQuery, queryCount, destination, destinationOffset)
	panicIf(err, "CommandEncoder.ResolveQuerySet failed")
}
func (p *ComputePassEncoder) End() { err := p.TryEnd(); panicIf(err, "ComputePassEncoder.End failed") }

func (p *Device) CreateBindGroup(descriptor *BindGroupDescriptor) *BindGroup {
	r0, err := p.TryCreateBindGroup(descriptor)
	panicIf(err, "Device.CreateBindGroup failed")
	return r0
}

func (p *Device) CreateBindGroupLayout(descriptor *BindGroupLayoutDescriptor) *BindGroupLayout {
	r0, err := p.TryCreateBindGroupLayout(descriptor)
	panicIf(err, "Device.CreateBindGroupLayout failed")
	return r0
}

func (p *Device) CreateBuffer(descriptor *BufferDescriptor) *Buffer {
	r0, err := p.TryCreateBuffer(descriptor)
	panicIf(err, "Device.CreateBuffer failed")
	return r0
}

func (p *Device) CreateCommandEncoder(descriptor *CommandEncoderDescriptor) *CommandEncoder {
	r0, err := p.TryCreateCommandEncoder(descriptor)
	panicIf(err, "Device.CreateCommandEncoder failed")
	return r0
}

func (p *Device) CreateComputePipeline(descriptor *ComputePipelineDescriptor) *ComputePipeline {
	r0, err := p.TryCreateComputePipeline(descriptor)
	panicIf(err, "Device.CreateComputePipeline failed")
	return r0
}

func (p *Device) CreatePipelineLayout(descriptor *PipelineLayoutDescriptor) *PipelineLayout {
	r0, err := p.TryCreatePipelineLayout(descriptor)
	panicIf(err, "Device.CreatePipelineLayout failed")
	return r0
}

func (p *Device) CreateQuerySet(descriptor *QuerySetDescriptor) *QuerySet {
	r0, err := p.TryCreateQuerySet(descriptor)
	panicIf(err, "Device.CreateQuerySet failed")
	return r0
}

func (p *Device) CreateRenderBundleEncoder(descriptor *RenderBundleEncoderDescriptor) *RenderBundleEncoder {
	r0, err := p.TryCreateRenderBundleEncoder(descriptor)
	panicIf(err, "Device.CreateRenderBundleEncoder failed")
	return r0
}

func (p *Device) CreateRenderPipeline(descriptor *RenderPipelineDescriptor) *RenderPipeline {
	r0, err := p.TryCreateRenderPipeline(descriptor)
	panicIf(err, "Device.CreateRenderPipeline failed")
	return r0
}

func (p *Device) CreateSampler(descriptor *SamplerDescriptor) *Sampler {
	r0, err := p.TryCreateSampler(descriptor)
	panicIf(err, "Device.CreateSampler failed")
	return r0
}

func (p *Device) CreateShaderModule(descriptor *ShaderModuleDescriptor) *ShaderModule {
	r0, err := p.TryCreateShaderModule(descriptor)
	panicIf(err, "Device.CreateShaderModule failed")
	return r0
}

func (p *Device) CreateTexture(descriptor *TextureDescriptor) *Texture {
	r0, err := p.TryCreateTexture(descriptor)
	panicIf(err, "Device.CreateTexture failed")
	return r0
}
func (p *Device) CreateBufferInit(descriptor *BufferInitDescriptor) *Buffer {
	r0, err := p.TryCreateBufferInit(descriptor)
	panicIf(err, "Device.CreateBufferInit failed")
	return r0
}

func (p *Queue) WriteBuffer(buffer *Buffer, bufferOffset uint64, data []byte) {
	err := p.TryWriteBuffer(buffer, bufferOffset, data)
	panicIf(err, "Queue.WriteBuffer failed")
}

func (p *Queue) WriteTexture(destination *TexelCopyTextureInfo, data []byte, dataLayout *TexelCopyBufferLayout, writeSize *Extent3D) {
	err := p.TryWriteTexture(destination, data, dataLayout, writeSize)
	panicIf(err, "Queue.WriteTexture failed")
}
func (p *RenderPassEncoder) End() { err := p.TryEnd(); panicIf(err, "RenderPassEncoder.End failed") }

// NOTE: you should typically not call [Texture.Release] on the returned texture.
// Instead, you should call [TextureView.Release] on any [TextureView] you create from it.
func (p *Surface) GetCurrentTexture() *Texture {
	r0, err := p.TryGetCurrentTexture()
	panicIf(err, "Surface.GetCurrentTexture failed")
	return r0
}
func (p *Texture) CreateView(descriptor *TextureViewDescriptor) *TextureView {
	r0, err := p.TryCreateView(descriptor)
	panicIf(err, "Texture.CreateView failed")
	return r0
}
