package main

import (
	"fmt"
	"log"
	"runtime"

	"another-tiny-render/internal/vk"

	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	vkInstance vk.PInstance
	vkSurface  vk.PSurfaceKHR
)

func doVulkanStuff(wnd *glfw.Window) {
	var (
		createInfo         vk.InstanceCreateInfo
		appInfo            vk.ApplicationInfo
		extNames, lrsNames []string
	)

	extCount, extensions, _ := vk.EnumerateInstanceExtensionProperties(nil)
	extNames = make([]string, 0)

	fmt.Printf("extensions count - %v\n", extCount)
	for _, ext := range extensions {
		fmt.Printf("%v\n", ext.ExtensionName)
		extNames = append(extNames, ext.ExtensionName)
	}

	lrsCount, layers, _ := vk.EnumerateInstanceLayerProperties()
	lrsNames = make([]string, 0)

	fmt.Printf("layers count - %v\n", lrsCount)
	for _, lr := range layers {
		fmt.Printf("%v\n", lr.LayerName)
		lrsNames = append(lrsNames, lr.LayerName)
	}

	appInfo.SType = vk.STRUCTURE_TYPE_APPLICATION_INFO
	appInfo.PApplicationName = "basic"
	appInfo.EngineVersion = vk.MakeAPIVersion(0, 1, 0, 0)
	appInfo.PEngineName = "no engine"
	appInfo.EngineVersion = vk.MakeAPIVersion(0, 1, 0, 0)
	appInfo.ApiVersion = vk.MakeAPIVersion(0, 1, 0, 0)

	createInfo.SType = vk.STRUCTURE_TYPE_INSTANCE_CREATE_INFO
	createInfo.Flags = 0
	createInfo.PApplicationInfo = &appInfo
	createInfo.EnabledLayerCount = lrsCount
	createInfo.PpEnabledLayerNames = lrsNames
	createInfo.EnabledExtensionCount = extCount
	createInfo.PpEnabledExtensionNames = extNames

	var res vk.Result
	vkInstance, res = vk.CreateInstance(createInfo, nil)

	if res != vk.SUCCESS {
		log.Fatalf("error creating instance with - %v\n", res)
	} else {
		fmt.Printf("creating instalnce success with - %v\n", vkInstance)
	}

	devCount, devs, _ := vk.EnumeratePhysicalDevices(vkInstance)
	fmt.Printf("physical devices count - %v\n", devCount)

	devProps := vk.GetPhysicalDeviceProperties(devs[0])
	fmt.Printf("device name - %v\n", devProps.DeviceName)

	devFeats := vk.GetPhysicalDeviceFeatures(devs[0])
	fmt.Printf("device features - %v\n", devFeats)

	queueCount, _ := vk.GetPhysicalDeviceQueueFamilyProperties(devs[0])
	fmt.Printf("queue count - %v\n", queueCount)

	vkSurface, _ = CreateGLFWvkPSurfaceKHR(wnd, vkInstance)

	surfaceCapKHR, res := vk.GetPhysicalDeviceSurfaceCapabilitiesKHR(devs[0], vkSurface)
	fmt.Printf("phys dev surface capabilities - %v\n", surfaceCapKHR)
}

func main() {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		log.Fatalln("Error! Can't init glfw!")
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	glfw.WindowHint(glfw.DoubleBuffer, glfw.True)

	wnd, err := glfw.CreateWindow(800, 600, "basic-cgo", nil, nil)
	if err != nil {
		log.Fatalln("Error! Can't create window!")
	}

	glfwExtensions := wnd.GetRequiredInstanceExtensions()
	fmt.Printf("glfw instance extensions - %v\n", glfwExtensions)

	keyCallback := func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if key == glfw.KeyEscape {
			w.SetShouldClose(true)
		}
	}
	wnd.SetKeyCallback(keyCallback)

	doVulkanStuff(wnd)

	for !wnd.ShouldClose() {
		glfw.PollEvents()
	}

	vk.DestroySurfaceKHR(vkInstance, vkSurface, nil)
	vk.DestroyInstance(vkInstance, nil)

	glfw.Terminate()
}
