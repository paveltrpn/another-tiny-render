package vk

/*
#cgo LDFLAGS: -lX11 -lXi -lXrandr -lXxf86vm -lXinerama -lXcursor -lrt -lvulkan
#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include <vulkan/vulkan.h>
*/
import "C"

import (
	"unsafe"
)

func MakeAPIVersion(variant, major, minor, patch uint32) uint32 {
	return (((variant) << 29) | ((major) << 22) | ((minor) << 12) | (patch))
}

func EnumerateInstanceExtensionProperties(pLayerName *uint8) (uint32, []ExtensionProperties, Result) {
	var (
		pPropertyCount C.uint32_t
		goProperties   []ExtensionProperties
		pProperties    *C.VkExtensionProperties
	)

	err := C.vkEnumerateInstanceExtensionProperties(
		(*C.char)(unsafe.Pointer(pLayerName)),
		&pPropertyCount,
		nil)
	if err != SUCCESS {
		return 0, nil, Result(err)
	}

	pProperties = allocVkStructs(pProperties, uint32(pPropertyCount))
	defer cFree(pProperties)

	err = C.vkEnumerateInstanceExtensionProperties(
		(*C.char)(unsafe.Pointer(pLayerName)),
		(*C.uint32_t)(&pPropertyCount),
		pProperties)
	if err != SUCCESS {
		return 0, nil, Result(err)
	}

	goProperties = make([]ExtensionProperties, pPropertyCount)
	decoded := unsafe.Slice(pProperties, int(pPropertyCount))

	for i := range goProperties {
		goProperties[i].ExtensionName = C.GoString(&decoded[i].extensionName[0])
		goProperties[i].SpecVersion = uint32(decoded[i].specVersion)
	}

	return uint32(pPropertyCount), goProperties, Result(err)
}

func EnumerateInstanceLayerProperties() (uint32, []LayerProperties, Result) {
	var (
		pPropertyCount C.uint32_t
		goProperties   []LayerProperties
		pProperties    *C.VkLayerProperties
	)

	err := C.vkEnumerateInstanceLayerProperties(&pPropertyCount, nil)
	if err != SUCCESS {
		return 0, nil, Result(err)
	}

	pProperties = allocVkStructs(pProperties, uint32(pPropertyCount))
	defer cFree(pProperties)

	err = C.vkEnumerateInstanceLayerProperties(&pPropertyCount, pProperties)
	if err != SUCCESS {
		return 0, nil, Result(err)
	}

	goProperties = make([]LayerProperties, pPropertyCount)
	decoded := unsafe.Slice(pProperties, int(pPropertyCount))

	for i := range goProperties {
		goProperties[i].LayerName = C.GoString(&decoded[i].layerName[0])
		goProperties[i].SpecVersion = uint32(decoded[i].specVersion)
		goProperties[i].ImplementationVersion = uint32(decoded[i].implementationVersion)
		goProperties[i].Description = C.GoString(&decoded[i].description[0])
	}

	return uint32(pPropertyCount), goProperties, SUCCESS
}

