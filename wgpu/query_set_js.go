package wgpu

import "syscall/js"

type QuerySet struct {
	jsValue js.Value
}

func (g *QuerySet) toJS() any {
	return g.jsValue
}

func (*QuerySet) Release() {
	// no-op
}
