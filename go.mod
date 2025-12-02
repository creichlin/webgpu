module github.com/oliverbestmann/webgpu

go 1.25

require (
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20250301202403-da16c1255728
	github.com/oliverbestmann/webgpu/libs-android v0.0.0-20251123130330-255708779f8b
	github.com/oliverbestmann/webgpu/libs-darwin v0.0.0-20251123130615-bce73357dc8d
	github.com/oliverbestmann/webgpu/libs-ios v0.0.0-20251123130710-78131685b429
	github.com/oliverbestmann/webgpu/libs-linux v0.0.0-20251123130508-41a15cd9f5d4
	github.com/oliverbestmann/webgpu/libs-windows v0.0.0-20251123134324-0b4e31ddbbf9
)

retract (
	v1.27.0 // published before deciding on a version scheme. we start at v1.0.0
)