func CreateInstance(createInfo InstanceCreateInfo, pAllocator *AllocationCallbacks) (PInstance, Result) {

	var (
		pCreateInfo C.VkInstanceCreateInfo
		// using pointer to C.VkApplicationInfo instead of stack allocated
		// value because of "cgo argument has Go pointer to Go pointer" panic
		pApplicationInfo *C.VkApplicationInfo
		// pAllocationCallbacks *C.VkAllocationCallbacks
		pInstance C.VkInstance
	)

	pApplicationInfo = allocVkStructs(pApplicationInfo, 1)
	defer cFree(pApplicationInfo)

	pApplicationInfo.sType = C.VkStructureType(createInfo.PApplicationInfo.SType)
	pApplicationInfo.pNext = nil
	pApplicationInfo.pApplicationName = C.CString(createInfo.PApplicationInfo.PApplicationName)
	defer cFree(pApplicationInfo.pApplicationName)
	pApplicationInfo.applicationVersion = C.uint(createInfo.PApplicationInfo.ApplicationVersion)
	pApplicationInfo.pEngineName = C.CString(createInfo.PApplicationInfo.PEngineName)
	defer cFree(pApplicationInfo.pEngineName)
	pApplicationInfo.engineVersion = C.uint(createInfo.PApplicationInfo.EngineVersion)
	pApplicationInfo.apiVersion = C.uint(createInfo.PApplicationInfo.ApiVersion)

	pCreateInfo.sType = C.VkStructureType(createInfo.SType)
	pCreateInfo.pNext = nil
	pCreateInfo.flags = C.uint(createInfo.Flags)
	pCreateInfo.pApplicationInfo = pApplicationInfo
	pCreateInfo.enabledLayerCount = C.uint(createInfo.EnabledLayerCount)
	// FIXME!!! - free ppEnabledLayerNames **char
	pCreateInfo.ppEnabledLayerNames = strSliceToCArrray(createInfo.PpEnabledLayerNames)
	pCreateInfo.enabledExtensionCount = C.uint(createInfo.EnabledExtensionCount)
	// FIXME!!! - free ppEnabledExtensionNames **char
	pCreateInfo.ppEnabledExtensionNames = strSliceToCArrray(createInfo.PpEnabledExtensionNames)

	err := C.vkCreateInstance(&pCreateInfo, nil, &pInstance)

	return (PInstance)(pInstance), Result(err)
}

func DestroyInstance(instance PInstance, pAllocator *AllocationCallbacks) {
	C.vkDestroyInstance(instance, (*C.struct_VkAllocationCallbacks)(pAllocator))
}

func EnumeratePhysicalDevices(instance PInstance) (uint32, []PhysicalDevice, Result) {
	var (
		pPhysicalDeviceCount C.uint32_t
		pPhysicalDevices     *C.VkPhysicalDevice
		goPhysicalDevices    []PhysicalDevice
	)

	err := C.vkEnumeratePhysicalDevices(instance, &pPhysicalDeviceCount, nil)
	if err != SUCCESS {
		return 0, nil, Result(err)
	}

	pPhysicalDevices = allocVkStructs(pPhysicalDevices, uint32(pPhysicalDeviceCount))
	defer cFree(pPhysicalDevices)

	err = C.vkEnumeratePhysicalDevices(instance, &pPhysicalDeviceCount, pPhysicalDevices)
	if err != SUCCESS {
		return 0, nil, Result(err)
	}

	goPhysicalDevices = make([]PhysicalDevice, uint32(pPhysicalDeviceCount))
	decoded := unsafe.Slice(pPhysicalDevices, int(pPhysicalDeviceCount))

	for i := range goPhysicalDevices {
		goPhysicalDevices[i] = PhysicalDevice(decoded[i])
	}

	return uint32(pPhysicalDeviceCount), goPhysicalDevices, SUCCESS
}

