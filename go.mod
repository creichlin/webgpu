module github.com/oliverbestmann/webgpu

go 1.25

require (
	github.com/go-gl/glfw/v3.4/glfw v0.1.0-pre.1.0.20260406072232-3ac4aa2bb164
	github.com/oliverbestmann/webgpu/libs-android v0.0.0-20260509160813-48db59792a15
	github.com/oliverbestmann/webgpu/libs-darwin v0.0.0-20260509160802-b09403b07cd3
	github.com/oliverbestmann/webgpu/libs-ios v0.0.0-20260509160803-765e39d2a48b
	github.com/oliverbestmann/webgpu/libs-linux v0.0.0-20260509160809-2fefaf7c9ead
	github.com/oliverbestmann/webgpu/libs-windows v0.0.0-20260509160807-0bc32b12c7bc
	github.com/stretchr/testify v1.11.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract v1.27.0 // published before deciding on a version scheme. we start at v1.0.0
