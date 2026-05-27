package wgpu

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkCGoOverhead(b *testing.B) {
	instance := CreateInstance(nil)

	adapter, err := instance.RequestAdapter(&RequestAdapterOptions{
		ForceFallbackAdapter: true,
	})
	require.NoError(b, err)
	defer adapter.Release()

	device, err := adapter.RequestDevice(nil)
	require.NoError(b, err)
	defer device.Release()

	queue := device.GetQueue()
	defer queue.Release()

	b.ReportAllocs()

	texture := device.CreateTexture(&TextureDescriptor{
		Label:         "test",
		Usage:         TextureUsageCopyDst,
		Dimension:     TextureDimension2D,
		Size:          Extent3D{Width: 16, Height: 16, DepthOrArrayLayers: 1},
		Format:        TextureFormatRGBA8Unorm,
		MipLevelCount: 1,
		SampleCount:   1,
	})

	for range b.N {
		view := texture.CreateView(&TextureViewDescriptor{
			Label:           "test label",
			MipLevelCount:   1,
			ArrayLayerCount: 1,
		})

		view.Release()
	}
}
