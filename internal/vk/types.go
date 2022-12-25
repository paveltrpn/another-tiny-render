package vk

/*
#include <vulkan/vulkan.h>
*/
import "C"

type StructureType int

type Result int

type PhysicalDeviceType int

type Flags uint32
type InstanceCreateFlags Flags
type SampleCountFlags Flags
type QueueFlags Flags
type DeviceQueueCreateFlags Flags
type Bool32 uint32
type VkSurfaceTransformFlagsKHR Flags
type VkSurfaceTransformFlagBitsKHR Flags
type VkCompositeAlphaFlagsKHR Flags
type VkImageUsageFlags Flags

type AllocationCallbacks C.VkAllocationCallbacks

type ExtensionProperties struct {
	ExtensionName string
	SpecVersion   uint32
}

type LayerProperties struct {
	LayerName             string
	SpecVersion           uint32
	ImplementationVersion uint32
	Description           string
}

type InstanceCreateInfo struct {
	SType StructureType
	//const void* pNext;
	Flags                   InstanceCreateFlags
	PApplicationInfo        *ApplicationInfo
	EnabledLayerCount       uint32
	PpEnabledLayerNames     []string
	EnabledExtensionCount   uint32
	PpEnabledExtensionNames []string
}

type ApplicationInfo struct {
	SType StructureType
	//const void*  pNext;
	PApplicationName   string
	ApplicationVersion uint32
	PEngineName        string
	EngineVersion      uint32
	ApiVersion         uint32
}

type PhysicalDevice C.VkPhysicalDevice
type PInstance C.VkInstance
type PSurfaceKHR C.VkSurfaceKHR
type DeviceSize uint64