func GetPhysicalDeviceProperties(physicalDevice PhysicalDevice) PhysicalDeviceProperties {
	var (
		pProperties  C.VkPhysicalDeviceProperties
		goProperties PhysicalDeviceProperties
	)

	C.vkGetPhysicalDeviceProperties(physicalDevice, &pProperties)

	goProperties.ApiVersion = uint32(pProperties.apiVersion)
	goProperties.DriverVersion = uint32(pProperties.driverVersion)
	goProperties.VendorID = uint32(pProperties.vendorID)
	goProperties.DeviceID = uint32(pProperties.deviceID)
	goProperties.DeviceType = PhysicalDeviceType(pProperties.deviceType)
	goProperties.DeviceName = C.GoString(&pProperties.deviceName[0])
	// FIXME!!!
	//goProperties.PipelineCacheUUID = append(goProperties.PipelineCacheUUID, 0)

	goProperties.Limits.MaxImageDimension1D = uint32(pProperties.limits.maxImageDimension1D)
	goProperties.Limits.MaxImageDimension2D = uint32(pProperties.limits.maxImageDimension2D)
	goProperties.Limits.MaxImageDimension3D = uint32(pProperties.limits.maxImageDimension3D)
	goProperties.Limits.MaxImageDimensionCube = uint32(pProperties.limits.maxImageDimensionCube)
	goProperties.Limits.MaxImageArrayLayers = uint32(pProperties.limits.maxImageArrayLayers)
	goProperties.Limits.MaxTexelBufferElements = uint32(pProperties.limits.maxTexelBufferElements)
	goProperties.Limits.MaxUniformBufferRange = uint32(pProperties.limits.maxUniformBufferRange)
	goProperties.Limits.MaxStorageBufferRange = uint32(pProperties.limits.maxStorageBufferRange)
	goProperties.Limits.MaxPushConstantsSize = uint32(pProperties.limits.maxPushConstantsSize)
	goProperties.Limits.MaxMemoryAllocationCount = uint32(pProperties.limits.maxMemoryAllocationCount)
	goProperties.Limits.MaxSamplerAllocationCount = uint32(pProperties.limits.maxSamplerAllocationCount)
	goProperties.Limits.BufferImageGranularity = DeviceSize(pProperties.limits.bufferImageGranularity)
	goProperties.Limits.SparseAddressSpaceSize = DeviceSize(pProperties.limits.sparseAddressSpaceSize)
	goProperties.Limits.MaxBoundDescriptorSets = uint32(pProperties.limits.maxBoundDescriptorSets)
	goProperties.Limits.MaxPerStageDescriptorSamplers = uint32(pProperties.limits.maxPerStageDescriptorSamplers)
	goProperties.Limits.MaxPerStageDescriptorUniformBuffers = uint32(pProperties.limits.maxPerStageDescriptorUniformBuffers)
	goProperties.Limits.MaxPerStageDescriptorStorageBuffers = uint32(pProperties.limits.maxPerStageDescriptorStorageBuffers)
	goProperties.Limits.MaxPerStageDescriptorSampledImages = uint32(pProperties.limits.maxPerStageDescriptorSampledImages)
	goProperties.Limits.MaxPerStageDescriptorStorageImages = uint32(pProperties.limits.maxPerStageDescriptorStorageImages)
	goProperties.Limits.MaxPerStageDescriptorInputAttachments = uint32(pProperties.limits.maxPerStageDescriptorInputAttachments)
	goProperties.Limits.MaxPerStageResources = uint32(pProperties.limits.maxPerStageResources)
	goProperties.Limits.MaxDescriptorSetSamplers = uint32(pProperties.limits.maxDescriptorSetSamplers)
	goProperties.Limits.MaxDescriptorSetUniformBuffers = uint32(pProperties.limits.maxDescriptorSetUniformBuffers)
	goProperties.Limits.MaxDescriptorSetUniformBuffersDynamic = uint32(pProperties.limits.maxDescriptorSetUniformBuffersDynamic)
	goProperties.Limits.MaxDescriptorSetStorageBuffers = uint32(pProperties.limits.maxDescriptorSetStorageBuffers)
	goProperties.Limits.MaxDescriptorSetStorageBuffersDynamic = uint32(pProperties.limits.maxDescriptorSetStorageBuffersDynamic)
	goProperties.Limits.MaxDescriptorSetSampledImages = uint32(pProperties.limits.maxDescriptorSetSampledImages)
	goProperties.Limits.MaxDescriptorSetStorageImages = uint32(pProperties.limits.maxDescriptorSetStorageImages)
	goProperties.Limits.MaxDescriptorSetInputAttachments = uint32(pProperties.limits.maxDescriptorSetInputAttachments)
	goProperties.Limits.MaxVertexInputAttributes = uint32(pProperties.limits.maxVertexInputAttributes)
	goProperties.Limits.MaxVertexInputBindings = uint32(pProperties.limits.maxVertexInputBindings)
	goProperties.Limits.MaxVertexInputAttributeOffset = uint32(pProperties.limits.maxVertexInputAttributeOffset)
	goProperties.Limits.MaxVertexInputBindingStride = uint32(pProperties.limits.maxVertexInputBindingStride)
	goProperties.Limits.MaxVertexOutputComponents = uint32(pProperties.limits.maxVertexOutputComponents)
	goProperties.Limits.MaxTessellationGenerationLevel = uint32(pProperties.limits.maxTessellationGenerationLevel)
	goProperties.Limits.MaxTessellationPatchSize = uint32(pProperties.limits.maxTessellationPatchSize)
	goProperties.Limits.MaxTessellationControlPerVertexInputComponents = uint32(pProperties.limits.maxTessellationControlPerVertexInputComponents)
	goProperties.Limits.MaxTessellationControlPerVertexOutputComponents = uint32(pProperties.limits.maxTessellationControlPerVertexOutputComponents)
	goProperties.Limits.MaxTessellationControlPerPatchOutputComponents = uint32(pProperties.limits.maxTessellationControlPerPatchOutputComponents)
	goProperties.Limits.MaxTessellationControlTotalOutputComponents = uint32(pProperties.limits.maxTessellationControlTotalOutputComponents)
	goProperties.Limits.MaxTessellationEvaluationInputComponents = uint32(pProperties.limits.maxTessellationEvaluationInputComponents)
	goProperties.Limits.MaxTessellationEvaluationOutputComponents = uint32(pProperties.limits.maxTessellationEvaluationOutputComponents)
	goProperties.Limits.MaxGeometryShaderInvocations = uint32(pProperties.limits.maxGeometryShaderInvocations)
	goProperties.Limits.MaxGeometryInputComponents = uint32(pProperties.limits.maxGeometryInputComponents)
	goProperties.Limits.MaxGeometryOutputComponents = uint32(pProperties.limits.maxGeometryOutputComponents)
	goProperties.Limits.MaxGeometryOutputVertices = uint32(pProperties.limits.maxGeometryOutputVertices)
	goProperties.Limits.MaxGeometryTotalOutputComponents = uint32(pProperties.limits.maxGeometryTotalOutputComponents)
	goProperties.Limits.MaxFragmentInputComponents = uint32(pProperties.limits.maxFragmentInputComponents)
	goProperties.Limits.MaxFragmentOutputAttachments = uint32(pProperties.limits.maxFragmentOutputAttachments)
	goProperties.Limits.MaxFragmentDualSrcAttachments = uint32(pProperties.limits.maxFragmentDualSrcAttachments)
	goProperties.Limits.MaxFragmentCombinedOutputResources = uint32(pProperties.limits.maxFragmentCombinedOutputResources)
	goProperties.Limits.MaxComputeSharedMemorySize = uint32(pProperties.limits.maxComputeSharedMemorySize)
	goProperties.Limits.MaxComputeWorkGroupCount[0] = uint32(pProperties.limits.maxComputeWorkGroupCount[0])
	goProperties.Limits.MaxComputeWorkGroupCount[1] = uint32(pProperties.limits.maxComputeWorkGroupCount[1])
	goProperties.Limits.MaxComputeWorkGroupCount[2] = uint32(pProperties.limits.maxComputeWorkGroupCount[2])
	goProperties.Limits.MaxComputeWorkGroupInvocations = uint32(pProperties.limits.maxComputeWorkGroupInvocations)
	goProperties.Limits.MaxComputeWorkGroupSize[0] = uint32(pProperties.limits.maxComputeWorkGroupSize[0])
	goProperties.Limits.MaxComputeWorkGroupSize[1] = uint32(pProperties.limits.maxComputeWorkGroupSize[1])
	goProperties.Limits.MaxComputeWorkGroupSize[2] = uint32(pProperties.limits.maxComputeWorkGroupSize[2])
	goProperties.Limits.SubPixelPrecisionBits = uint32(pProperties.limits.subPixelPrecisionBits)
	goProperties.Limits.SubTexelPrecisionBits = uint32(pProperties.limits.subTexelPrecisionBits)
	goProperties.Limits.MipmapPrecisionBits = uint32(pProperties.limits.mipmapPrecisionBits)
	goProperties.Limits.MaxDrawIndexedIndexValue = uint32(pProperties.limits.maxDrawIndexedIndexValue)
	goProperties.Limits.MaxDrawIndirectCount = uint32(pProperties.limits.maxDrawIndirectCount)
	goProperties.Limits.MaxSamplerLodBias = float32(pProperties.limits.maxSamplerLodBias)
	goProperties.Limits.MaxSamplerAnisotropy = float32(pProperties.limits.maxSamplerAnisotropy)
	goProperties.Limits.MaxViewports = uint32(pProperties.limits.maxViewports)
	goProperties.Limits.MaxViewportDimensions[0] = uint32(pProperties.limits.maxViewportDimensions[0])
	goProperties.Limits.MaxViewportDimensions[1] = uint32(pProperties.limits.maxViewportDimensions[1])
	goProperties.Limits.ViewportBoundsRange[0] = float32(pProperties.limits.viewportBoundsRange[0])
	goProperties.Limits.ViewportBoundsRange[1] = float32(pProperties.limits.viewportBoundsRange[1])
	goProperties.Limits.ViewportSubPixelBits = uint32(pProperties.limits.viewportSubPixelBits)
	goProperties.Limits.MinMemoryMapAlignment = uint64(pProperties.limits.minMemoryMapAlignment)
	goProperties.Limits.MinTexelBufferOffsetAlignment = DeviceSize(pProperties.limits.minTexelBufferOffsetAlignment)
	goProperties.Limits.MinUniformBufferOffsetAlignment = DeviceSize(pProperties.limits.minUniformBufferOffsetAlignment)
	goProperties.Limits.MinStorageBufferOffsetAlignment = DeviceSize(pProperties.limits.minStorageBufferOffsetAlignment)
	goProperties.Limits.MinTexelOffset = int32(pProperties.limits.minTexelOffset)
	goProperties.Limits.MaxTexelOffset = uint32(pProperties.limits.maxTexelOffset)
	goProperties.Limits.MinTexelGatherOffset = int32(pProperties.limits.minTexelGatherOffset)
	goProperties.Limits.MaxTexelGatherOffset = uint32(pProperties.limits.maxTexelGatherOffset)
	goProperties.Limits.MinInterpolationOffset = float32(pProperties.limits.minInterpolationOffset)
	goProperties.Limits.MaxInterpolationOffset = float32(pProperties.limits.maxInterpolationOffset)
	goProperties.Limits.SubPixelInterpolationOffsetBits = uint32(pProperties.limits.subPixelInterpolationOffsetBits)
	goProperties.Limits.MaxFramebufferWidth = uint32(pProperties.limits.maxFramebufferWidth)
	goProperties.Limits.MaxFramebufferHeight = uint32(pProperties.limits.maxFramebufferHeight)
	goProperties.Limits.MaxFramebufferLayers = uint32(pProperties.limits.maxFramebufferLayers)
	goProperties.Limits.FramebufferColorSampleCounts = SampleCountFlags(pProperties.limits.framebufferColorSampleCounts)
	goProperties.Limits.FramebufferDepthSampleCounts = SampleCountFlags(pProperties.limits.framebufferDepthSampleCounts)
	goProperties.Limits.FramebufferStencilSampleCounts = SampleCountFlags(pProperties.limits.framebufferStencilSampleCounts)
	goProperties.Limits.FramebufferNoAttachmentsSampleCounts = SampleCountFlags(pProperties.limits.framebufferNoAttachmentsSampleCounts)
	goProperties.Limits.MaxColorAttachments = uint32(pProperties.limits.maxColorAttachments)
	goProperties.Limits.SampledImageColorSampleCounts = SampleCountFlags(pProperties.limits.sampledImageColorSampleCounts)
	goProperties.Limits.SampledImageIntegerSampleCounts = SampleCountFlags(pProperties.limits.sampledImageIntegerSampleCounts)
	goProperties.Limits.SampledImageDepthSampleCounts = SampleCountFlags(pProperties.limits.sampledImageDepthSampleCounts)
	goProperties.Limits.SampledImageStencilSampleCounts = SampleCountFlags(pProperties.limits.sampledImageStencilSampleCounts)
	goProperties.Limits.StorageImageSampleCounts = SampleCountFlags(pProperties.limits.storageImageSampleCounts)
	goProperties.Limits.MaxSampleMaskWords = uint32(pProperties.limits.maxSampleMaskWords)
	goProperties.Limits.TimestampComputeAndGraphics = Bool32(pProperties.limits.timestampComputeAndGraphics)
	goProperties.Limits.TimestampPeriod = float32(pProperties.limits.timestampPeriod)
	goProperties.Limits.MaxClipDistances = uint32(pProperties.limits.maxClipDistances)
	goProperties.Limits.MaxCullDistances = uint32(pProperties.limits.maxCullDistances)
	goProperties.Limits.MaxCombinedClipAndCullDistances = uint32(pProperties.limits.maxCombinedClipAndCullDistances)
	goProperties.Limits.DiscreteQueuePriorities = uint32(pProperties.limits.discreteQueuePriorities)
	goProperties.Limits.PointSizeRange[0] = float32(pProperties.limits.pointSizeRange[0])
	goProperties.Limits.PointSizeRange[1] = float32(pProperties.limits.pointSizeRange[1])
	goProperties.Limits.LineWidthRange[0] = float32(pProperties.limits.lineWidthRange[0])
	goProperties.Limits.LineWidthRange[1] = float32(pProperties.limits.lineWidthRange[1])
	goProperties.Limits.PointSizeGranularity = float32(pProperties.limits.pointSizeGranularity)
	goProperties.Limits.LineWidthGranularity = float32(pProperties.limits.lineWidthGranularity)
	goProperties.Limits.StrictLines = Bool32(pProperties.limits.strictLines)
	goProperties.Limits.StandardSampleLocations = Bool32(pProperties.limits.standardSampleLocations)
	goProperties.Limits.OptimalBufferCopyOffsetAlignment = DeviceSize(pProperties.limits.optimalBufferCopyOffsetAlignment)
	goProperties.Limits.OptimalBufferCopyRowPitchAlignment = DeviceSize(pProperties.limits.optimalBufferCopyRowPitchAlignment)
	goProperties.Limits.NonCoherentAtomSize = DeviceSize(pProperties.limits.nonCoherentAtomSize)

	//FIXME!!!
	//goProperties.SparseProperties = pProperties.sparseProperties

	return goProperties
}

