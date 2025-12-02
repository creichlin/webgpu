//go:build js

package wgpu

// GetQueue returns a Queue as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-queue
func (g *Device) GetQueue() *Queue {
	jsQueue := g.jsValue.Get("queue")
	return &Queue{
		jsValue: jsQueue,
	}
}

func (g *Device) TryCreateQuerySet(descriptor *QuerySetDescriptor) (*QuerySet, error) {
	jsQuerySet := g.jsValue.Call("createQuerySet", pointerToJS(descriptor))
	return &QuerySet{jsQuerySet}, nil
}

// TryCreateCommandEncoder as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createcommandencoder
func (g *Device) TryCreateCommandEncoder(descriptor *CommandEncoderDescriptor) (*CommandEncoder, error) {
	jsEncoder := g.jsValue.Call("createCommandEncoder", pointerToJS(descriptor))
	return &CommandEncoder{
		jsValue: jsEncoder,
	}, nil
}

// TryCreateBuffer as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createbuffer
func (g *Device) TryCreateBuffer(descriptor *BufferDescriptor) (*Buffer, error) {
	jsBuffer := g.jsValue.Call("createBuffer", pointerToJS(descriptor))
	return &Buffer{
		jsValue: jsBuffer,
	}, nil
}

// TryCreateShaderModule as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createshadermodule
func (g *Device) TryCreateShaderModule(desc *ShaderModuleDescriptor) (*ShaderModule, error) {
	jsShader := g.jsValue.Call("createShaderModule", pointerToJS(desc))
	return &ShaderModule{
		jsValue: jsShader,
	}, nil
}

// TryCreateRenderPipeline as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createrenderpipeline
func (g *Device) TryCreateRenderPipeline(descriptor *RenderPipelineDescriptor) (*RenderPipeline, error) {
	jsPipeline := g.jsValue.Call("createRenderPipeline", pointerToJS(descriptor))
	return &RenderPipeline{
		jsValue: jsPipeline,
	}, nil
}

// TryCreateBindGroup as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createbindgroup
func (g *Device) TryCreateBindGroup(descriptor *BindGroupDescriptor) (*BindGroup, error) {
	jsBindGroup := g.jsValue.Call("createBindGroup", pointerToJS(descriptor))
	return &BindGroup{
		jsValue: jsBindGroup,
	}, nil
}

// TryCreateBindGroupLayout as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createbindgrouplayout
func (g *Device) TryCreateBindGroupLayout(descriptor *BindGroupLayoutDescriptor) (*BindGroupLayout, error) {
	jsLayout := g.jsValue.Call("createBindGroupLayout", pointerToJS(descriptor))
	return &BindGroupLayout{
		jsValue: jsLayout,
	}, nil
}

// TryCreatePipelineLayout as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createpipelinelayout
func (g *Device) TryCreatePipelineLayout(descriptor *PipelineLayoutDescriptor) (*PipelineLayout, error) {
	jsLayout := g.jsValue.Call("createPipelineLayout", pointerToJS(descriptor))
	return &PipelineLayout{
		jsValue: jsLayout,
	}, nil
}

// TryCreateComputePipeline as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createcomputepipeline
func (g *Device) TryCreateComputePipeline(descriptor *ComputePipelineDescriptor) (*ComputePipeline, error) {
	jsPipeline := g.jsValue.Call("createComputePipeline", pointerToJS(descriptor))
	return &ComputePipeline{
		jsValue: jsPipeline,
	}, nil
}

// TryCreateTexture as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createtexture
func (g *Device) TryCreateTexture(descriptor *TextureDescriptor) (*Texture, error) {
	jsTexture := g.jsValue.Call("createTexture", pointerToJS(descriptor))
	return &Texture{
		jsValue: jsTexture,
	}, nil
}

// TryCreateSampler as described:
// https://gpuweb.github.io/gpuweb/#dom-gpudevice-createsampler
func (g *Device) TryCreateSampler(descriptor *SamplerDescriptor) (*Sampler, error) {
	jsSampler := g.jsValue.Call("createSampler", pointerToJS(descriptor))
	return &Sampler{
		jsValue: jsSampler,
	}, nil
}

func (g *Device) GetLimits() Limits {
	return limitsFromJS(g.jsValue.Get("limits"))
}

func (g *Device) Poll(wait bool, wrappedSubmissionIndex *uint64) (queueEmpty bool) {
	return false // no-op
}

func (g *Device) TryCreateRenderBundleEncoder(descriptor *RenderBundleEncoderDescriptor) (*RenderBundleEncoder, error) {
	// TODO implement this
	panic("unimplemented")
}