type PhysicalDeviceLimits struct {
	MaxImageDimension1D                             uint32
	MaxImageDimension2D                             uint32
	MaxImageDimension3D                             uint32
	MaxImageDimensionCube                           uint32
	MaxImageArrayLayers                             uint32
	MaxTexelBufferElements                          uint32
	MaxUniformBufferRange                           uint32
	MaxStorageBufferRange                           uint32
	MaxPushConstantsSize                            uint32
	MaxMemoryAllocationCount                        uint32
	MaxSamplerAllocationCount                       uint32
	BufferImageGranularity                          DeviceSize
	SparseAddressSpaceSize                          DeviceSize
	MaxBoundDescriptorSets                          uint32
	MaxPerStageDescriptorSamplers                   uint32
	MaxPerStageDescriptorUniformBuffers             uint32
	MaxPerStageDescriptorStorageBuffers             uint32
	MaxPerStageDescriptorSampledImages              uint32
	MaxPerStageDescriptorStorageImages              uint32
	MaxPerStageDescriptorInputAttachments           uint32
	MaxPerStageResources                            uint32
	MaxDescriptorSetSamplers                        uint32
	MaxDescriptorSetUniformBuffers                  uint32
	MaxDescriptorSetUniformBuffersDynamic           uint32
	MaxDescriptorSetStorageBuffers                  uint32
	MaxDescriptorSetStorageBuffersDynamic           uint32
	MaxDescriptorSetSampledImages                   uint32
	MaxDescriptorSetStorageImages                   uint32
	MaxDescriptorSetInputAttachments                uint32
	MaxVertexInputAttributes                        uint32
	MaxVertexInputBindings                          uint32
	MaxVertexInputAttributeOffset                   uint32
	MaxVertexInputBindingStride                     uint32
	MaxVertexOutputComponents                       uint32
	MaxTessellationGenerationLevel                  uint32
	MaxTessellationPatchSize                        uint32
	MaxTessellationControlPerVertexInputComponents  uint32
	MaxTessellationControlPerVertexOutputComponents uint32
	MaxTessellationControlPerPatchOutputComponents  uint32
	MaxTessellationControlTotalOutputComponents     uint32
	MaxTessellationEvaluationInputComponents        uint32
	MaxTessellationEvaluationOutputComponents       uint32
	MaxGeometryShaderInvocations                    uint32
	MaxGeometryInputComponents                      uint32
	MaxGeometryOutputComponents                     uint32
	MaxGeometryOutputVertices                       uint32
	MaxGeometryTotalOutputComponents                uint32
	MaxFragmentInputComponents                      uint32
	MaxFragmentOutputAttachments                    uint32
	MaxFragmentDualSrcAttachments                   uint32
	MaxFragmentCombinedOutputResources              uint32
	MaxComputeSharedMemorySize                      uint32
	MaxComputeWorkGroupCount                        [3]uint32
	MaxComputeWorkGroupInvocations                  uint32
	MaxComputeWorkGroupSize                         [3]uint32
	SubPixelPrecisionBits                           uint32
	SubTexelPrecisionBits                           uint32
	MipmapPrecisionBits                             uint32
	MaxDrawIndexedIndexValue                        uint32
	MaxDrawIndirectCount                            uint32
	MaxSamplerLodBias                               float32
	MaxSamplerAnisotropy                            float32
	MaxViewports                                    uint32
	MaxViewportDimensions                           [2]uint32
	ViewportBoundsRange                             [2]float32
	ViewportSubPixelBits                            uint32
	MinMemoryMapAlignment                           uint64
	MinTexelBufferOffsetAlignment                   DeviceSize
	MinUniformBufferOffsetAlignment                 DeviceSize
	MinStorageBufferOffsetAlignment                 DeviceSize
	MinTexelOffset                                  int32
	MaxTexelOffset                                  uint32
	MinTexelGatherOffset                            int32
	MaxTexelGatherOffset                            uint32
	MinInterpolationOffset                          float32
	MaxInterpolationOffset                          float32
	SubPixelInterpolationOffsetBits                 uint32
	MaxFramebufferWidth                             uint32
	MaxFramebufferHeight                            uint32
	MaxFramebufferLayers                            uint32
	FramebufferColorSampleCounts                    SampleCountFlags
	FramebufferDepthSampleCounts                    SampleCountFlags
	FramebufferStencilSampleCounts                  SampleCountFlags
	FramebufferNoAttachmentsSampleCounts            SampleCountFlags
	MaxColorAttachments                             uint32
	SampledImageColorSampleCounts                   SampleCountFlags
	SampledImageIntegerSampleCounts                 SampleCountFlags
	SampledImageDepthSampleCounts                   SampleCountFlags
	SampledImageStencilSampleCounts                 SampleCountFlags
	StorageImageSampleCounts                        SampleCountFlags
	MaxSampleMaskWords                              uint32
	TimestampComputeAndGraphics                     Bool32
	TimestampPeriod                                 float32
	MaxClipDistances                                uint32
	MaxCullDistances                                uint32
	MaxCombinedClipAndCullDistances                 uint32
	DiscreteQueuePriorities                         uint32
	PointSizeRange                                  [2]float32
	LineWidthRange                                  [2]float32
	PointSizeGranularity                            float32
	LineWidthGranularity                            float32
	StrictLines                                     Bool32
	StandardSampleLocations                         Bool32
	OptimalBufferCopyOffsetAlignment                DeviceSize
	OptimalBufferCopyRowPitchAlignment              DeviceSize
	NonCoherentAtomSize                             DeviceSize
}