func GetPhysicalDeviceFeatures(physicalDevice PhysicalDevice) PhysicalDeviceFeatures {
	var (
		pFeatures  C.VkPhysicalDeviceFeatures
		goFeatures PhysicalDeviceFeatures
	)

	C.vkGetPhysicalDeviceFeatures(physicalDevice, &pFeatures)

	goFeatures.RobustBufferAccess = Bool32(pFeatures.robustBufferAccess)
	goFeatures.FullDrawIndexUint32 = Bool32(pFeatures.fullDrawIndexUint32)
	goFeatures.ImageCubeArray = Bool32(pFeatures.imageCubeArray)
	goFeatures.IndependentBlend = Bool32(pFeatures.independentBlend)
	goFeatures.GeometryShader = Bool32(pFeatures.geometryShader)
	goFeatures.TessellationShader = Bool32(pFeatures.tessellationShader)
	goFeatures.SampleRateShading = Bool32(pFeatures.sampleRateShading)
	goFeatures.DualSrcBlend = Bool32(pFeatures.dualSrcBlend)
	goFeatures.LogicOp = Bool32(pFeatures.logicOp)
	goFeatures.MultiDrawIndirect = Bool32(pFeatures.multiDrawIndirect)
	goFeatures.DrawIndirectFirstInstance = Bool32(pFeatures.drawIndirectFirstInstance)
	goFeatures.DepthClamp = Bool32(pFeatures.depthClamp)
	goFeatures.DepthBiasClamp = Bool32(pFeatures.depthBiasClamp)
	goFeatures.FillModeNonSolid = Bool32(pFeatures.fillModeNonSolid)
	goFeatures.DepthBounds = Bool32(pFeatures.depthBounds)
	goFeatures.WideLines = Bool32(pFeatures.wideLines)
	goFeatures.LargePoints = Bool32(pFeatures.largePoints)
	goFeatures.AlphaToOne = Bool32(pFeatures.alphaToOne)
	goFeatures.MultiViewport = Bool32(pFeatures.multiViewport)
	goFeatures.SamplerAnisotropy = Bool32(pFeatures.samplerAnisotropy)
	goFeatures.TextureCompressionETC2 = Bool32(pFeatures.textureCompressionETC2)
	goFeatures.TextureCompressionASTC_LDR = Bool32(pFeatures.textureCompressionASTC_LDR)
	goFeatures.TextureCompressionBC = Bool32(pFeatures.textureCompressionBC)
	goFeatures.OcclusionQueryPrecise = Bool32(pFeatures.occlusionQueryPrecise)
	goFeatures.PipelineStatisticsQuery = Bool32(pFeatures.pipelineStatisticsQuery)
	goFeatures.VertexPipelineStoresAndAtomics = Bool32(pFeatures.vertexPipelineStoresAndAtomics)
	goFeatures.FragmentStoresAndAtomics = Bool32(pFeatures.fragmentStoresAndAtomics)
	goFeatures.ShaderTessellationAndGeometryPointSize = Bool32(pFeatures.shaderTessellationAndGeometryPointSize)
	goFeatures.ShaderImageGatherExtended = Bool32(pFeatures.shaderImageGatherExtended)
	goFeatures.ShaderStorageImageExtendedFormats = Bool32(pFeatures.shaderStorageImageExtendedFormats)
	goFeatures.ShaderStorageImageMultisample = Bool32(pFeatures.shaderStorageImageMultisample)
	goFeatures.ShaderStorageImageReadWithoutFormat = Bool32(pFeatures.shaderStorageImageReadWithoutFormat)
	goFeatures.ShaderStorageImageWriteWithoutFormat = Bool32(pFeatures.shaderStorageImageWriteWithoutFormat)
	goFeatures.ShaderUniformBufferArrayDynamicIndexing = Bool32(pFeatures.shaderUniformBufferArrayDynamicIndexing)
	goFeatures.ShaderSampledImageArrayDynamicIndexing = Bool32(pFeatures.shaderSampledImageArrayDynamicIndexing)
	goFeatures.ShaderStorageBufferArrayDynamicIndexing = Bool32(pFeatures.shaderStorageBufferArrayDynamicIndexing)
	goFeatures.ShaderStorageImageArrayDynamicIndexing = Bool32(pFeatures.shaderStorageImageArrayDynamicIndexing)
	goFeatures.ShaderClipDistance = Bool32(pFeatures.shaderClipDistance)
	goFeatures.ShaderCullDistance = Bool32(pFeatures.shaderCullDistance)
	goFeatures.ShaderFloat64 = Bool32(pFeatures.shaderFloat64)
	goFeatures.ShaderInt64 = Bool32(pFeatures.shaderInt64)
	goFeatures.ShaderInt16 = Bool32(pFeatures.shaderInt16)
	goFeatures.ShaderResourceResidency = Bool32(pFeatures.shaderResourceResidency)
	goFeatures.ShaderResourceMinLod = Bool32(pFeatures.shaderResourceMinLod)
	goFeatures.SparseBinding = Bool32(pFeatures.sparseBinding)
	goFeatures.SparseResidencyBuffer = Bool32(pFeatures.sparseResidencyBuffer)
	goFeatures.SparseResidencyImage2D = Bool32(pFeatures.sparseResidencyImage2D)
	goFeatures.SparseResidencyImage3D = Bool32(pFeatures.sparseResidencyImage3D)
	goFeatures.SparseResidency2Samples = Bool32(pFeatures.sparseResidency2Samples)
	goFeatures.SparseResidency4Samples = Bool32(pFeatures.sparseResidency4Samples)
	goFeatures.SparseResidency8Samples = Bool32(pFeatures.sparseResidency8Samples)
	goFeatures.SparseResidency16Samples = Bool32(pFeatures.sparseResidency16Samples)
	goFeatures.SparseResidencyAliased = Bool32(pFeatures.sparseResidencyAliased)
	goFeatures.VariableMultisampleRate = Bool32(pFeatures.variableMultisampleRate)
	goFeatures.InheritedQueries = Bool32(pFeatures.inheritedQueries)

	return goFeatures
}

