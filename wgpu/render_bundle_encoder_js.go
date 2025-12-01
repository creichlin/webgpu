package wgpu

import "syscall/js"

type RenderBundleEncoder struct {
	jsValue js.Value
}

func (g *RenderBundleEncoder) Release() {}
