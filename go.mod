module github.com/oliverbestmann/webgpu

go 1.25

require (
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20250301202403-da16c1255728
	github.com/oliverbestmann/webgpu/libs-android v0.0.0-20260321175638-35553c66247e
	github.com/oliverbestmann/webgpu/libs-darwin v0.0.0-20260321175628-865c2b8d8a62
	github.com/oliverbestmann/webgpu/libs-ios v0.0.0-20260321175629-0e29e7d88738
	github.com/oliverbestmann/webgpu/libs-linux v0.0.0-20260321175635-fec3b53b1724
	github.com/oliverbestmann/webgpu/libs-windows v0.0.0-20260321175633-752a2b4a21c3
)

retract v1.27.0 // published before deciding on a version scheme. we start at v1.0.0