func GetPhysicalDeviceQueueFamilyProperties(physicalDevice PhysicalDevice) (uint32, []QueueFamilyProperties) {
	var (
		pQueueFamilyPropertyCount C.uint32_t
		pQueueFamilyProperties    *C.VkQueueFamilyProperties
		goQueueFamilyProperties   []QueueFamilyProperties
	)

	C.vkGetPhysicalDeviceQueueFamilyProperties(physicalDevice, &pQueueFamilyPropertyCount, nil)

	pQueueFamilyProperties = allocVkStructs(pQueueFamilyProperties, uint32(pQueueFamilyPropertyCount))
	defer cFree(pQueueFamilyProperties)

	C.vkGetPhysicalDeviceQueueFamilyProperties(physicalDevice, &pQueueFamilyPropertyCount, pQueueFamilyProperties)

	goQueueFamilyProperties = make([]QueueFamilyProperties, pQueueFamilyPropertyCount)
	decoded := unsafe.Slice(pQueueFamilyProperties, int(pQueueFamilyPropertyCount))

	for i := range decoded {
		goQueueFamilyProperties[i].QueueFlags = QueueFlags(decoded[i].queueFlags)
		goQueueFamilyProperties[i].QueueCount = uint32(decoded[i].queueCount)
		goQueueFamilyProperties[i].TimestampValidBits = uint32(decoded[i].timestampValidBits)
		goQueueFamilyProperties[i].MinImageTransferGranularity.depth = uint32(decoded[i].minImageTransferGranularity.depth)
		goQueueFamilyProperties[i].MinImageTransferGranularity.height = uint32(decoded[i].minImageTransferGranularity.height)
		goQueueFamilyProperties[i].MinImageTransferGranularity.width = uint32(decoded[i].minImageTransferGranularity.width)
	}

	return uint32(pQueueFamilyPropertyCount), goQueueFamilyProperties
}