type PhysicalDeviceFeatures struct {
	RobustBufferAccess                      Bool32
	FullDrawIndexUint32                     Bool32
	ImageCubeArray                          Bool32
	IndependentBlend                        Bool32
	GeometryShader                          Bool32
	TessellationShader                      Bool32
	SampleRateShading                       Bool32
	DualSrcBlend                            Bool32
	LogicOp                                 Bool32
	MultiDrawIndirect                       Bool32
	DrawIndirectFirstInstance               Bool32
	DepthClamp                              Bool32
	DepthBiasClamp                          Bool32
	FillModeNonSolid                        Bool32
	DepthBounds                             Bool32
	WideLines                               Bool32
	LargePoints                             Bool32
	AlphaToOne                              Bool32
	MultiViewport                           Bool32
	SamplerAnisotropy                       Bool32
	TextureCompressionETC2                  Bool32
	TextureCompressionASTC_LDR              Bool32
	TextureCompressionBC                    Bool32
	OcclusionQueryPrecise                   Bool32
	PipelineStatisticsQuery                 Bool32
	VertexPipelineStoresAndAtomics          Bool32
	FragmentStoresAndAtomics                Bool32
	ShaderTessellationAndGeometryPointSize  Bool32
	ShaderImageGatherExtended               Bool32
	ShaderStorageImageExtendedFormats       Bool32
	ShaderStorageImageMultisample           Bool32
	ShaderStorageImageReadWithoutFormat     Bool32
	ShaderStorageImageWriteWithoutFormat    Bool32
	ShaderUniformBufferArrayDynamicIndexing Bool32
	ShaderSampledImageArrayDynamicIndexing  Bool32
	ShaderStorageBufferArrayDynamicIndexing Bool32
	ShaderStorageImageArrayDynamicIndexing  Bool32
	ShaderClipDistance                      Bool32
	ShaderCullDistance                      Bool32
	ShaderFloat64                           Bool32
	ShaderInt64                             Bool32
	ShaderInt16                             Bool32
	ShaderResourceResidency                 Bool32
	ShaderResourceMinLod                    Bool32
	SparseBinding                           Bool32
	SparseResidencyBuffer                   Bool32
	SparseResidencyImage2D                  Bool32
	SparseResidencyImage3D                  Bool32
	SparseResidency2Samples                 Bool32
	SparseResidency4Samples                 Bool32
	SparseResidency8Samples                 Bool32
	SparseResidency16Samples                Bool32
	SparseResidencyAliased                  Bool32
	VariableMultisampleRate                 Bool32
	InheritedQueries                        Bool32
}

type PhysicalDeviceSparseProperties struct {
	residencyStandard2DBlockShape            Bool32
	residencyStandard2DMultisampleBlockShape Bool32
	residencyStandard3DBlockShape            Bool32
	residencyAlignedMipSize                  Bool32
	residencyNonResidentStrict               Bool32
}

type PhysicalDeviceProperties struct {
	ApiVersion        uint32
	DriverVersion     uint32
	VendorID          uint32
	DeviceID          uint32
	DeviceType        PhysicalDeviceType
	DeviceName        string
	PipelineCacheUUID []uint8
	Limits            PhysicalDeviceLimits
	SparseProperties  PhysicalDeviceSparseProperties
}

type Extent2D struct {
	width  uint32
	height uint32
}

type Extent3D struct {
	width  uint32
	height uint32
	depth  uint32
}

type QueueFamilyProperties struct {
	QueueFlags                  QueueFlags
	QueueCount                  uint32
	TimestampValidBits          uint32
	MinImageTransferGranularity Extent3D
}

type DeviceQueueCreateInfo struct {
	sType StructureType
	//const void*                 pNext;
	flags            DeviceQueueCreateFlags
	queueFamilyIndex uint32
	queueCount       uint32
	pQueuePriorities float32
}

type SurfaceCapabilitiesKHR struct {
	MinImageCount           uint32
	MaxImageCount           uint32
	CurrentExtent           Extent2D
	MinImageExtent          Extent2D
	MaxImageExtent          Extent2D
	MaxImageArrayLayers     uint32
	SupportedTransforms     VkSurfaceTransformFlagsKHR
	CurrentTransform        VkSurfaceTransformFlagBitsKHR
	SupportedCompositeAlpha VkCompositeAlphaFlagsKHR
	SupportedUsageFlags     VkImageUsageFlags
}
