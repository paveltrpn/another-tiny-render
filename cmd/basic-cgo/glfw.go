package main

import (
	"another-tiny-render/internal/vk"
	"unsafe"

	"github.com/go-gl/glfw/v3.3/glfw"
)

// This function cast *VkSurfaceKHR returned by golang GLFW bindings and
// represented as uintptr to "normal" form.
//
// GLFW return newly created VkSurfaceKHR as uintptr value. Why?
// But we must pass *VkSurfaceKHR to vkDestrySurfaceKHR function instead of uintptr.
// Golang SDL2 bindings return this guy as unsafe.Pointer value indeed.
//
// And why go vet tell "passible missuse of unsafe.Pointer" here???
func CreateGLFWvkPSurfaceKHR(wnd *glfw.Window, instance vk.PInstance) (vk.PSurfaceKHR, error) {
	vkSurfaceGLFW, e := wnd.CreateWindowSurface(vkInstance, nil)
	return *(*vk.PSurfaceKHR)(unsafe.Pointer(vkSurfaceGLFW)), e
}