func DestroySurfaceKHR(instance PInstance, surface PSurfaceKHR, pAllocator *AllocationCallbacks) {
	C.vkDestroySurfaceKHR(instance, surface, nil)
}

func GetPhysicalDeviceSurfaceCapabilitiesKHR(physicalDevice PhysicalDevice, surface PSurfaceKHR) (SurfaceCapabilitiesKHR, Result) {
	var (
		pSurfaceCapabilities  C.VkSurfaceCapabilitiesKHR
		goSurfaceCapabilities SurfaceCapabilitiesKHR
	)

	err := C.vkGetPhysicalDeviceSurfaceCapabilitiesKHR(physicalDevice, surface, &pSurfaceCapabilities)
	if err != SUCCESS {
		return SurfaceCapabilitiesKHR{}, Result(err)
	}

	goSurfaceCapabilities.MinImageCount = uint32(pSurfaceCapabilities.minImageCount)
	goSurfaceCapabilities.MaxImageCount = uint32(pSurfaceCapabilities.maxImageCount)
	goSurfaceCapabilities.CurrentExtent.height = uint32(pSurfaceCapabilities.currentExtent.height)
	goSurfaceCapabilities.CurrentExtent.width = uint32(pSurfaceCapabilities.currentExtent.width)
	goSurfaceCapabilities.MinImageExtent.height = uint32(pSurfaceCapabilities.currentExtent.height)
	goSurfaceCapabilities.MinImageExtent.width = uint32(pSurfaceCapabilities.currentExtent.width)
	goSurfaceCapabilities.MaxImageExtent.height = uint32(pSurfaceCapabilities.maxImageExtent.height)
	goSurfaceCapabilities.MaxImageArrayLayers = uint32(pSurfaceCapabilities.maxImageArrayLayers)
	goSurfaceCapabilities.SupportedTransforms = VkSurfaceTransformFlagsKHR(pSurfaceCapabilities.supportedTransforms)
	goSurfaceCapabilities.CurrentTransform = VkSurfaceTransformFlagBitsKHR(pSurfaceCapabilities.currentTransform)
	goSurfaceCapabilities.SupportedCompositeAlpha = VkCompositeAlphaFlagsKHR(pSurfaceCapabilities.supportedCompositeAlpha)
	goSurfaceCapabilities.SupportedUsageFlags = VkImageUsageFlags(pSurfaceCapabilities.supportedUsageFlags)

	return goSurfaceCapabilities, SUCCESS
}
